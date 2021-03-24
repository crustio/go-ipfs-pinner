package dspinner

import (
	r "math/rand"
	"time"

	"github.com/ipfs/go-cid"
)

var sealedMap map[cid.Cid]map[cid.Cid]string
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func startSeal(root cid.Cid, value []byte) error {
	sealedMap[root] = make(map[cid.Cid]string)
	r.Seed(time.Now().UnixNano())
	sealedMap[root][root] = "crust/" + randStringRunes(10)
	return nil
}

func seal(root cid.Cid, leaf cid.Cid, value []byte) error {
	sealedMap[root][leaf] = "crust/" + randStringRunes(10)
	return nil
}

func endSeal(root cid.Cid) (map[cid.Cid]string, error) {
	resMap := sealedMap[root]
	delete(sealedMap, root)
	return resMap, nil
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}
