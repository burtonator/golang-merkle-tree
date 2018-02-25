package merkletree

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// computeSplits take the given fileSize and length and computes block offsets
// within the file to return SlabReferences. The SlabReferences can then be used
// to split the inputs directly.
func ComputeSlabReferences(fileSize int64, length int64) []SlabReference {

	var result = []SlabReference{}

	var offset = int64(0)

	for ; offset < fileSize; offset += length {

		var start = offset;
		var end = min(offset + length, fileSize);

		result = append(result, SlabReference{start, end - start})

	}

	return result

}