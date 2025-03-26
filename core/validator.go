package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{bc: bc}
}

func (bv *BlockValidator) ValidateBlock(b *Block) error {
	if bv.bc.HasBlock(b.Height) {
		return fmt.Errorf("block %d already exists", b.Height)
	}

	if b.Height != bv.bc.Height()+1 {
		return fmt.Errorf("height %d is greater than blockchain height %d", b.Height, bv.bc.Height())
	}

	prevHeader, err := bv.bc.GetHeader(b.Height - 1)
	if err != nil {
		return err
	}

	hash := BlockHasher{}.Hash(prevHeader)
	if hash != b.PrevBlockHash {
		return fmt.Errorf("invalid previous block hash")
	}

	if err := b.Verify(); err != nil {
		return err
	}

	return nil
}
