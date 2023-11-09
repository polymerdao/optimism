package eth

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"math/big"
	"strconv"

	bfttypes "github.com/cometbft/cometbft/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type Bytes []byte

type Header struct {
	ChainID string `json:"chain_id"`
	Height  int64  `json:"height"`
	Time    uint64 `json:"time"`

	// prev block hash
	LastBlockHash []byte `json:"last_block_hash"`

	// hashes of block data
	LastCommitHash Bytes `json:"last_commit_hash"` // commit from validators from the last block
	DataHash       Bytes `json:"data_hash"`        // transactions

	// hashes from the app output from the prev block
	ValidatorsHash     Bytes `json:"validators_hash"`      // validators for the current block
	NextValidatorsHash Bytes `json:"next_validators_hash"` // validators for the next block
	ConsensusHash      Bytes `json:"consensus_hash"`       // consensus params for current block
	// root hash of all results from the txs from the previous block
	LastResultsHash Bytes `json:"last_results_hash"`

	// consensus info
	EvidenceHash Bytes `json:"evidence_hash"` // evidence included in the block
}

func (h *Header) Populate(cosmosHeader *tmproto.Header) *Header {
	h.ChainID = cosmosHeader.ChainID
	h.Height = cosmosHeader.Height
	h.Time = uint64(cosmosHeader.Time.Unix())
	h.LastBlockHash = cosmosHeader.LastBlockId.Hash
	h.LastCommitHash = cosmosHeader.LastCommitHash
	h.DataHash = cosmosHeader.DataHash
	h.ValidatorsHash = cosmosHeader.ValidatorsHash
	h.NextValidatorsHash = cosmosHeader.NextValidatorsHash
	h.ConsensusHash = cosmosHeader.ConsensusHash
	h.LastResultsHash = cosmosHeader.LastResultsHash
	h.EvidenceHash = cosmosHeader.EvidenceHash
	return h
}

// fixed size hash of 32 bytes generated by SHA256 or Keccak256
// type Hash = common.Hash
type Address = common.Address

type Block struct {
	Txs             bfttypes.Txs    `json:"txs"`
	Header          *Header         `json:"header"`
	ParentBlockHash common.Hash     `json:"parentHash"`
	L1Txs           []Data          `json:"l1Txs"`
	GasLimit        *hexutil.Uint64 `json:"gasLimit"`
}

var _ BlockData = (*Block)(nil)

func BlockUnmarshaler(bytes []byte) (BlockData, error) {
	b := Block{}
	if err := json.Unmarshal(bytes, &b); err != nil {
		panic(err)
	}
	return &b, nil
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

// ExecutionPayload is an ethereum type with data that our block is missing.
// It seems like the op-node is happy accepting these fields since it won't look into the
// missing ones.
func (b *Block) ToExecutionPayload() (*ExecutionPayload, error) {
	return &ExecutionPayload{
		BlockHash:    b.Hash(),
		BlockNumber:  hexutil.Uint64(b.Height()),
		ParentHash:   b.ParentHash(),
		Timestamp:    hexutil.Uint64(b.Time()),
		Transactions: b.L1Txs,
		GasLimit:     *b.GasLimit,
	}, nil
}

func (b *Block) Time() uint64 {
	return b.Header.Time
}

func (b *Block) NumberU64() uint64 {
	return uint64(b.Height())
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

func (b *Block) Populate(block EthBlock) {
	// TODO: use the interface to populate the block
	blockCoerced := block.(*Block)
	b.Txs = blockCoerced.Txs
	b.Header = blockCoerced.Header
	b.ParentBlockHash = blockCoerced.ParentBlockHash
	b.L1Txs = blockCoerced.L1Txs
	b.GasLimit = blockCoerced.GasLimit
}

// This is only a representation of the block from the block-store point of view
// The interface is missing key accessors to fields within the block but that are not needed
// by the block store.
type BlockData interface {
	Height() int64
	Bytes() []byte
	Hash() Hash
	ParentHash() Hash
}
