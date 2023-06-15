/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access SkyHigh/Push full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"skyhigh-api-graphql/internal/types"
)

// StoreSkhBurn stores the given native SKH burn per block record into the persistent storage.
func (p *proxy) StoreSkhBurn(burn *types.SkhBurn) error {
	p.cache.SkhBurnUpdate(burn, p.db.BurnTotal)
	return p.db.StoreBurn(burn)
}

// SkhBurnTotal provides the total amount of burned native SKH.
func (p *proxy) SkhBurnTotal() (int64, error) {
	return p.cache.SkhBurnTotal(p.db.BurnTotal)
}

// SkhBurnList provides list of per-block burned native SKH tokens.
func (p *proxy) SkhBurnList(count int64) ([]types.SkhBurn, error) {
	return p.db.BurnList(count)
}
