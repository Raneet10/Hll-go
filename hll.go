package hll

import (
	"math"
)

const TWO_TO_32 = 1 << 32

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

//Operation to add an elemnt in the register set
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

// Count gives the cardinality of the register set
/*
	Z = Z + 1 / 2 ^ M[j]
	E = alpha(m) * m^2 * Z
*/

func (h *Hll) CountElements() uint {
	estimate := Estimate(h.M)

	if estimate < 2.5*float64(h.m) {
		if v := ZeroValueRegisters(h.M); v != 0 {
			return uint(float64(h.m) * math.Log(float64(h.m/v)))
		}
		return uint(estimate)

	} else if estimate < TWO_TO_32/30 {
		return uint(estimate)
	}

	return uint(-TWO_TO_32 * math.Log(1-estimate/TWO_TO_32))
}

//Merge operation for two Hlls
/*
	hll[j] := MAX(hll1[j], hll2[j])

*/
func (h *Hll) MergeHll(k *Hll) {

	for i, e := range k.M {
		if e > h.M[i] {
			h.M[i] = e
		}
	}
}
