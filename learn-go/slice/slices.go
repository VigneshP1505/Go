package slice

import (
	"fmt"
	"slices"
)

func Slices() {
	var s = make([]string, 3)
	fmt.Println(s == nil)
	s = append(s, "s")
	s = append(s, "t")
	s[0] = "vignesh"
	s[1] = "dev"
	s[2] = "toronto"
	fmt.Println(s)
	var t = make([]string, len(s))
	copy(t, s)
	fmt.Println(t)
	fmt.Println(t == nil)
	fmt.Println(slices.Compare(t, s))

	var nums = make([]int64, 3)
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	var test = slices.Clone(nums)
	fmt.Println(test[0])
	test[1] = 40
	fmt.Println(nums[1])
	fmt.Println(slices.Compare(nums, test))
	fmt.Println(slices.Compare(test, nums))
	fmt.Println(slices.Index(nums, 10))
	fmt.Println(slices.Index(nums, 100))
}
