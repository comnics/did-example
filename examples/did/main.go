package main

import (
	"fmt"
	"github.com/comnics/did-example/core"
)

func main() {
	did := new(core.DIDManager)

	did.MakeDID("test", "2345")

	fmt.Printf("DID: [%s]", did.String())

}
