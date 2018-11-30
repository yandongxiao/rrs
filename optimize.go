package rrs

import (
	"errors"
	"fmt"
	"math/rand"
)

// take @n samples according to the constraints from
// the @params paramters
func take(n int, params []Parameter) []Sample {
	ps := make([]Sample, n)
	for i, p := range ps {
		p = make(Sample, len(params))
		for i := range p {
			p[i].Low = params[i].Low
			p[i].Up = params[i].Up
			p[i].Name = params[i].Name
			p[i].Val = rand.Float64()*(p[i].Up-p[i].Low) + p[i].Low
		}
		ps[i] = p
	}
	return ps
}

// Optimize is the interface for optimizing parameters
func Optimize(req Request) (Response, error) {
	if req.Round < GS.n {
		msg := fmt.Sprintf("round value error: must >= %d", GS.n)
		return Response{}, errors.New(msg)
	}
	counter := 0
	exploreReponses := make(responses, 0, req.Round)

	// the algorithm uses the value of f(xn(1)) in the first n samples as the threshold value yr.
	// take n random samples xi, i = 1 . . . n from parameter space D
	// x0 ← arg min1≤i≤n(f(xi)), yr ← f(x0), add f(x0) to the threshold set F
	// NOTE: x0, yr, F can be calculated from exploreReponses
	for _, s := range take(GS.n, req.Params) {
		exploreReponses = append(exploreReponses, Response{
			Sample: s,
			result: req.Metric(s),
		})
		counter++
	}
	cand := exploreReponses.getBest()

	// the adjustment for the balance between exploration and exploitation.
	promisingResult := exploreReponses.getPromisingResult()
	exploitFlag := 1
	best := cand

	for counter < req.Round {
		if exploitFlag == 1 {
			// Exploit flag is set, start exploitation process
			// j←0, fc ←f(x0), xl ←x0, ρ←r;
			j := 0
			p := GS.r
			cand.bound(p)
			for p > LS.st && counter < req.Round {
				// take a random sample x′ from bounded parameter space
				s := cand.take()
				resp := Response{
					Sample: s,
					result: req.Metric(s),
				}
				counter++
				if resp.result < cand.result {
					// Find a better point, re-align the center
					// of sample space to the new point
					cand = resp
					cand.align(req.Params)
					j = 0
				} else {
					j++
				}

				if j == LS.l {
					// If random sampling fails to find a better point in l samples, that suggests
					// φN(x0)(f(x0)) is smaller than the expected level υ.
					// Fail to find a better point, then shrink the sample space
					// generate a new neighborhood N′(x0) whose size is c · m(N(x0))
					p = LS.c * p
					cand.bound(LS.c)
					j = 0
				}
			}
			exploitFlag = 0
			if cand.result < best.result {
				best = cand
			}
		}

		// do exploration again
		// any future sample with a smaller function value than yr
		// is considered to belong to AD(r)
		s := take(1, req.Params)[0]
		v := req.Metric(s)
		counter++
		exploreReponses = append(exploreReponses, Response{s, v})
		if v < promisingResult {
			// Find a promising point, set the flag to exploit
			exploitFlag = 1
			cand = Response{s, v}
		}

		// In later exploration, a new xn(1) is obtained every n samples.
		// and yr is updated with the average of these xn(1)
		if len(exploreReponses)%GS.n == 0 {
			promisingResult = exploreReponses.getPromisingResult()
		}
	}
	return best, nil
}
