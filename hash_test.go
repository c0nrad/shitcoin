package main

import "testing"
import "bytes"

func TestHash(t *testing.T) {
	in := []byte("Hello, world!0")
	out := FromHex([]byte("1312af178c253f84028d480a6adc1e25e81caa44c749ec81976192e2ec934c64"))
	if !bytes.Equal(out, Hash(in)) {
		t.Errorf("Failed to correctly take hash %s, %s", in, out, Hash(in))
	}
}

func TestHex(t *testing.T) {
	in := []byte("I AM SOME DATER")
	if !bytes.Equal(in, FromHex(ToHex(in))) {
		t.Error("Failed to correctly convert string to and from hex")
	}
}

func TestPairHash(t *testing.T) {
	p1 := []byte("00baf6626abc2df808da36a518c69f09b0d2ed0a79421ccfde4f559d2e42128b")
	p2 := []byte("91c5e9f288437262f218c60f986e8bc10fb35ab3b9f6de477ff0eb554da89dea")
	out := []byte("2ec6aa182d2c2a79e697222784fd30dffaaaa3ba2adf04c3d660f5591ee6dc70")
	if !bytes.Equal(PairHash(p1, p2), out) {
		t.Error("Failed to correctly take pair hash")
	}
}

func TestRipeMD160(t *testing.T) {
	in := []byte("The quick brown fox jumps over the lazy dog")
	sln := []byte("37f332f68db77bd9d7edd4969571ad671cf9dd3b")
	if !bytes.Equal(RIPEMD160(in), FromHex(sln)) {
		t.Error("Failed to correctly take ripe hash")
	}
}

// RIPEMD-160("The quick brown fox jumps over the lazy dog") =
// 37f332f68db77bd9d7edd4969571ad671cf9dd3b
