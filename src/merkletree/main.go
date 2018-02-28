package merkletree

import (
	"fmt"
	"os"
	"log"
)

func main() {

	path := os.Args[1]

	fmt.Printf("Computing merkle tree for: %s\n", path)

	if slabs, err := SplitFile(path, 1024); err == nil {

		tree := BuildTree(slabs)

		// TODO: we could do a better job printing this out visually via ascii art
		// but that's beyond the scope of this simple proof of concept

		fmt.Printf("%#v\n", tree)

	} else {
		log.Fatal("Could not split file: ", err)
	}

}