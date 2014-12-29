package main

import (
	"bytes"
	"math/big"
	"testing"
)

func SampleBlock() (b Block) {
	b.Version = 2
	b.PreviousBlockHash = []byte("000000000000000117c80378b8da0e33559b5997f2ad55e2f7d18ec1975b9717")
	b.MerkleRoot = []byte("871714dcbae6c8193a2bb9b2a69fe1c0440399f38d94b3a0f1b447275a29978a")
	b.Timestamp = 0x53058b35
	b.Bits = 0x19015f53
	b.Nonce = 0
	return b
}

func TestBlockHash(t *testing.T) {
	b := SampleBlock()

	hash0 := []byte("5c56c2883435b38aeba0e69fb2e0e3db3b22448d3e17b903d774dd5650796f76")

	if !bytes.Equal(b.Hash(), hash0) {
		t.Error("Failed to correctly take hash of block /w Nonce=0")
	}

	b.Nonce = 856192328
	hash856192328 := []byte("0000000000000000e067a478024addfecdc93628978aa52d91fabd4292982a50")
	if !bytes.Equal(b.Hash(), hash856192328) {
		t.Error("Failed to correctly take hash of block /w Nonce=856192328")
	}
}

func TestIsValidNonce(t *testing.T) {
	b := SampleBlock()

	b.Nonce = 0
	if b.IsValid() {
		t.Error("Incorrectly said bad nonce as valid /w Nonce=0")
	}

	b.Nonce = 856192327
	if b.IsValid() {
		t.Error("Incorrectly said bad nonce as valid /w Nonce=856192327")
	}

	b.Nonce = 856192328
	if !b.IsValid() {
		t.Error("Didn't recognize valid block")
	}
}

func TestDifficulty(t *testing.T) {
	var sln big.Int
	sln.SetString("00000000000404CB000000000000000000000000000000000000000000000000", 16)

	b := SampleBlock()
	b.Bits = 0x1b0404cb
	diff := b.Difficulty()

	if sln.Cmp(&diff) != 0 {
		t.Error("Failed to correctly calculate difficultly")
	}
}

func TestBlockHeader(t *testing.T) {
	header := []byte{2, 0, 0, 0, 23, 151, 91, 151, 193, 142, 209, 247, 226, 85, 173, 242, 151, 89, 155, 85, 51, 14, 218, 184, 120, 3, 200, 23, 1, 0, 0, 0, 0, 0, 0, 0, 138, 151, 41, 90, 39, 71, 180, 241, 160, 179, 148, 141, 243, 153, 3, 68, 192, 225, 159, 166, 178, 185, 43, 58, 25, 200, 230, 186, 220, 20, 23, 135, 53, 139, 5, 83, 83, 95, 1, 25, 0, 0, 0, 0}

	b := SampleBlock()

	if !bytes.Equal(b.Header(), header) {
		t.Error("Failed to correctly generate header")
	}
}

func TestMine(t *testing.T) {
	b := SampleBlock()
	b.Nonce = 856191328
	if !b.Mine() {
		t.Error("Failed to mine block")
	}

	if b.Nonce != 856192328 {
		t.Error("mined block with inncorect nonce")
	}
}
