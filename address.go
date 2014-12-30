package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

type Address struct {
	X, Y big.Int

	D big.Int
}

func NewAddress() (a Address) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	a.X = *priv.PublicKey.X
	a.Y = *priv.PublicKey.Y
	a.D = *priv.D
	return a
}

func VerifyWIFChecksum(wif []byte) bool {
	address := FromBase58(wif)
	length := len(address)

	checksum := address[length-4 : length]

	address = address[0 : length-4]
	guessChecksum := Hash(Hash(address))[0:4]

	return bytes.Equal(checksum, guessChecksum)
}

func WIFToPrivateKey(wif []byte) []byte {
	if !VerifyWIFChecksum(wif) {
		panic("Failed to cerify checksum" + string(wif))
	}

	address := FromBase58(wif)
	length := len(address)

	// Drop the 0x80 for mainnet
	// Drop the 4byte checksum
	address = address[1 : length-4]
	return address
}

func (a *Address) PrivateWIF() []byte {
	privateKey := a.D
	address := privateKey.Bytes()

	//Add a 0x80 byte in front of it for mainnet addresses or 0xef for testnet addresses.
	//Also add a 0x01 byte at the end if the private key will correspond to a compressed public key
	address = append([]byte{0x80}, address...)

	//Calculate Checksum
	checksum := Hash(Hash(address))[0:4]

	//Add the 4 checksum bytes to extended key
	address = append(address, checksum...)

	//Return base58 encoded
	return ToBase58(address)
}
