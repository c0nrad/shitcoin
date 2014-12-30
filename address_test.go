package main

import (
	"bytes"
	"math/big"
	"testing"
)

func TestPrivateWIF(t *testing.T) {
	var D big.Int
	D.SetString("0C28FCA386C7A227600B2FE50B7CAE11EC86D3BF1FBE471BE89827E19D72AA1D", 16)
	WIF := []byte("5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ")
	a := Address{*big.NewInt(0), *big.NewInt(0), D}

	if !bytes.Equal(a.PrivateWIF(), WIF) {
		t.Error("Failed to correctly generate PrivateWIF from key")
	}
}

func TestWIFToPrivate(t *testing.T) {
	wif := []byte("5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ")
	private := FromHex([]byte("0C28FCA386C7A227600B2FE50B7CAE11EC86D3BF1FBE471BE89827E19D72AA1D"))

	if !bytes.Equal(WIFToPrivateKey(wif), private) {
		t.Error("Failed to convert wif to private key")
	}
}

func TestVerifyWIFChecksum(t *testing.T) {
	wif := []byte("5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ")

	if !VerifyWIFChecksum(wif) {
		t.Error("Failed to verify valid WIF checksum")
	}

	badWIF := []byte("1337CGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ")
	if VerifyWIFChecksum(badWIF) {
		t.Error("Verified invalid WIF checksum")
	}
}
