package sources

import (
	"context"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/peptide"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// PolymerClient is a L2ClientGeneric implementation that interacts with the Polymer's ABCI app as an L2 Execution
// Engine via RPC bindings.
type PolymerClient struct {
	client client.RPC
}

func NewPolymerClient(client client.RPC) *PolymerClient {
	return &PolymerClient{client: client}
}

func (p *PolymerClient) PayloadByLabel(ctx context.Context, label eth.BlockLabel) (*eth.ExecutionPayload, error) {
	var payload *eth.ExecutionPayload
	err := p.client.CallContext(ctx, &payload, "ee_getPayloadByLabel", label)
	return payload, err
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

func (p *PolymerClient) BlockByNumber(ctx context.Context, number uint64) (*peptide.Block, error) {
	var block *peptide.Block
	err := p.client.CallContext(ctx, block, "peptide_getBlockByNumber", number)
	return block, err
}

func (p *PolymerClient) Close() {
	p.client.Close()
}
