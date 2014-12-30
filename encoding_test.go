package main

import (
	"bytes"
	"testing"
)

func TestBase58(t *testing.T) {
	in := []byte("a lifetime of yesterdays")

	if !bytes.Equal(FromBase58(ToBase58(in)), in) {
		t.Error("Failed to correctly encode a string to and from base58")
	}
}

func TestBase58Encode(t *testing.T) {
	in := FromHex([]byte("800C28FCA386C7A227600B2FE50B7CAE11EC86D3BF1FBE471BE89827E19D72AA1D507A5B8D"))
	out := []byte("5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ")

	if !bytes.Equal(ToBase58(in), out) {
		t.Error("Failed to correctly encode a string to base58")
	}
}

func TestBase58Decoding(t *testing.T) {
	in := []byte("5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ")
	out := FromHex([]byte("800C28FCA386C7A227600B2FE50B7CAE11EC86D3BF1FBE471BE89827E19D72AA1D507A5B8D"))

	if !bytes.Equal(FromBase58(in), out) {
		t.Error("Failed to correctly decode a string from base58")
	}

}
