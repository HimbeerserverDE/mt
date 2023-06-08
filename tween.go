package mt

type TweenF32 struct {
	Start, End float32
}

type TweenV2F32 struct {
	Start, End [2]float32
}

type TweenV3F32 struct {
	Start, End [3]float32
}

type TweenRangeF32 struct {
	Start, End RangeF32
}

type TweenRangeV3F32 struct {
	Start, End RangeV3F32
}
