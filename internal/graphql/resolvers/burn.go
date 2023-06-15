// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"math/big"
	"skyhigh-api-graphql/internal/repository"
	"skyhigh-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// SkhBurnedTotal resolves total amount of burned SKH tokens in WEI units.
func (rs *rootResolver) SkhBurnedTotal() hexutil.Big {
	val, err := repository.R().SkhBurnTotal()
	if err != nil {
		log.Criticalf("failed to load burned total; %s", err.Error())
		return hexutil.Big{}
	}
	return hexutil.Big(*new(big.Int).Mul(big.NewInt(val), types.BurnDecimalsCorrection))
}

// SkhBurnedTotalAmount resolves total amount of burned SKH tokens in SKH units.
func (rs *rootResolver) SkhBurnedTotalAmount() float64 {
	val, err := repository.R().SkhBurnTotal()
	if err != nil {
		log.Criticalf("failed to load burned total; %s", err.Error())
		return 0
	}
	return float64(val) / types.BurnSKHDecimalsCorrection
}

// SkhLatestBlockBurnList resolves a list of the latest block burns.
func (rs *rootResolver) SkhLatestBlockBurnList(args struct{ Count int32 }) ([]types.SkhBurn, error) {
	if args.Count < 1 || args.Count > 50 {
		args.Count = 25
	}
	return repository.R().SkhBurnList(int64(args.Count))
}
