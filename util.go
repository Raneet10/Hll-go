package hll

import (
	"hash/adler32"
	"math/bits"
)

func Hash32(data []byte) uint32 {
	return adler32.Checksum(data)
}

func LeadingZeroes32(e uint32) int {
	return bits.LeadingZeros32(e)
}
