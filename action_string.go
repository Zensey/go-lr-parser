// Code generated by "stringer -type=Action"; DO NOT EDIT.

package parser

import "strconv"

const _Action_name = "actShiftactReduceactAcceptactNoneactNop"

var _Action_index = [...]uint8{0, 8, 17, 26, 33, 39}

func (i Action) String() string {
	if i < 0 || i >= Action(len(_Action_index)-1) {
		return "Action(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Action_name[_Action_index[i]:_Action_index[i+1]]
}
