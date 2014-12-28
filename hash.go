package main

import "encoding/hex"
import "crypto/sha256"

func Hash(in []byte) (out []byte) {
	h := sha256.New()
	h.Write(in)
	return h.Sum(nil)
}

func PairHash(p1 []byte, p2 []byte) (out []byte) {
	p1, p2 = Reverse(FromHex(p1)), Reverse(FromHex(p2))
	hashPair := Hash(append(p1, p2...))
	out = Hash(hashPair)
	return ToHex(Reverse(out))
}

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
