package merkletree

import (
	"os"
	"log"
)

//import "os"
//
// Spit takes a file path, opens the file, then returns an array of slabs
// representing the underlying data regions.
func SplitFile(path string, length int64) []Slab {

	file, _ := os.Open("/home/burton/test1.dat")

	// now stream through the file reading in the slabs
	
	if stat, err := file.Stat(); err == nil {

		var fileSize = stat.Size()

		slabReferences := ComputeSlabReferences(fileSize, length)

		result := make([]Slab, 0)

		for _, slabReference := range slabReferences {

			data := make([]byte, slabReference.length)

			if _, err := file.ReadAt(data, slabReference.offset); err == nil {

				// FIXME: we need to include offset + length
				slab := Slab{data:data}

				result = append(result, slab)
			} else {
				// FIXME: return this...
				log.Fatal(err)
			}

		}

		// now read the regions from the files specified by the slab references

		return result

	} else {
		// FIXME: return this...
		log.Fatal(err)
	}

}
