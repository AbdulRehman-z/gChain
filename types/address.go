package types

import (
	"encoding/hex"
	"fmt"
)

type Address [20]uint8

func (addr Address) String() string {
	return hex.EncodeToString(addr[:])
}

func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		err := fmt.Sprintf("address lenght %d should be 32", len(b))
		panic(err)
	}

	var addr Address
	for i := range b {
		addr[i] = b[i]
	}

	return Address(addr)
}
