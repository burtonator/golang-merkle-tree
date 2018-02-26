package merkletree

import (
	"crypto/sha256"
	"errors"
	"log"
	"math"
)

const SHA256_LEN = 256 / 8

type HashNodePair struct {

	left HashNode

	// Right will have a zero object with a zero length hashcode when we are
	// working with an odd number of records.
	right HashNode

}

// TODO: what if there is an ODD number of leaf nodes???  should we force them
// to be even?

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

		leafHashNode := LeafHashNode{&HashNode{hashcode, nil, nil}, slab}

		leafHashNodes = append(leafHashNodes, leafHashNode)

	}

	// TODO: now that we have the leaf nodes computed, with their hashcode, we need to
	// merge these until we have a root node, then we have the full merkle
	// hash tree.

}

func mergeForRoot(hashNodes []HashNode) HashNode {

	return mergeForIntermediate(hashNodes)[0];

}


func mergeForIntermediate(hashNodes []HashNode) []HashNode {

	merged := make([]HashNode, 0)
	//
	//if len(hashNodes) <= 1 {
	//	// we're already done here.
	//	return hashNodes
	//}
	//

	return merged

}

func computeHashNodePairs(hashNodes []HashNode) []HashNodePair {

	result := make([]HashNodePair, 0)

	queue := make(chan HashNode, len(hashNodes))

	// add all the hashNodes to the queue
	for _, hashNode := range hashNodes {
		queue <- hashNode
	}

	nrPairs := int(math.Ceil(float64(len(hashNodes)) / float64(2)))

	for i := 0; i < nrPairs; i++ {

		if left, err := takeHashNode(queue); err == nil {

			right, _ := takeHashNode(queue)

			hashNodePair := HashNodePair{left, right}

			result = append(result, hashNodePair)

		} else {
			log.Fatal("Could not find left node: ", err)
		}

	}

	return result

}

// TODO: make working with the queue into an object 
func takeHashNode(queue chan HashNode) (HashNode, error) {

	select {
	case msg := <-queue:
		return msg, nil
	default:
		return HashNode{}, errors.New("no value")
	}

}