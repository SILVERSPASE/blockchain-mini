package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/silverspase/blockchain/types"
)

var Blockchain []types.Block

func calculateHash(block types.Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock types.Block, BPM int) (types.Block, error) {
	var newBlock types.Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func IsBlockValid(newBlock, oldBlock types.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func ReplaceChain(newBlocks []types.Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
