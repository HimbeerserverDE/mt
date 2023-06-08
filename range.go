package mt

type RangeF32 struct {
	Min, Max float32
	Bias     float32
}

type RangeV3F32 struct {
	Min, Max [3]float32
	Bias     float32
}
