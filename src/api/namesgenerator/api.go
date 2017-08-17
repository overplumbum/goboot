package namesgenerator

import (
	"fmt"
	"hash/fnv"
)

func NameForHash(s string) string {
	f := fnv.New64()
	f.Write([]byte(s))
	val := f.Sum64() % uint64(len(left)*len(right))

	return fmt.Sprintf("%s_%s", left[val%uint64(len(left))], right[val/uint64(len(left))])
}
