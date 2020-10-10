package simple

func Slice() []int {
	slice1 := make([]int, 3)
	slice2 := []int{1, 2, 3}
	return append(slice1, slice2...)
}
