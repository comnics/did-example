package core

import (
	"fmt"
	"github.com/comnics/did-example/util"
)

type DIDManager struct {
	did    string
	method string
}

func (d *DIDManager) MakeDID(method string, pbKey string) {
	d.method = method
	specificIdentifier := util.MakeHashBase58(pbKey)
	d.did = fmt.Sprintf("did:%s:%s", d.method, specificIdentifier)
}

func (d *DIDManager) String() string {
	return d.did
}
