package crust

import (
	"encoding/json"
)

type SealedBlock struct {
	Hash string `json:"hash"`
	Size int    `json:"size"`
	Data []byte `json:"data"`
}

type SealedInfo struct {
	Sbs []SealedBlock `json:"sbs"`
}

func (si *SealedInfo) Bytes() []byte {
	bs, _ := json.Marshal(si)
	return bs
}

func TryGetSealedInfo(value []byte) (bool, *SealedInfo) {
	si := &SealedInfo{}
	err := json.Unmarshal(value, si)
	if err != nil {
		return false, nil
	}

	return true, si
}

func MergeSealedInfo(a *SealedInfo, b *SealedInfo) *SealedInfo {
	si := &SealedInfo{}
	si.Sbs = append(a.Sbs, b.Sbs...)
	return si
}
