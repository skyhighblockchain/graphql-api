/*
Package rpc implements bridge to Push full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an SkyHigh/Push node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Push RPC interface for remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Push RPC interface with connection limited to specified endpoints.

We strongly discourage opening Push RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum"
)

// skhHeadsObserverSubscribeTick represents the time between subscription attempts.
const skhHeadsObserverSubscribeTick = 30 * time.Second

// observeBlocks collects new blocks from the blockchain network
// and posts them into the proxy channel for processing.
func (skh *SkhBridge) observeBlocks() {
	var sub ethereum.Subscription
	defer func() {
		if sub != nil {
			sub.Unsubscribe()
		}
		skh.log.Noticef("block observer done")
		skh.wg.Done()
	}()

	sub = skh.blockSubscription()
	for {
		// re-subscribe if the subscription ref is not valid
		if sub == nil {
			tm := time.NewTimer(skhHeadsObserverSubscribeTick)
			select {
			case <-skh.sigClose:
				return
			case <-tm.C:
				sub = skh.blockSubscription()
				continue
			}
		}

		// use the subscriptions
		select {
		case <-skh.sigClose:
			return
		case err := <-sub.Err():
			skh.log.Errorf("block subscription failed; %s", err.Error())
			sub = nil
		}
	}
}

// blockSubscription provides a subscription for new blocks received
// by the connected blockchain node.
func (skh *SkhBridge) blockSubscription() ethereum.Subscription {
	sub, err := skh.rpc.EthSubscribe(context.Background(), skh.headers, "newHeads")
	if err != nil {
		skh.log.Criticalf("can not observe new blocks; %s", err.Error())
		return nil
	}
	return sub
}
