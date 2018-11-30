package rrs

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
)

// Sample means an instance from the parameter space
// e.g. it include length and width in a 2-dimensional space
type Sample []struct {
	Parameter
	Val float64
}

func (p Sample) String() string {
	var buff bytes.Buffer
	for i := range p {
		str := fmt.Sprintf("{low=%2f, up=%2f, val=%2f} ",
			p[i].Low, p[i].Up, p[i].Val)
		buff.WriteString(str)
	}
	return buff.String()
}

// bound the earch space in the neighborhood of p
// Currently a simple method is used to construct, assume the
// parameter space D is defined by the upper and lower limits
// for its ith element, [li, ui], the neighborhood of x0 with
// a size of r · m(D) is the original parameter space scaled down by r
func (p Sample) bound(r float64) {
	for i := range p {
		low := p[i].Low
		up := p[i].Up
		val := p[i].Val

		if up-low < 0.001 {
			continue
		}
		percent := 1 - math.Pow(r, 1.0/float64(len(p)))
		p[i].Low = math.Min(low+low*percent*(1-(val-low)/(up-low)), val)
		p[i].Up = math.Max(up-up*percent*(1-(up-val)/(up-low)), val)
	}
}

// align move the bounded parameter space to a new place, and
// keep the paramter space size unchanged.
// @p will be will be located at the center
// NOTE: sample is a reference type
func (p Sample) align(params []Parameter) {
	for i := range p {
		low := p[i].Low
		up := p[i].Up
		step := (up - low) / 2
		p[i].Up = math.Min(p[i].Val+step, params[i].Up)
		p[i].Low = math.Max(p[i].Val-step, params[i].Low)
	}
}

// take creates a new sample according to the constraints
// from another sample
func (p Sample) take() Sample {
	p2 := make(Sample, len(p))
	for i := range p {
		p2[i].Low = p[i].Low
		p2[i].Up = p[i].Up
		p2[i].Name = p[i].Name
		p2[i].Val = rand.Float64()*(p2[i].Up-p2[i].Low) + p2[i].Low
	}
	return p2
}
