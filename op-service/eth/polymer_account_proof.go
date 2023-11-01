package eth

import (
	"github.com/ethereum/go-ethereum/common"
)

type PolymerAccountResult struct {
	StorageHash common.Hash `json:"storageHash"`
}

func (res *PolymerAccountResult) GetStorageHash() common.Hash {
	return res.StorageHash
}

func (res *PolymerAccountResult) Verify(stateRoot common.Hash) error {
	return nil
}
