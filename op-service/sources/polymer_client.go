package sources

import (
	"context"
	"fmt"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Hash []byte

type Transaction struct {
	inner []byte
	hash  []byte
}

func (tx *Transaction) Type() uint8 {
	return types.DynamicFeeTxType
}

func (tx *Transaction) MarshalBinary() ([]byte, error) {
	return tx.inner, nil
}

func (tx *Transaction) Data() []byte {
	return tx.inner
}

func (tx *Transaction) Hash() common.Hash {
	return common.BytesToHash(tx.hash)
}

type Transactions []*Transaction

func (t Transactions) Len() int {
	return t.Len()
}

type Block struct {
	height uint64
	hash   Hash
	parent Hash
	time   uint64
	txs    Transactions
	data   []byte
}

func (m *Block) Height() uint64 {
	return m.height
}

func (m *Block) NumberU64() uint64 {
	return m.height
}

func (m *Block) Bytes() []byte {
	return []byte(fmt.Sprintf("%v %v %v", m.hash, m.parent, m.height))
}

func (m *Block) Hash() common.Hash {
	return common.BytesToHash(m.hash)

}

func (m *Block) ParentHash() common.Hash {
	return common.BytesToHash(m.parent)
}

func (m *Block) Time() uint64 {
	//TODO implement me
	panic("implement me")
}

func (m *Block) Transactions() Transactions {
	//TODO implement me
	panic("implement me")
}

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

func (p *PolymerClient) BlockByNumber(ctx context.Context, number uint64) (*Block, error) {
	var block *Block
	err := p.client.CallContext(ctx, block, "peptide_getBlockByNumber", number)
	return block, err
}

func (p *PolymerClient) Close() {
	p.client.Close()
}
