package merkletree

import (
	"os"
)

// Spit takes a file path, opens the file, then returns an array of slabs
// representing the underlying data regions.
func SplitFile(path string, length int64) ([]Slab, error) {

	file, _ := os.Open(path)

	// now stream through the file reading in the slabs
	
	if stat, err := file.Stat(); err == nil {

		var fileSize = stat.Size()

		slabReferences := ComputeSlabReferences(fileSize, length)

		result := make([]Slab, 0)

		for _, slabReference := range slabReferences {

			data := make([]byte, slabReference.length)

			if _, err := file.ReadAt(data, slabReference.offset); err == nil {

				slab := Slab{&slabReference, data}

				result = append(result, slab)
				
			} else {
				return nil, err
			}

		}

		// now read the regions from the files specified by the slab references

		return result, nil

	} else {
		return nil, err
	}

}
