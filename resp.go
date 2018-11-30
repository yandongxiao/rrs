package rrs

import (
	"strconv"
)

// Response represents the optimized result
type Response struct {
	Sample
	result int
}

type responses []Response

func (resp Response) String() string {
	s := resp.Sample.String()
	return s + strconv.Itoa(resp.result)
}

// getBest return the the smallest result
func (resps responses) getBest() Response {
	var best Response
	for i := range resps {
		if i == 0 {
			best = resps[i]
			continue
		}
		if resps[i].result < best.result {
			best = resps[i]
		}
	}
	return best
}
