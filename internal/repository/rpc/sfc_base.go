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

//go:generate tools/abigen.sh --abi ./contracts/abi/sfc-1.1.abi --pkg contracts --type SfcV1Contract --out ./contracts/sfc-v1.go
//go:generate tools/abigen.sh --abi ./contracts/abi/sfc-2.0.4-rc.2.abi --pkg contracts --type SfcV2Contract --out ./contracts/sfc-v2.go
//go:generate tools/abigen.sh --abi ./contracts/abi/sfc-3.0-rc.1.abi --pkg contracts --type SfcContract --out ./contracts/sfc-v3.go
//go:generate tools/abigen.sh --abi ./contracts/abi/sfc-tokenizer.abi --pkg contracts --type SfcTokenizer --out ./contracts/sfc_tokenizer.go

import (
	"math/big"
	"skyhigh-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// sfcFirstLockEpoch represents the first epoch with stake locking available.
const sfcFirstLockEpoch uint64 = 1600

// SfcVersion returns current version of the SFC contract as a single number.
func (skh *SkhBridge) SfcVersion() (hexutil.Uint64, error) {
	// get the version information from the contract
	var ver [3]byte
	var err error
	ver, err = skh.SfcContract().Version(nil)
	if err != nil {
		skh.log.Criticalf("failed to get the SFC version; %s", err.Error())
		return 0, err
	}
	return hexutil.Uint64((uint64(ver[0]) << 16) | (uint64(ver[1]) << 8) | uint64(ver[2])), nil
}

// CurrentEpoch extract the current epoch id from SFC smart contract.
func (skh *SkhBridge) CurrentEpoch() (hexutil.Uint64, error) {
	// get the value from the contract
	epoch, err := skh.SfcContract().CurrentEpoch(skh.DefaultCallOpts())
	if err != nil {
		skh.log.Errorf("failed to get the current sealed epoch: %s", err.Error())
		return 0, err
	}
	return hexutil.Uint64(epoch.Uint64()), nil
}

// CurrentSealedEpoch extract the current sealed epoch id from SFC smart contract.
func (skh *SkhBridge) CurrentSealedEpoch() (hexutil.Uint64, error) {
	// get the value from the contract
	epoch, err := skh.SfcContract().CurrentSealedEpoch(skh.DefaultCallOpts())
	if err != nil {
		skh.log.Errorf("failed to get the current sealed epoch: %s", err.Error())
		return 0, err
	}
	return hexutil.Uint64(epoch.Uint64()), nil
}

// Epoch extract information about an epoch from SFC smart contract.
func (skh *SkhBridge) Epoch(id hexutil.Uint64) (*types.Epoch, error) {
	// extract epoch snapshot
	epo, err := skh.SfcContract().GetEpochSnapshot(nil, big.NewInt(int64(id)))
	if err != nil {
		skh.log.Errorf("failed to extract epoch information: %s", err.Error())
		return nil, err
	}

	return &types.Epoch{
		Id:                    id,
		EndTime:               hexutil.Uint64(epo.EndTime.Uint64()),
		EpochFee:              (hexutil.Big)(*epo.EpochFee),
		TotalBaseRewardWeight: (hexutil.Big)(*epo.TotalBaseRewardWeight),
		TotalTxRewardWeight:   (hexutil.Big)(*epo.TotalTxRewardWeight),
		BaseRewardPerSecond:   (hexutil.Big)(*epo.BaseRewardPerSecond),
		StakeTotalAmount:      (hexutil.Big)(*epo.TotalStake),
		TotalSupply:           (hexutil.Big)(*epo.TotalSupply),
	}, nil
}

// RewardsAllowed returns if the rewards can be manipulated with.
func (skh *SkhBridge) RewardsAllowed() (bool, error) {
	skh.log.Debug("rewards lock always open")
	return true, nil
}

// LockingAllowed indicates if the stake locking has been enabled in SFC.
func (skh *SkhBridge) LockingAllowed() (bool, error) {
	// get the current sealed epoch value from the contract
	epoch, err := skh.SfcContract().CurrentSealedEpoch(nil)
	if err != nil {
		skh.log.Errorf("failed to get the current sealed epoch: %s", err.Error())
		return false, err
	}

	return epoch.Uint64() >= sfcFirstLockEpoch, nil
}

// TotalStaked returns the total amount of staked tokens.
func (skh *SkhBridge) TotalStaked() (*big.Int, error) {
	return skh.SfcContract().TotalStake(skh.DefaultCallOpts())
}

// SfcMinValidatorStake extracts a value of minimal validator self stake.
func (skh *SkhBridge) SfcMinValidatorStake() (*big.Int, error) {
	return skh.SfcContract().MinSelfStake(skh.DefaultCallOpts())
}

// SfcMaxDelegatedRatio extracts a ratio between self delegation and received stake.
func (skh *SkhBridge) SfcMaxDelegatedRatio() (*big.Int, error) {
	return skh.SfcContract().MaxDelegatedRatio(skh.DefaultCallOpts())
}

// SfcMinLockupDuration extracts a minimal lockup duration.
func (skh *SkhBridge) SfcMinLockupDuration() (*big.Int, error) {
	return skh.SfcContract().MinLockupDuration(skh.DefaultCallOpts())
}

// SfcMaxLockupDuration extracts a maximal lockup duration.
func (skh *SkhBridge) SfcMaxLockupDuration() (*big.Int, error) {
	return skh.SfcContract().MaxLockupDuration(skh.DefaultCallOpts())
}

// SfcWithdrawalPeriodEpochs extracts a minimal number of epochs between un-delegate and withdraw.
func (skh *SkhBridge) SfcWithdrawalPeriodEpochs() (*big.Int, error) {
	return skh.SfcContract().WithdrawalPeriodEpochs(skh.DefaultCallOpts())
}

// SfcWithdrawalPeriodTime extracts a minimal number of seconds between un-delegate and withdraw.
func (skh *SkhBridge) SfcWithdrawalPeriodTime() (*big.Int, error) {
	return skh.SfcContract().WithdrawalPeriodTime(skh.DefaultCallOpts())
}
