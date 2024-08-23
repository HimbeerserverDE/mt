package mt

type TweenStyle uint8

const (
	Fwd TweenStyle = iota // Normal, linear interpolation.
	Rev
	Pulse
	Flicker
)

type TweenF32 struct {
	Style      TweenStyle
	Reps       uint16
	Beginning  float32
	Start, End float32
}

type TweenV2F32 struct {
	Style      TweenStyle
	Reps       uint16
	Beginning  float32
	Start, End [2]float32
}

type TweenV3F32 struct {
	Style      TweenStyle
	Reps       uint16
	Beginning  float32
	Start, End [3]float32
}

type TweenRangeF32 struct {
	Style      TweenStyle
	Reps       uint16
	Beginning  float32
	Start, End RangeF32
}

type TweenRangeV3F32 struct {
	Style      TweenStyle
	Reps       uint16
	Beginning  float32
	Start, End RangeV3F32
}
