package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
}

// Blockchain is a series of validated Blocks
var Blockchain []Block

// CalculateHash returns the hash of a Block
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// NewBlock creates a new Block using previous block's hash
func NewBlock(data string, oldBlock Block) Block {
	var newBlock Block

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

// DisplayAllBlocks prints all blocks in the blockchain
func DisplayAllBlocks() {
	for _, block := range Blockchain {
		println("Index:", block.Index)
		println("Timestamp:", block.Timestamp)
		println("Data:", block.Data)
		println("Hash:", block.Hash)
		println("PrevHash:", block.PrevHash)
		println()
	}
}

// ModifyBlock allows modification of a block at a given index
func ModifyBlock(newData string, blockIndex int) {
	if blockIndex < 0 || blockIndex >= len(Blockchain) {
		println("No block found at the given index")
		return
	}

	Blockchain[blockIndex].Data = newData
	Blockchain[blockIndex].Hash = calculateHash(Blockchain[blockIndex])

	// To maintain chain integrity, re-hash the subsequent blocks
	for i := blockIndex + 1; i < len(Blockchain); i++ {
		Blockchain[i].PrevHash = Blockchain[i-1].Hash
		Blockchain[i].Hash = calculateHash(Blockchain[i])
	}
}
