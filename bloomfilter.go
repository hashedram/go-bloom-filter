package bloomfilter

import (
	"hash"
)

type BF struct {
	max     int
	k       int
	h       hash.Hash32
	bitList []int
}
