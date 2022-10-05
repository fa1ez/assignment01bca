package assignment1bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

var blockchain []block

type block struct {
	nonce  int
	data   string //transactions
	P_Hash string //hash of previous block
	C_Hash string //hash of current block
}

func concat(a int, b string, c string) string {
	var aa string
	aa = strconv.Itoa(a)
	aa = aa + b + c
	return aa
}

func NewBlock(transaction string, nonce int, previousHash string) (bloc *block) {
	new_block := block{nonce: nonce, data: transaction, P_Hash: previousHash}
	hash := CalculateHash(concat(nonce, previousHash, transaction))
	new_block.C_Hash = hash
	bloc = &new_block
	return bloc
}

// Calculating Hash for block
func CalculateHash(stringToHash string) (output string) {
	output = fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
	return output
}

func Display_blocks() {
	for i := 0; i < len(blockchain); i++ {
		fmt.Printf("\n\nBlock: %d\nNonce: %d\nData: %v\nPrevious Hash: %v\nCurrent Hash: %v", i, blockchain[i].nonce, blockchain[i].data, blockchain[i].P_Hash, blockchain[i].C_Hash)
		fmt.Println()
	}
}

// Change block
func ChangeBlock() {
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
	blockchain[num].data = transaction
	make_hash := concat(blockchain[num].nonce, blockchain[num].P_Hash, transaction)
	hash := CalculateHash(make_hash)
	blockchain[num].C_Hash = hash
}

// verification
func VerifyChain() {
	if len(blockchain) == 0 {
		fmt.Println("Blockchain empty")
	} else {
		for i := 0; i < (len(blockchain) - 1); i++ {
			if blockchain[i].C_Hash != blockchain[i+1].P_Hash {
				fmt.Printf("Block %d changed", i)
				fmt.Println()
				return
			}
		}
		fmt.Println("Verification successful. Blockchain is unchanged")
	}

}
