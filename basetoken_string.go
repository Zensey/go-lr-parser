// Code generated by "stringer -type=BaseToken"; DO NOT EDIT.

package parser

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TokUndefined-0]
	_ = x[Beg-1]
	_ = x[End-2]
	_ = x[None-3]
}

const _BaseToken_name = "TokUndefinedBegEndNone"

var _BaseToken_index = [...]uint8{0, 12, 15, 18, 22}

func (i BaseToken) String() string {
	if i < 0 || i >= BaseToken(len(_BaseToken_index)-1) {
		return "BaseToken(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BaseToken_name[_BaseToken_index[i]:_BaseToken_index[i+1]]
}