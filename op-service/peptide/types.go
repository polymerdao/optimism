package peptide

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

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

type Transactions []*Transaction

func (t Transactions) Len() int {
	return t.Len()
}

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

func (tx *Transaction) Len() int {
	return len(tx.inner)
}

func (tx *Transaction) Hash() common.Hash {
	return common.BytesToHash(tx.hash)
}
