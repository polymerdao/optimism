package sources

import (
	"context"

	"math/big"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
)

type L2ClientGeneric interface {
	PayloadByLabel(ctx context.Context, label eth.BlockLabel) (*eth.ExecutionPayload, error)
	PayloadByNumber(ctx context.Context, num uint64) (*eth.ExecutionPayload, error)
	PayloadByHash(ctx context.Context, hash common.Hash) (*eth.ExecutionPayload, error)
	InfoByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, error)
	GetProof(ctx context.Context, address common.Address, storage []common.Hash, blockTag string) (*eth.AccountResult, error)

	InfoByLabel(ctx context.Context, label eth.BlockLabel) (eth.BlockInfo, error)

	ChainID(context.Context) (*big.Int, error)
	// L2BlockRefByNumber(context.Context, uint64) (eth.L2BlockRef, error)

	Close()
}
