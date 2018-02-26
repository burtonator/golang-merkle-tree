package merkletree


type SlabReference struct {

	// Offset is the byte offset position in the file that this slab starts
	offset int64

	// Length holds the length of the struct. We parse the file based into
	// slabs and keep the offset and length. Technically the length is stored
	// redundantly since we are given the length as input in our split but this
	// allows us to work with slabs elsewhere and keep the length handy.
	length int64

}

type Slab struct {

	*SlabReference

	// Data contains an array of bytes for the underlying representation of
	// slab content
	data []byte

}

type HashReference struct {

	// Hashcode has the actual raw bytes used to encode the underlying data for
	// this hash node.
	hashcode []byte

}

type HashNode struct {

	*HashReference

	left *HashNode

	right *HashNode

}

// LeafNode is just like a normal HashNode but it has no left or right nodes
// (they are nil) and it includes an extra pointer to the Slab that was used
// to construct it.
type LeafHashNode struct {

	*HashNode

	slab Slab

}
