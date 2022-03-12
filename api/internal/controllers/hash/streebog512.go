package hash

import (
	"github.com/pedroalbanese/gogost/gost34112012512"
)

func hashWithStreebog512(data []byte) []byte {
	hasher := gost34112012512.New()
	hasher.Write(data)

	rawHash := hasher.Sum(nil)
	return rawHash[:]
}
