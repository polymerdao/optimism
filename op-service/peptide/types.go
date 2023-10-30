package peptide

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	bfttypes "github.com/cometbft/cometbft/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
)

type Data = hexutil.Bytes
type Hash = common.Hash

type Block struct {
	Txs             bfttypes.Txs    `json:"txs"`
	Header          *tmproto.Header `json:"header"`
	ParentBlockHash Hash            `json:"parentHash"`
	L1Txs           []Data          `json:"l1Txs"`
}

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

func (b *Block) Height() int64 {
	return b.Header.Height
}

func (b *Block) Bytes() []byte {
	bytes, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return bytes
}

// Hash returns a unique hash of the block, used as the block identifier
func (b *Block) Hash() Hash {
	data := append(b.Bytes(), b.ParentBlockHash[:]...)
	hash := sha256.Sum256(data)
	return hash
}

func (b *Block) ParentHash() Hash {
	return b.ParentBlockHash
}

func (b *Block) NumberU64() uint64 {
	return uint64(b.Height())
}

func (b *Block) Time() uint64 {
	return uint64(b.Header.Time.Unix())
}

func (b *Block) Transactions() types.Transactions {
	var txs types.Transactions
	for _, l1tx := range b.L1Txs {
		var tx types.Transaction
		if err := tx.UnmarshalBinary(l1tx); err != nil {
			panic("failed to unmarshal l2 txs")
		}
		txs = append(txs, &tx)
	}

	chainId, err := strconv.Atoi(b.Header.ChainID)
	if err != nil {
		panic(fmt.Sprintf("block chain id is not an integer (%s): %w", b.Header.ChainID, err))
	}

	for _, l2tx := range b.Txs {
		//TODO: update to use proper Gas and To values if possible
		txData := &types.DynamicFeeTx{
			ChainID: big.NewInt(int64(chainId)),
			Data:    l2tx,
			Gas:     0,
			Value:   big.NewInt(0),
			To:      nil,
		}
		tx := types.NewTx(txData)
		txs = append(txs, tx)
	}
	return txs
}
