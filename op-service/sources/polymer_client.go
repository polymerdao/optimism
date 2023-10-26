package sources

import (
	"context"
	"fmt"
	"github.com/ethereum-optimism/optimism/op-service/peptide"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"strconv"

	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"strings"
)

// PolymerClient is a L2ClientGeneric implementation that interacts with the Polymer's ABCI app as an L2 Execution
// Engine via RPC bindings.
type PolymerClient struct {
	client client.RPC
}

var _ L2ClientGeneric = (*PolymerClient)(nil)

func NewPolymerClient(client client.RPC) *PolymerClient {
	return &PolymerClient{client: client}
}

func (p *PolymerClient) L2BlockRefByHash(ctx context.Context, l2Hash common.Hash) (eth.L2BlockRef, error) {
	var blockRef eth.L2BlockRef
	err := p.client.CallContext(ctx, &blockRef, "ee_getL2BlockRefByHash", l2Hash)
	return blockRef, err
}

func (p *PolymerClient) L2BlockRefByLabel(ctx context.Context, label eth.BlockLabel) (eth.L2BlockRef, error) {
	var blockRef eth.L2BlockRef
	err := p.client.CallContext(ctx, &blockRef, "ee_getL2BlockRefByLabel", label)
	if err != nil {
		x, ok := err.(interface{ ErrorData() interface{} })
		if ok {
			if strings.Contains(x.ErrorData().(string), "not found") {
				err = ethereum.NotFound

			}
		}

		// w%: wrap to preserve ethereum.NotFound case
		return eth.L2BlockRef{}, fmt.Errorf("failed to determine L2BlockRef of %s, could not get payload: %w", label, err)
	}

	return blockRef, err
}

func (p *PolymerClient) L2BlockRefByNumber(ctx context.Context, num uint64) (eth.L2BlockRef, error) {
	var blockRef eth.L2BlockRef
	err := p.client.CallContext(ctx, &blockRef, "ee_getL2BlockRefByNumber", strconv.FormatUint(num, 10))
	if err != nil {
		x, ok := err.(interface{ ErrorData() interface{} })
		if ok {
			if strings.Contains(x.ErrorData().(string), "not found") {
				err = ethereum.NotFound

			}
		}

		// w%: wrap to preserve ethereum.NotFound case
		return eth.L2BlockRef{}, fmt.Errorf("failed to determine L2BlockRef of %d, could not get payload: %w", num, err)
	}

	return blockRef, err

}

func (p *PolymerClient) PayloadByLabel(ctx context.Context, label eth.BlockLabel) (*eth.ExecutionPayload, error) {
	var payload *eth.ExecutionPayload
	err := p.client.CallContext(ctx, &payload, "ee_getPayloadByLabel", label)
	return payload, err
}

func (p *PolymerClient) PayloadByNumber(ctx context.Context, num uint64) (*eth.ExecutionPayload, error) {
	var payload *eth.ExecutionPayload
	err := p.client.CallContext(ctx, &payload, "ee_getPayloadByNumber", strconv.FormatUint(num, 10))
	return payload, err
}

func (p *PolymerClient) PayloadByHash(ctx context.Context, hash common.Hash) (*eth.ExecutionPayload, error) {
	var payload *eth.ExecutionPayload
	err := p.client.CallContext(ctx, &payload, "ee_getPayloadByHash", hash)
	return payload, err
}

func (p *PolymerClient) InfoByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, error) {
	var info eth.BlockInfo
	err := p.client.CallContext(ctx, &info, "ee_getInfoByHash", hash)
	return info, err
}

func (p *PolymerClient) GetProof(ctx context.Context, address common.Address, storage []common.Hash, blockTag string) (*eth.AccountResult, error) {
	var result *eth.AccountResult
	err := p.client.CallContext(ctx, &result, "ee_getProof", address, storage, blockTag)
	return result, err
}

func (p *PolymerClient) InfoByLabel(ctx context.Context, label eth.BlockLabel) (eth.BlockInfo, error) {
	var info eth.BlockInfo
	err := p.client.CallContext(ctx, &info, "ee_getInfoByLabel", label)
	return info, err
}

func (p *PolymerClient) ChainID(ctx context.Context) (*big.Int, error) {
	var chainID *big.Int
	err := p.client.CallContext(ctx, &chainID, "ee_getChainID")
	return chainID, err
}

func (p *PolymerClient) BlockByNumber(ctx context.Context, number *big.Int) (peptide.EthBlock, error) {
	var block *peptide.Block
	err := p.client.CallContext(ctx, block, "ee_getBlockByNumber", number)
	return block, err
}

func (p *PolymerClient) Close() {
	p.client.Close()
}

func (p *PolymerClient) InfoAndTxsByNumber(ctx context.Context, number uint64) (eth.BlockInfo, types.Transactions, error) {
	var info eth.BlockInfo
	err := p.client.CallContext(ctx, &info, "ee_getInfoByNumber", number)
	return info, types.Transactions{}, err
}

func (p *PolymerClient) InfoAndTxsByLabel(ctx context.Context, label eth.BlockLabel) (eth.BlockInfo, types.Transactions, error) {
	var info eth.BlockInfo
	err := p.client.CallContext(ctx, &info, "ee_getInfoByLabel", label)
	return info, types.Transactions{}, err
}

func (p *PolymerClient) InfoAndTxsByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, types.Transactions, error) {
	var info eth.BlockInfo
	err := p.client.CallContext(ctx, &info, "ee_getInfoByHash", hash)
	return info, types.Transactions{}, err
}

// ----------------------------------------------------------
// TODO make the test happy for now
func (p *PolymerClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return big.NewInt(10000000), nil
}

func (p *PolymerClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}
