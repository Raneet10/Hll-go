package hll

import (
	"hash/adler32"
	"math/bits"
)

//Returns 32-bit hash of data
func Hash32(data []byte) uint32 {
	return adler32.Checksum(data)
}

//Returns leading zeroes of a 32-bit uint
func LeadingZeroes32(e uint32) int {
	return bits.LeadingZeros32(e)
}
