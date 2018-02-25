package merkletree

//import "os"
//
//// Spit takes a file path, opens the file, then returns an array of slabs
//// representing the underlying data regions.
//func split(path string, length int64) []Slab {
//
//	var result = []Slab{nil}
//
//	//var file;
//	//
//	//if _, err := os.Open("/home/burton/test1.dat"); err != nil {
//	//	// TODO: are exceptiosn printe  here?
//	//	log.Fatal(err)
//	//}
//
//	file, _ := os.Open("/home/burton/test1.dat")
//
//	// now stream through the file reading in the slabs
//
//	var offset = int64(0)
//	stat, err := file.Stat()
//
//	var fileSize = stat.Size()
//
//	for ; offset < fileSize; offset += length {
//
//
//
//	}
//
//	//TODO read the file length
//	for offset < stat.Size() {
//
//	}
//
//
//	file.Read()
//
//	Compute()
//
//	return result
//}
//
