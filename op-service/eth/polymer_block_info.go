package eth

import "github.com/ethereum/go-ethereum/common"

type PolymerBlockInfo struct {
	Hash       common.Hash `json:"hash"`
	Number     uint64      `json:"blockNumber"`
	ParentHash common.Hash `json:"parentHash"`
	Time       uint64      `json:"timestamp"`
	StateRoot  common.Hash `json:"stateRoot"`
}

func (p PolymerBlockInfo) Root() common.Hash {
	return p.StateRoot
}
