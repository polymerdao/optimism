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
	// TODO use big.Int for all num related api? requires some changes to op-node and friends though
	PayloadByNumber(ctx context.Context, num uint64) (*eth.ExecutionPayload, error)
	PayloadByHash(ctx context.Context, hash common.Hash) (*eth.ExecutionPayload, error)
	InfoByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, error)
	GetProof(ctx context.Context, address common.Address, storage []common.Hash, blockTag string) (eth.Proof, error)

	InfoByLabel(ctx context.Context, label eth.BlockLabel) (eth.BlockInfo, error)

	ChainID(context.Context) (*big.Int, error)

	InfoAndTxsByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, types.Transactions, error)
	InfoAndTxsByNumber(ctx context.Context, number uint64) (eth.BlockInfo, types.Transactions, error)
	InfoAndTxsByLabel(ctx context.Context, label eth.BlockLabel) (eth.BlockInfo, types.Transactions, error)

	Close()

	// TODO add these to make the test happy
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)

	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	BlockByNumber(ctx context.Context, number *big.Int) (peptide.EthBlock, error)
}
