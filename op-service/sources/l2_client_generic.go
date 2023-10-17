package sources

import (
	"context"

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
	InfoByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, error)
	GetProof(ctx context.Context, address common.Address, storage []common.Hash, blockTag string) (*eth.AccountResult, error)

	InfoByLabel(ctx context.Context, label eth.BlockLabel) (eth.BlockInfo, error)

	ChainID(context.Context) (*big.Int, error)

	Close()
}
