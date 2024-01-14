package core

// const dbFile = "blockchain.db"
// const blocksBucket = "blocks"

// type BlockchainIterator struct {
// 	currentHash []byte
// 	Db          *bolt.DB
// }

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	NewBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, NewBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
