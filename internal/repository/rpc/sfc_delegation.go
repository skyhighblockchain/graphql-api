/*
Package rpc implements bridge to Push full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an SkyHigh/Push node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Push RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Push RPC interface with connection limited to specified endpoints.

We strongly discourage opening Push RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"context"
	"fmt"
	"math/big"
	"skyhigh-api-graphql/internal/repository/rpc/contracts"
	"skyhigh-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// AmountStaked returns the current amount at stake for the given staker address and target validator
func (skh *SkhBridge) AmountStaked(addr *common.Address, valID *big.Int) (*big.Int, error) {
	// keep track of the operation
	skh.log.Debugf("verifying amount staked by %s to %d", addr.String(), valID.Uint64())
	return skh.SfcContract().GetStake(skh.DefaultCallOpts(), *addr, valID)
}

// AmountStakeLocked returns the current locked amount at stake for the given staker address and target validator.
func (skh *SkhBridge) AmountStakeLocked(addr *common.Address, valID *big.Int) (*big.Int, error) {
	return skh.SfcContract().GetLockedStake(skh.DefaultCallOpts(), *addr, valID)
}

// AmountStakeUnlocked returns the current unlocked amount at stake for the given staker address and target validator.
func (skh *SkhBridge) AmountStakeUnlocked(addr *common.Address, valID *big.Int) (*big.Int, error) {
	return skh.SfcContract().GetUnlockedStake(skh.DefaultCallOpts(), *addr, valID)
}

// StakeUnlockPenalty returns the expected penalty of a premature stake unlock.
func (skh *SkhBridge) StakeUnlockPenalty(addr *common.Address, valID *big.Int, amount *big.Int) (*big.Int, error) {
	// pack call data
	cd, err := skh.SfcAbi().Pack("unlockStake", valID, amount)
	if err != nil {
		skh.log.Errorf("penalty for unlocking %d of %s to %d not available; %s", amount.Uint64(), addr.String(), valID.Uint64(), err.Error())
		return nil, err
	}

	// make the UnlockStake call as a view call to get the penalty value
	data, err := skh.eth.CallContract(context.Background(), ethereum.CallMsg{
		From: *addr,
		To:   &skh.sfcConfig.SFCContract,
		Data: cd,
	}, nil)
	if err != nil {
		skh.log.Errorf("penalty for unlocking %d of %s to %d not available; %s", amount.Uint64(), addr.String(), valID.Uint64(), err.Error())
		return nil, err
	}

	// make response size sanity check; we expect single big integer value
	if len(data) != 32 {
		skh.log.Errorf("penalty for unlocking %d of %s to %d response not valid; expected 32 bytes, received %d bytes", amount.Uint64(), addr.String(), valID.Uint64(), len(data))
		return nil, err
	}

	// parse data in to the value and return it
	return new(big.Int).SetBytes(data), nil
}

// PendingRewards returns a detail of delegation rewards waiting to be claimed for the given delegation.
func (skh *SkhBridge) PendingRewards(addr *common.Address, valID *big.Int) (*types.PendingRewards, error) {
	// prep the empty value
	pr := types.PendingRewards{
		Address: *addr,
		Staker:  hexutil.Big(*valID),
		Amount:  hexutil.Big{},
	}

	// get the pending rewards amount
	amo, err := skh.SfcContract().PendingRewards(skh.DefaultCallOpts(), *addr, valID)
	if err != nil {
		skh.log.Criticalf("can not calculate pending rewards of %s to %d; %s", addr.String(), valID.Uint64(), err.Error())
		return &pr, nil
	}

	// update the amount
	pr.Amount = hexutil.Big(*amo)
	return &pr, nil
}

// DelegationLock returns delegation lock information using SFC contract binding.
func (skh *SkhBridge) DelegationLock(addr *common.Address, valID *hexutil.Big) (dll *types.DelegationLock, err error) {
	// recover from panic here
	defer func() {
		if r := recover(); r != nil {
			skh.log.Criticalf("can not get SFC lock status on delegation %s to %d; SFC call panic", addr.String(), valID.String())
			dll = &types.DelegationLock{}
		}
	}()

	// get staker locking detail
	lock, err := skh.SfcContract().GetLockupInfo(skh.DefaultCallOpts(), *addr, valID.ToInt())
	if err != nil {
		skh.log.Errorf("delegation lock query failed; %v", err)
		return nil, err
	}

	// are lock timers available?
	if lock.FromEpoch == nil || lock.EndTime == nil {
		skh.log.Errorf("delegation lock details not available")
		return nil, fmt.Errorf("delegation lock missing")
	}

	// make a new delegation lock
	return &types.DelegationLock{
		LockedAmount:    hexutil.Big(*lock.LockedStake),
		LockedFromEpoch: hexutil.Uint64(lock.FromEpoch.Uint64()),
		LockedUntil:     hexutil.Uint64(lock.EndTime.Uint64()),
		Duration:        hexutil.Uint64(lock.Duration.Uint64()),
	}, nil
}

// DelegationOutstandingSSKH returns the amount of sSKH tokens for the delegation
// identified by the delegator address and the stakerId.
func (skh *SkhBridge) DelegationOutstandingSSKH(addr *common.Address, valID *big.Int) (*big.Int, error) {
	// log action
	skh.log.Debugf("checking outstanding sSKH of %s to %d", addr.String(), valID.Uint64())

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcTokenizer(skh.sfcConfig.TokenizerContract, skh.eth)
	if err != nil {
		skh.log.Criticalf("failed to instantiate SFC Tokenizer contract; %s", err.Error())
		return nil, err
	}

	// get the amount of outstanding sSKH
	return contract.OutstandingSSKH(skh.DefaultCallOpts(), *addr, valID)
}

// DelegationTokenizerUnlocked returns the status of SFC Tokenizer lock
// for a delegation identified by the address and staker id.
func (skh *SkhBridge) DelegationTokenizerUnlocked(addr *common.Address, valID *big.Int) (bool, error) {
	// log action
	skh.log.Debugf("checking SFC tokenizer lock of %s to %d", addr.String(), valID.Uint64())

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcTokenizer(skh.sfcConfig.TokenizerContract, skh.eth)
	if err != nil {
		skh.log.Criticalf("failed to instantiate SFC Tokenizer contract: %s", err.Error())
		return false, err
	}

	// get the lock status
	lock, err := contract.AllowedToWithdrawStake(skh.DefaultCallOpts(), *addr, valID)
	if err != nil {
		skh.log.Criticalf("failed to get SFC Tokenizer lock status of %s to %d; %s", addr.String(), valID.Uint64(), err.Error())
		return false, err
	}

	return lock, nil
}
