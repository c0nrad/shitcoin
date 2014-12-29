package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
)

type Block struct {
	Version           uint32
	PreviousBlockHash []byte
	MerkleRoot        []byte
	Timestamp         uint32
	Bits              uint32
	Nonce             uint32
}

const MaxUint = ^uint32(0)

// It's possible that only adjusting the Nonce isn't enough
// Also, I start at the blocks previous nonce value...
func (b *Block) Mine() bool {
	nonce := uint32(b.Nonce)
	for nonce = 0; nonce < MaxUint; nonce++ {
		b.Nonce = nonce
		fmt.Println(b.Nonce)

		if b.IsValid() {
			return true
		}
	}
	return false
}

func (b *Block) IsValid() bool {
	hashString := b.Hash()

	var hash, diff big.Int
	hash.SetBytes(FromHex(hashString))
	diff = b.Difficulty()

	if hash.Cmp(&diff) < 0 {
		return true
	}

	return false
}

func (b *Block) Hash() (out []byte) {
	out = Hash(Hash(b.Header()))
	out = ToHex(Reverse(out))
	return
}

func (b *Block) Header() []byte {
	buff := new(bytes.Buffer)

	binary.Write(buff, binary.LittleEndian, b.Version)
	buff.Write(Reverse(FromHex(b.PreviousBlockHash)))
	buff.Write(Reverse(FromHex(b.MerkleRoot)))
	binary.Write(buff, binary.LittleEndian, b.Timestamp)
	binary.Write(buff, binary.LittleEndian, b.Bits)
	binary.Write(buff, binary.LittleEndian, b.Nonce)

	return buff.Bytes()
}

func (b *Block) Difficulty() (diff big.Int) {
	exponent := b.Bits >> 24     //upper byte
	mantisa := b.Bits & 0xffffff //lower three bytes

	diff.Exp(big.NewInt(2), big.NewInt(8*(int64(exponent)-3)), nil)
	diff.Mul(&diff, big.NewInt(int64(mantisa)))
	return diff
}
