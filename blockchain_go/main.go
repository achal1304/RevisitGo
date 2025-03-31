package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index        int
	PreviousHash string
	Hash         string
	Timestamp    string
	Data         string
}

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PreviousHash)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func createGenesisBlock() Block {
	genesis := Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Data:         "This is genesis block first",
		PreviousHash: "",
	}
	genesis.Hash = calculateHash(genesis)
	return genesis
}

func genesisBlock(prevBlock Block, data string) Block {
	newBlock := Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Data:         data,
		PreviousHash: prevBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func isBlockValid(prev, curr Block) bool {
	if curr.Index != prev.Index+1 {
		return false
	} else if curr.PreviousHash != prev.Hash {
		return false
	} else if curr.Hash != calculateHash(curr) {
		return false
	}
	return true
}

func main() {
	blockChain := []Block{createGenesisBlock()}

	blockChain = append(blockChain, genesisBlock(blockChain[len(blockChain)-1], "this is second block"))
	blockChain = append(blockChain, genesisBlock(blockChain[len(blockChain)-1], "this is third block"))

	for i := 1; i < len(blockChain); i++ {
		fmt.Printf("\n--- Block #%d ---\n", blockChain[i].Index)
		fmt.Println("Timestamp   :", blockChain[i].Timestamp)
		fmt.Println("Data        :", blockChain[i].Data)
		fmt.Println("Prev. Hash  :", blockChain[i].PreviousHash)
		fmt.Println("Hash        :", blockChain[i].Hash)
		fmt.Println("is blockchain valid ? ", isBlockValid(blockChain[i-1], blockChain[i]))
	}

	blockChain[2].Data = "data is tampered"

	for i := 1; i < len(blockChain); i++ {
		fmt.Printf("\n--- Block #%d ---\n", blockChain[i].Index)
		fmt.Println("Timestamp   :", blockChain[i].Timestamp)
		fmt.Println("Data        :", blockChain[i].Data)
		fmt.Println("Prev. Hash  :", blockChain[i].PreviousHash)
		fmt.Println("Hash        :", blockChain[i].Hash)
		fmt.Println("is blockchain valid ? ", isBlockValid(blockChain[i-1], blockChain[i]))
	}
}
