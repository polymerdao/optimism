package sources

import (
	"context"
	"fmt"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/peptide"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

// PolymerClient is a L2ClientGeneric implementation that interacts with the Polymer's ABCI app as an L2 Execution
// Engine via RPC bindings.
type PolymerClient struct {
	client client.RPC
}

func (p PolymerClient) ForkchoiceUpdate(ctx context.Context, state *eth.ForkchoiceState, attr *eth.PayloadAttributes) (*eth.ForkchoiceUpdatedResult, error) {
	//TODO implement me
	panic("implement me")
}

func (p PolymerClient) NewPayload(ctx context.Context, payload *eth.ExecutionPayload) (*eth.PayloadStatusV1, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PolymerClient) GetPayload(ctx context.Context, payloadId eth.PayloadID) (*eth.ExecutionPayload, error) {

	//TODO implement me
	panic("implement me")
}

func (p *PolymerClient) L2BlockRefByHash(ctx context.Context, l2Hash common.Hash) (eth.L2BlockRef, error) {
	var blockRef eth.L2BlockRef
	err := p.client.CallContext(ctx, &blockRef, "ee_getL2BlockRefByHash", l2Hash)
	return blockRef, err
}

func (p *PolymerClient) SystemConfigByL2Hash(ctx context.Context, hash common.Hash) (eth.SystemConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PolymerClient) L2BlockRefByNumber(ctx context.Context, num uint64) (eth.L2BlockRef, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PolymerClient) OutputV0AtBlock(ctx context.Context, blockHash common.Hash) (*eth.OutputV0, error) {
	//TODO implement me
	panic("implement me")
}

func NewPolymerClient(client client.RPC) *PolymerClient {
	return &PolymerClient{client: client}
}

func (p *PolymerClient) payloadCall(ctx context.Context, method string, id rpcBlockID) (*eth.ExecutionPayload, error) {
	var block *rpcBlock
	err := p.client.CallContext(ctx, &block, method, id.Arg(), true)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, ethereum.NotFound
	}
	payload, err := block.ExecutionPayload(true)
	if err != nil {
		return nil, err
	}
	if err := id.CheckID(payload.ID()); err != nil {
		return nil, fmt.Errorf("fetched payload does not match requested ID: %w", err)
	}
	return payload, nil
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

func (p *PolymerClient) PayloadByLabel(ctx context.Context, label eth.BlockLabel) (*eth.ExecutionPayload, error) {
	var payload *eth.ExecutionPayload
	err := p.client.CallContext(ctx, &payload, "ee_getPayloadByLabel", label)
	return payload, err
	//return p.payloadCall(ctx, "eth_getBlockByNumber", label)
}

func (p *PolymerClient) PayloadByNumber(ctx context.Context, num uint64) (*eth.ExecutionPayload, error) {
	var payload *eth.ExecutionPayload
	err := p.client.CallContext(ctx, &payload, "ee_getPayloadByNumber", num)
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

func (p *PolymerClient) InfoAndTxsByHash(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Transactions, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PolymerClient) Close() {
	p.client.Close()
}
