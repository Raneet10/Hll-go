package hll

import (
	"math"
)

//const TWO_TO_32 = 1 << 32

//Hyperloglog struct consists of a register M of size 'm'

type Hll struct {
	M []uint32
	m uint32
}

//Creates a new Hyperloglog struct

func NewHll(m uint32) *Hll {
	M := make([]uint32, m)
	return &Hll{
		M: M,
		m: m,
	}
}

//Operation to add an elemnt in the register
func (h *Hll) AddElement(element []byte) {

	/*
		ALGORITHM:
		1. x := hash32(v)
		2. j := 1 + <x1 x2..xb>
		3. w := xb+1 xb2 ....
		4. M[j] := MAX(M[j], rho(w)), rho(w) = leftmostone(w)
	*/
	b := uint32(math.Log2(float64(h.m)))
	x := Hash32(element)
	j := x >> (32 - b)
	lmo := uint32(LeadingZeroes32((x << b))) + 1

	if lmo > h.M[j] {
		h.M[j] = lmo
	}
}

//TODO : Count

//Merge operation for two Hlls
func (h *Hll) MergeHll(k *Hll) {

	//hll[j] := MAX(hll1[j], hll2[j])

	for i, e := range k.M {
		if e > h.M[i] {
			h.M[i] = e
		}
	}
}
