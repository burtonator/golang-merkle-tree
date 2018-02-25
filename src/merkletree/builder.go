package merkletree

import "crypto/sha256"

const SHA256_LEN = 256 / 8

// BuildTree takes a set of Slabs and then builds a merkle hash tree from them
func BuildTree(slabs []Slab) {

	// first build the leaf nodes from the data slabs

	leafHashNodes := make([]LeafHashNode, 0)

	for _, slab := range slabs {

		hashcode := make([]byte, SHA256_LEN)

		// compute the hash from the underlying bytes
		hash := sha256.New()
		hash.Write(slab.data)
		hashcode = hash.Sum(hashcode)

		leafHashNode := LeafHashNode{&HashNode{&HashReference{hashcode}, nil, nil}, slab}

		leafHashNodes = append(leafHashNodes, leafHashNode)

	}

	// TODO: now that we have the leaf nodes computed, with their hashcode, we need to
	// merge these until we have a root node, then we have the full merkle
	// hash tree.

}

// Merge a given array of HashNodes by compressing the two nodes into a parent
// node and then returning that directly.  
func merge(hashNodes []HashNode) {

}