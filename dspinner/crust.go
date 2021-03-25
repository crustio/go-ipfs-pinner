package dspinner

import (
	"context"

	"github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
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

func sealBlock(root cid.Cid, leaf cid.Cid, value []byte) error {
	sealedMap[root][leaf] = append([]byte("crust/"), value...)
	return nil
}

func endSeal(root cid.Cid) (map[cid.Cid][]byte, error) {
	resMap := sealedMap[root]
	delete(sealedMap, root)
	return resMap, nil
}

func seal(ctx context.Context, root cid.Cid, serv ipld.DAGService) (map[cid.Cid][]byte, error) {
	rootNode, err := serv.Get(ctx, root)
	if err != nil {
		return nil, err
	}

	err = startSeal(rootNode.Cid(), rootNode.RawData())
	if err != nil {
		return nil, err
	}

	err = deepSeal(ctx, rootNode, serv)
	if err != nil {
		return nil, err
	}
	return endSeal(root)
}

func deepSeal(ctx context.Context, rootNode ipld.Node, serv ipld.DAGService) error {
	for i := 0; i < len(rootNode.Links()); i++ {
		leafNode, err := serv.Get(ctx, rootNode.Links()[i].Cid)
		if err != nil {
			return err
		}

		err = deepSeal(ctx, leafNode, serv)
		if err != nil {
			return err
		}

		err = sealBlock(rootNode.Cid(), leafNode.Cid(), leafNode.RawData())
		if err != nil {
			return err
		}
	}

	return nil
}
