package core

type Blockchain struct {
	headers   []*Header
	validator Validator
	store     Storage
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStore(),
	}

	bc.validator = NewBlockValidator(bc)
	if err := bc.addBlockWithoutValidation(genesis); err != nil {
		return nil, err
	}

	return bc, nil
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

// AddBlock adds a block to the blockchain after validating it.
func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	if err := bc.addBlockWithoutValidation(b); err != nil {
		return err
	}

	return nil
}

// HasBlock checks if a block with the given height exists in the blockchain.
func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

// [0,1,2,3,4] => len is 5
// [0,1,2,3,4] => len-1 is 4
func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

// addBlockWithoutValidation adds a genesis block(initial block) to the blockchain without validation. It is used for initializing the blockchain.
func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.headers = append(bc.headers, b.Header)
	return bc.store.Put(b)
}
