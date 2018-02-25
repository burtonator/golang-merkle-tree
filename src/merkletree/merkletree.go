package merkletree

// Node implements a struct that
type Node struct {


}

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
//
//
//func main() {
//
//	fmt.Printf("Hello world\n")
//
//	var foo = []int{}
//
//	foo = append(foo, 1)
//
//	//
//	//_, err := os.Open("/home/burton/test1.dat")
//	//if err != nil {
//	//	// TODO: are exceptiosn printe  here?
//	//	log.Fatal(err)
//	//}
//
//	// now parse the files into slabs...
//
//	// now take these slabs and build hash node
//
//	// now build the tree
//
//}