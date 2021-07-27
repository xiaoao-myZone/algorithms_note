package utils

// import "fmt"

func IntReverse(s []int, isCopy bool) []int {
	slice := make([]int, len(s))
	if isCopy {
		copy(slice, s)

	} else {
		slice = s
	}

	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[j], slice[i] = slice[i], slice[j]
	}
	if isCopy {
		return slice
	} else {
		return []int{}
	}
	// fmt.Println("yes")

}
