package mt

type RangeF32 struct {
	min, max float32
	bias     float32
}

type RangeV3F32 struct {
	min, max [3]float32
	bias     float32
}
