package dspinner

import (
	"github.com/ipfs/go-cid"
)

var sealedMap map[cid.Cid]map[cid.Cid][]byte

func init() {
	sealedMap = make(map[cid.Cid]map[cid.Cid][]byte)
}

func startSeal(root cid.Cid, value []byte) error {
	sealedMap[root] = make(map[cid.Cid][]byte)
	sealedMap[root][root] = append([]byte("crust/"), value...)
	return nil
}

func seal(root cid.Cid, leaf cid.Cid, value []byte) error {
	sealedMap[root][leaf] = append([]byte("crust/"), value...)
	return nil
}

func endSeal(root cid.Cid) (map[cid.Cid][]byte, error) {
	resMap := sealedMap[root]
	delete(sealedMap, root)
	return resMap, nil
}
