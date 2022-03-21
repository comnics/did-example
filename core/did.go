package core

import (
	"fmt"
	"github.com/comnics/did-example/util"
)

type DID struct {
	did    string
	method string
}

func NewDID(method string, pbKey string) (did *DID) {
	var newDid = new(DID)
	newDid.method = method
	specificIdentifier := util.MakeHashBase58(pbKey)
	newDid.did = fmt.Sprintf("did:%s:%s", method, specificIdentifier)

	return newDid
}

func (d *DID) String() string {
	return d.did
}
