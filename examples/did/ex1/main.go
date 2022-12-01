package main

import "fmt"

func main() {
	method := "ssikr"
	specificIdentifier := "abcd1234"

	did := fmt.Sprintf("did:%s:%s", method, specificIdentifier)

	fmt.Printf("DID: %s\n", did)
}
