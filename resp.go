package rrs

import "fmt"

// Response represents the optimized result
type Response struct {
	Sample
	Result []int
	// the extra values of rrs process
	Val interface{}
}

type responses []Response

func less(this []int, that []int) bool {
	for i := 0; i < len(this) && i < len(that); i++ {
		if this[i] < that[i] {
			return true
		} else if this[i] > that[i] {
			return false
		}
	}

	if len(this) < len(that) {
		return true
	} else if len(this) > len(that) {
		return false
	}
	return true
}

func (resp Response) String() string {
	s := resp.Sample.String()
	return fmt.Sprintf("%v %v", s, resp.Result)
}

// getBest return the the smallest result
func (resps responses) getBest() Response {
	var best Response
	for i := range resps {
		if i == 0 {
			best = resps[i]
			continue
		}
		if less(resps[i].Result, best.Result) {
			best = resps[i]
		}
	}
	return best
}
