package rrs

import "math"

// exploration phase, random sampling is used to identify
// a point in AD(r) for exploitation. The value of r should be first
// chosen. Based on this value and a predefined confidence probability p,
// the number of can be calculated as: n = ln(1−p)/ln(1-r)
type exploration struct {
	// p the confidence probability p should choose a value close
	// to 1, e.g. 0.99
	p float64
	// r decides the balance between exploration and exploitation
	r float64
	// n when r=0.1 and p=0.99, gs.n = 44
	n int
}

// exploitation As soon as exploration finds a promising point x0 whose
// function value is smaller than yr, we start a recursive random sampling
// procedure in the neighborhood N(x0) of x0.
type exploitation struct {
	q  float64 // confidence
	v  float64 // percentile
	c  float64 // ratio of shrink
	st float64 // min local ratio value of shrink
	l  int     // local number of samples in every interation
}

// GS settings about exploration
var GS exploration

// LS settings about exploitation
var LS exploitation

func init() {
	GS = exploration{
		p: 0.99,
		r: 0.1,
	}

	GS.n = int(math.Log(1-GS.p) / math.Log(1-GS.r))

	// TODO: how to set a correct value
	// Initialize exploitation parameters q, υ, c, st, l←ln(1−q)/ln(1−υ);
	LS = exploitation{
		q:  0.99,
		v:  0.3,
		c:  0.1,
		st: 0.000001,
	}
	LS.l = int(math.Log(1-LS.q) / math.Log(1-LS.v))
}
