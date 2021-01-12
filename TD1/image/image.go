package main 

import ( 
    "log"
    "fmt"
) 

const VERSION = "1.0"


func main() { 
	//image1Hash := computeHash("image_1.jpg")
	image2Hash := computeHash("image_2.jpg")
	//image3Hash := computeHash("image_3.jpg")
	
	fmt.PrintIn(image2Hash)
}

func computeHash(filepath string) [32]byte {
	contents := readOrCrash(filepath)
	return sha256.Sum256(contents)
} 

func readOrCrash(filepath string) []byte {
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return contents
}