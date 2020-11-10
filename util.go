package hll

import (
	"hash/adler32"
	"math/bits"
)

func Alpha(m uint32) float64 {

	if m == 16 {
		return 0.673
	} else if m == 32 {
		return 0.697
	} else if m == 64 {
		return 0.709
	}

	return 0.7213 / (1 + 1.079/float64(m))

}

//Returns 32-bit hash of data
func Hash32(data []byte) uint32 {
	return adler32.Checksum(data)
}

//Returns leading zeroes of a 32-bit uint
func LeadingZeroes32(e uint32) int {
	return bits.LeadingZeros32(e)
}

func Estimate(M []uint32) float64 {
	Z := 0.0

	for _, e := range M {
		Z += 1.0 / float64(uint32(1)<<e)
	}

	m := uint32(len(M))
	return Alpha(m) * float64(m) * float64(m) * Z
}

func ZeroValueRegisters(M []uint32) uint32 {
	zeroes := 0

	for _, e := range M {
		if e == 0 {
			zeroes++
		}
	}

	return uint32(zeroes)
}
