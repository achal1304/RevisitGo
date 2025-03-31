package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
	Nonce        int
}

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s%d", block.Index, block.Timestamp, block.Data, block.PreviousHash, block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func mineBlock(block *Block, difficulty int) {
	prefix := strings.Repeat("0", difficulty)
	for {
		block.Hash = calculateHash(*block)
		if strings.HasPrefix(block.Hash, prefix) {
			break
		}
		block.Nonce++
	}
}

func createGenesisBlock(difficulty int) Block {
	genesis := Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Data:         "Genesis Block",
		PreviousHash: "",
		Nonce:        0,
	}
	mineBlock(&genesis, difficulty)
	return genesis
}

func generateBlock(prev Block, data string, difficulty int) Block {
	newBlock := Block{
		Index:        prev.Index + 1,
		Timestamp:    time.Now().String(),
		Data:         data,
		PreviousHash: prev.Hash,
		Nonce:        0,
	}
	mineBlock(&newBlock, difficulty)
	return newBlock
}

func isBlockValid(newBlock, prevBlock Block, difficulty int) bool {
	if newBlock.Index != prevBlock.Index+1 {
		return false
	}
	if newBlock.PreviousHash != prevBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	prefix := strings.Repeat("0", difficulty)
	if !strings.HasPrefix(newBlock.Hash, prefix) {
		return false
	}
	return true
}

func main() {
	difficulty := 4
	blockchain := []Block{createGenesisBlock(difficulty)}

	block1 := generateBlock(blockchain[len(blockchain)-1], "Send 1 BTC to Alice", difficulty)
	block2 := generateBlock(blockchain[len(blockchain)-1], "Send 2 BTC to Bob", difficulty)

	blockchain = append(blockchain, block1, block2)

	for _, block := range blockchain {
		fmt.Printf("\n--- Block #%d ---\n", block.Index)
		fmt.Println("Timestamp   :", block.Timestamp)
		fmt.Println("Data        :", block.Data)
		fmt.Println("Nonce       :", block.Nonce)
		fmt.Println("Prev. Hash  :", block.PreviousHash)
		fmt.Println("Hash        :", block.Hash)
	}
}
