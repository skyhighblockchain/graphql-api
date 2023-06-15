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
	"skyhigh-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	retypes "github.com/ethereum/go-ethereum/core/types"
)

// Transaction returns information about a blockchain transaction by hash.
func (skh *SkhBridge) Transaction(hash *common.Hash) (*types.Transaction, error) {
	// keep track of the operation
	skh.log.Debugf("loading transaction %s", hash.String())

	// call for data
	var trx types.Transaction
	err := skh.rpc.Call(&trx, "skh_getTransactionByHash", hash)
	if err != nil {
		skh.log.Error("transaction could not be extracted")
		return nil, err
	}

	// is there a block reference already?
	if trx.BlockNumber != nil {
		// get transaction receipt
		var rec struct {
			Index             hexutil.Uint64  `json:"transactionIndex"`
			CumulativeGasUsed hexutil.Uint64  `json:"cumulativeGasUsed"`
			GasUsed           hexutil.Uint64  `json:"gasUsed"`
			ContractAddress   *common.Address `json:"contractAddress,omitempty"`
			Status            hexutil.Uint64  `json:"status"`
			Logs              []retypes.Log   `json:"logs"`
		}

		// call for the transaction receipt data
		err := skh.rpc.Call(&rec, "skh_getTransactionReceipt", hash)
		if err != nil {
			skh.log.Errorf("can not get receipt for transaction %s", hash)
			return nil, err
		}

		// copy some data
		trx.Index = &rec.Index
		trx.CumulativeGasUsed = &rec.CumulativeGasUsed
		trx.GasUsed = &rec.GasUsed
		trx.ContractAddress = rec.ContractAddress
		trx.Status = &rec.Status
		trx.Logs = rec.Logs
	}

	// keep track of the operation
	skh.log.Debugf("transaction %s loaded", hash.String())
	return &trx, nil
}

// SendTransaction sends raw signed and RLP encoded transaction to the block chain.
func (skh *SkhBridge) SendTransaction(tx hexutil.Bytes) (*common.Hash, error) {
	// keep track of the operation
	skh.log.Debug("sending new transaction to block chain")

	var hash common.Hash
	err := skh.rpc.Call(&hash, "eth_sendRawTransaction", tx)
	if err != nil {
		skh.log.Error("transaction could not be sent")
		return nil, err
	}

	// keep track of the operation
	skh.log.Debugf("transaction has been accepted with hash %s", hash.String())
	return &hash, nil
}
