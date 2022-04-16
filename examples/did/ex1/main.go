package main

import "fmt"

func main() {
	method := "comnic"
	specificIdentifier := "abcd1234"

	did := fmt.Sprintf("did:%s:%s", method, specificIdentifier)

	fmt.Printf("DID: %s\n", did)
}
