package assignment1bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Block struct {
	Nonce  int
	Data   string //transactions
	P_hash string //hash of previous block
	C_hash string //hash of current block
}

func concat(a int, b string, c string) string {
	var aa string
	aa = strconv.Itoa(a)
	aa = aa + b + c
	return aa
}

func NewBlock(transaction string, Nonce int, previousHash string) (bloc *Block) {
	new_block := Block{Nonce: Nonce, Data: transaction, P_hash: previousHash}
	hash := CalculateHash(concat(Nonce, previousHash, transaction))
	new_block.C_hash = hash
	bloc = &new_block
	return bloc
}

// Calculating Hash for block
func CalculateHash(stringToHash string) (output string) {
	output = fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
	return output
}

func Display_blocks(blockchain []Block) {
	for i := 0; i < len(blockchain); i++ {
		fmt.Printf("\n\nBlock: %d\nNonce: %d\nData: %v\nPrevious Hash: %v\nCurrent Hash: %v", i, blockchain[i].Nonce, blockchain[i].Data, blockchain[i].P_hash, blockchain[i].C_hash)
		fmt.Println()
	}
}

// Change block
func ChangeBlock(blockchain []Block) {
	var num int
	fmt.Println("Enter number of the block(0,1,2...)")
	fmt.Scan(&num)
	if num < 0 || num > len(blockchain) {
		fmt.Println("Invalid number")
		return
	}
	var transaction string
	fmt.Println("Enter new Transaction(No spaces)")
	fmt.Scan(&transaction)
	blockchain[num].Data = transaction
	make_hash := concat(blockchain[num].Nonce, blockchain[num].P_hash, transaction)
	hash := CalculateHash(make_hash)
	blockchain[num].C_hash = hash
}

// verification
func VerifyChain(blockchain []Block) {
	if len(blockchain) == 0 {
		fmt.Println("Blockchain empty")
	} else {
		for i := 0; i < (len(blockchain) - 1); i++ {
			if blockchain[i].C_hash != blockchain[i+1].P_hash {
				fmt.Printf("Block %d changed", i)
				fmt.Println()
				return
			}
		}
		fmt.Println("Verification successful. Blockchain is unchanged")
	}

}
