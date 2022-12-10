package main

import (
	"fmt"
	"github.com/comnics/did-example/util"
)

func main() {
	str := "Hello~ SSIKR."
	hash1 := util.MakeHash(str)
	hash2 := util.MakeHashBase58(str)
	hash3 := util.MakeHashHex(str)

	fmt.Println("")
	fmt.Println("Plain Text: ", str)
	fmt.Println("MakeHash: ", hash1)
	fmt.Println("HashBase58: ", hash2)
	fmt.Println("MakeHashHex: ", hash3)
}
