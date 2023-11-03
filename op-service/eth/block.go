package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Hash = common.Hash

type EthBlock interface {
	Hash() Hash
	ParentHash() Hash
	NumberU64() uint64
	Transactions() types.Transactions
	Time() uint64
}

type ExtendedEthBlock interface {
	EthBlock
	Height() uint64
	Bytes() []byte
}
