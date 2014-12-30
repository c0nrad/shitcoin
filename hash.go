package main

import "crypto/sha256"
import "golang.org/x/crypto/ripemd160"

func Hash(in []byte) []byte {
	h := sha256.New()
	h.Write(in)
	return h.Sum(nil)
}

func RIPEMD160(in []byte) []byte {
	h := ripemd160.New()
	h.Write(in)
	return h.Sum(nil)
}

func PairHash(p1 []byte, p2 []byte) (out []byte) {
	p1, p2 = Reverse(FromHex(p1)), Reverse(FromHex(p2))
	hashPair := Hash(append(p1, p2...))
	out = Hash(hashPair)
	return ToHex(Reverse(out))
}
