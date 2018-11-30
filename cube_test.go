package rrs_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/yandongxiao/rrs"
)

func TestCube(t *testing.T) {
	fmt.Println(rrs.GS, rrs.LS)
	rand.Seed(int64(time.Now().Nanosecond()))
	params := []rrs.Parameter{
		{
			Name: "length",
			Low:  1,
			Up:   2000,
		},

		{
			Name: "width",
			Low:  1,
			Up:   2000,
		},

		{
			Name: "height",
			Low:  1,
			Up:   2000,
		},
	}
	fmt.Printf("value space=[%d, %d]\n", 1, 2000*2000*2000)

	round := 50
	for i := 1; i < 12; i++ {
		fmt.Println("round=", round)
		fmt.Println(rrs.Optimize(rrs.Request{
			Round:  round,
			Params: params,
			Metric: func(p rrs.Sample) int {
				getVolume := func(p rrs.Sample) float64 {
					res := 1.0
					for i := range p {
						res *= p[i].Val
					}
					return res
				}
				return int(math.Abs(50000 - getVolume(p)))
			},
		}))
		round *= 2
	}
}