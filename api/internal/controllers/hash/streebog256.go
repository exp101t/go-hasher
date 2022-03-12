package hash

import (
	"github.com/pedroalbanese/gogost/gost34112012256"
)

func hashWithStreebog256(data []byte) []byte {
	hasher := gost34112012256.New()
	hasher.Write(data)

	rawHash := hasher.Sum(nil)
	return rawHash[:]
}
