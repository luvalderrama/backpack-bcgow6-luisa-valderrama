package ordenamiento

import(
	"sort"
)

func Order(slice []int) []int{
	sort.Ints(slice)
	return slice
}