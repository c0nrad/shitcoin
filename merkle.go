package main

// Expect them to be in hex
func Merkle(hashes [][]byte) []byte {
	if len(hashes) == 1 {
		return hashes[0]
	}

	var tree [][]byte

	for i := 0; i < len(hashes)-1; i = i + 2 {
		pairHash := PairHash(hashes[i], hashes[i+1])
		tree = append(tree, pairHash)
	}

	//Hash tree is odd, double hash last element
	if len(hashes)%2 == 1 {
		lastHash := hashes[len(hashes)-1]
		pairHash := PairHash(lastHash, lastHash)
		tree = append(tree, pairHash)
	}
	return Merkle(tree)
}
