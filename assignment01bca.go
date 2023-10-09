//Ashish Jumani 20i-0494 Assignment 1

package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

var Blockchain []Block

func NewBlock(transaction string, nonce int, previousHash string) (*Block, error) {
	if len(transaction) == 0 {
		return nil, errors.New("Transaction cannot be empty")
	}

	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		CurrentHash:  "",
	}
	block.CurrentHash = block.CreateHash()
	return block, nil
}

func (b *Block) CreateHash() string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func DisplayBlocks() {
	for _, block := range Blockchain {
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.CurrentHash)
		fmt.Println("---------------------")
	}
}

func ChangeBlock(block *Block, newTransaction string) error {
	if len(newTransaction) == 0 {
		return errors.New("New transaction cannot be empty")
	}

	block.Transaction = newTransaction
	block.CurrentHash = block.CreateHash()
	return nil
}

func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ {
		currentBlock := Blockchain[i]
		previousBlock := Blockchain[i-1]

		if currentBlock.PreviousHash != previousBlock.CurrentHash {
			return false
		}
	}
	return true
}

func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
}

func main() {
	// Example usage of the functions
	genesisBlock, err := NewBlock("Genesis Block", 0, "")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	Blockchain = append(Blockchain, *genesisBlock)

	newBlock, err := NewBlock("Alice to Bob", 1, Blockchain[len(Blockchain)-1].CurrentHash)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	Blockchain = append(Blockchain, *newBlock)

	DisplayBlocks()

	fmt.Println("Changing block transaction...")
	err = ChangeBlock(&Blockchain[1], "Charlie to David")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	DisplayBlocks()

	fmt.Println("Verifying blockchain...")
	isValid := VerifyChain()
	if isValid {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}
