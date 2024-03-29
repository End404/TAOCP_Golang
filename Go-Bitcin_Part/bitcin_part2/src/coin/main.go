package main

import (
	"bitcin_part2/src/core"
	"fmt"
	"strconv"
)

func main() {
	bc := core.NewBlockchain() //初始化

	bc.AddBlock("Send 1 BTC to Ivan")      //加入一个区块，发送一个比特币
	bc.AddBlock("Send 2 more BTC to Ivan") //加入，发送更多

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash) //前一个区块的哈希值
		fmt.Printf("Data: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := core.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
