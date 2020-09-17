package functions

// RangeList ...
type RangeList struct {
	MinList int
	MaxList int
}

// MakeRange ...
func (r RangeList) MakeRange() []int {

	NewList := make([]int, r.MaxList-r.MinList+1)
	for Item := range NewList {
		NewList[Item] = r.MinList + Item
	}

	return NewList
}
