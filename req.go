package rrs

// Request describes the requirements about optimization
type Request struct {
	Round  int
	Params []Parameter
	Metric MetricFunc
}

// Parameter specify the range of paramters [low, up)
type Parameter struct {
	Name string
	Low  float64
	Up   float64
}

// MetricFunc the value of ps, the smaller the better
type MetricFunc func(p Sample) int
