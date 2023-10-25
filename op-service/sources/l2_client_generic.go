package sources

import (
	"context"
	"github.com/ethereum-optimism/optimism/op-service/peptide"
	"github.com/ethereum/go-ethereum/core/types"

	"math/big"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
)

// L2ClientGeneric provides a set of methods for interacting with Layer 2 (L2) data. It is implemented by different
// clients, such as PolymerClient and EthClient, to provide specific behaviors for these interactions. In l2_client.go,
// it is embedded by L2Client struct, extending its functionality by adding caching and configuration capabilities. In
// node.go, it is used by OpNode struct through l2Source field to interact with the L2 Execution Engine via RPC
// bindings.
type L2ClientGeneric interface {
	PayloadByLabel(ctx context.Context, label eth.BlockLabel) (*eth.ExecutionPayload, error)
	PayloadByNumber(ctx context.Context, num uint64) (*eth.ExecutionPayload, error)
	PayloadByHash(ctx context.Context, hash common.Hash) (*eth.ExecutionPayload, error)
	GetProof(ctx context.Context, address common.Address, storage []common.Hash, blockTag string) (*eth.AccountResult, error)
	L2BlockRefByLabel(ctx context.Context, label eth.BlockLabel) (eth.L2BlockRef, error)

	InfoByLabel(ctx context.Context, label eth.BlockLabel) (eth.BlockInfo, error)
	InfoByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, error)
	InfoAndTxsByHash(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Transactions, error)

	ChainID(context.Context) (*big.Int, error)

	Close()

	GetPayload(ctx context.Context, payloadId eth.PayloadID) (*eth.ExecutionPayload, error)
	ForkchoiceUpdate(ctx context.Context, state *eth.ForkchoiceState, attr *eth.PayloadAttributes) (*eth.ForkchoiceUpdatedResult, error)
	NewPayload(ctx context.Context, payload *eth.ExecutionPayload) (*eth.PayloadStatusV1, error)

	BlockByNumber(ctx context.Context, number *big.Int) (peptide.EthBlock, error)
}
