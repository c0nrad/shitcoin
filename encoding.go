package main

import (
	"bytes"
	"encoding/hex"
	"math/big"
)

func Reverse(in []byte) (out []byte) {
	out = make([]byte, len(in))
	for i, c := range in {
		out[len(in)-1-i] = c
	}
	return out
}

func ToHex(in []byte) (out []byte) {
	length := hex.EncodedLen(len(in))
	out = make([]byte, length)
	hex.Encode(out, in)
	return out
}

func FromHex(in []byte) (out []byte) {
	length := hex.DecodedLen(len(in))
	out = make([]byte, length)
	hex.Decode(out, in)
	return out
}

var B58 = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func ToBase58(in []byte) []byte {
	var n, mod big.Int
	var results []byte
	n.SetBytes(in)

	for n.Cmp(big.NewInt(0)) > 0 {
		mod.Mod(&n, big.NewInt(58))
		results = append(results, B58[mod.Int64()])
		n.Div(&n, big.NewInt(58))
	}

	return Reverse(results)
}

func FromBase58(in []byte) []byte {
	var results big.Int
	results.SetInt64(0)

	for _, c := range in {
		results.Mul(&results, big.NewInt(58))
		index := int64(bytes.Index(B58, []byte{c}))
		if index <= -1 {
			panic("Not a valid base58 string")
		}
		results.Add(&results, big.NewInt(index))
	}
	return results.Bytes()
}
