// Code generated by "stringer -type Endianness"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LittleEndian-0]
	_ = x[BigEndian-1]
	_ = x[NonEndian-2]
}

const _Endianness_name = "LittleEndianBigEndianNonEndian"

var _Endianness_index = [...]uint8{0, 12, 21, 30}

func (i Endianness) String() string {
	if i < 0 || i >= Endianness(len(_Endianness_index)-1) {
		return "Endianness(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Endianness_name[_Endianness_index[i]:_Endianness_index[i+1]]
}
