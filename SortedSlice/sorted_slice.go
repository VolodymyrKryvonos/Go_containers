package SortedSlice

type SortSlice struct {
	slice []interface{}
	less  func(interface{}, interface{}) bool
}

func NewSortedSlice(less func(interface{}, interface{}) bool) *SortSlice {
	return &SortSlice{less: less}
}

func NewStringSlice() *SortSlice {
	return &SortSlice{less: func(a interface{}, b interface{}) bool {
		return a.(string) < b.(string)
	}}
}

func NewIntSlice() *SortSlice {
	return &SortSlice{less: func(a interface{}, b interface{}) bool {
		return a.(int) < b.(int)
	}}
}

func (slice *SortSlice) Clear() {
	slice.slice = nil
}
func (slice *SortSlice) Add(a interface{}) {
	if slice.slice == nil {
		slice.slice = []interface{}{a}
	} else if index := BinSearch(a, slice.slice, slice.less); len(slice.slice) == index {
		slice.slice = append(slice.slice, a)
	} else {
		newSlice := make([]interface{}, len(slice.slice)+1)
		at := copy(newSlice, slice.slice[:index])
		at += copy(newSlice[at:], []interface{}{a})
		copy(newSlice[at:], slice.slice[index:])
		slice.slice = newSlice
	}
}

func (slice *SortSlice) Remove(a interface{}) bool {
	index := BinSearch(a, slice.slice, slice.less)
	for ; index < len(slice.slice); index++ {
		if !slice.less(slice.slice[index], a) &&
			!slice.less(a, slice.slice[index]) {
			slice.slice = append(slice.slice[:index], slice.slice[index+1:]...)
			return true
		}
	}
	return false
}

func (slice *SortSlice) Index(a interface{}) int {
	index := BinSearch(a, slice.slice, slice.less)
	if index >= len(slice.slice) ||
		slice.less(slice.slice[index], a) ||
		slice.less(a, slice.slice[index]) {
		return -1
	}
	return index
}

func (slice SortSlice) At(index int) interface{} {
	if index < len(slice.slice) {
		return slice.slice[index]
	}
	panic("out of range")
}

func (slice SortSlice) Len() int {
	return len(slice.slice)
}

func BinSearch(a interface{}, slice []interface{}, less func(interface{}, interface{}) bool) int {
	left, right := 0, len(slice)
	for left < right {
		middle := int((left + right) / 2)
		if less(slice[middle], a) {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}
