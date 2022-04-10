package main

import (
	"fmt"
	core "github.com/comnics/did-example/core"
)

func main() {
	did, _ := core.NewDID("test", "12345")

	fmt.Printf("DID: [%s]", did.String())

}
