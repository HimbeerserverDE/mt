// Code generated by "stringer -linecomment -type AttractionKind"; DO NOT EDIT.

package mt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoAttraction-0]
	_ = x[PointAttraction-1]
	_ = x[LineAttraction-2]
	_ = x[PlaneAttraction-3]
}

const _AttractionKind_name = "NonePointLinePlane"

var _AttractionKind_index = [...]uint8{0, 4, 9, 13, 18}

func (i AttractionKind) String() string {
	if i >= AttractionKind(len(_AttractionKind_index)-1) {
		return "AttractionKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AttractionKind_name[_AttractionKind_index[i]:_AttractionKind_index[i+1]]
}
