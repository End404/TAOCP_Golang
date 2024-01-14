package core

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
