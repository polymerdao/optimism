package sources

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"strconv"
	"time"

	"math/big"
)

type ABCIClient struct {
	TxClient *cosmosclient.Client
}

func NewABCIClient(cosmosClient *cosmosclient.Client) *ABCIClient {
	return &ABCIClient{cosmosClient}
}

func NewABCIClientWithTimeout(timeout time.Duration, url string) (*ABCIClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	gasAdjFactor := 1.5

	// Create TxClient with in memory keyring backend
	txClient, err := cosmosclient.New(
		ctx,
		cosmosclient.WithKeyringBackend(cosmosaccount.KeyringMemory),
		cosmosclient.WithAddressPrefix("polymer"),
		cosmosclient.WithGas("auto"),
		cosmosclient.WithGasAdjustment(gasAdjFactor),
		cosmosclient.WithNodeAddress(url),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create a cosmos client client: %w", err)
	}

	// Cosmos client doens't automatically pass through customized gasAdjFactor to factory, do it manually here
	txClient.TxFactory = txClient.TxFactory.WithGasAdjustment(gasAdjFactor)

	return NewABCIClient(&txClient), nil
}

func (c *ABCIClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	height := number.Int64()
	block, err := c.TxClient.RPC.Block(ctx, &height)
	if err != nil {
		return nil, fmt.Errorf("unable to get a cosmos block at height %d: %w", height, err)
	}

	blockResults, err := c.TxClient.RPC.BlockResults(ctx, &height)
	if err != nil {
		return nil, fmt.Errorf("unable to get cosmos block results at height %d: %w", height, err)
	}

	//calculate gas used across all tx results
	var gasUsed uint64
	for _, txResult := range blockResults.TxsResults {
		gasUsed += uint64(txResult.GasUsed)
	}

	header := &types.Header{
		ParentHash:       common.BytesToHash(block.Block.LastBlockID.Hash.Bytes()),
		UncleHash:        common.Hash{},    // irrelevant for ABCI
		Coinbase:         common.Address{}, // irrelevant for ABCI
		Root:             common.BytesToHash(block.BlockID.Hash.Bytes()),
		TxHash:           common.BytesToHash(block.Block.Txs.Hash()),
		ReceiptHash:      common.BytesToHash(block.Block.Header.AppHash.Bytes()), // no direct alternative in ABCI?
		Bloom:            types.Bloom{},                                          // irrelevant for ABCI
		Difficulty:       nil,                                                    // irrelevant for ABCI
		Number:           big.NewInt(block.Block.Height),
		GasLimit:         0, // irrelevant for ABCI
		GasUsed:          gasUsed,
		Time:             uint64(block.Block.Time.Unix()),
		Extra:            nil,                // irrelevant for ABCI
		MixDigest:        common.Hash{},      // irrelevant for ABCI
		Nonce:            types.BlockNonce{}, // irrelevant for ABCI
		BaseFee:          nil,                // irrelevant for ABCI
		WithdrawalsHash:  nil,                // irrelevant for ABCI
		BlobGasUsed:      nil,                // irrelevant for ABCI
		ExcessBlobGas:    nil,                // irrelevant for ABCI
		ParentBeaconRoot: nil,                // irrelevant for ABCI
	}

	chainId, err := strconv.Atoi(block.Block.ChainID)
	if err != nil {
		return nil, fmt.Errorf("cosmos block chain id is not an integer (%s): %w", block.Block.ChainID, err)
	}

	txs := make([]*types.Transaction, len(blockResults.TxsResults))
	for i, txResult := range blockResults.TxsResults {
		txData := &types.DynamicFeeTx{
			ChainID: big.NewInt(int64(chainId)),
			Data:    txResult.Data,
			Gas:     uint64(txResult.GasUsed),
			Value:   big.NewInt(0),
			To:      nil, // TODO find how to extract this from a Cosmos Tx
		}
		txs[i] = types.NewTx(txData)
	}

	return types.NewBlockWithHeader(header).WithBody(txs, nil), nil
}
