package leetcode_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
	"sort"

	. "github.com/zacharychang/leetcode/384.shuffle-an-array"
)

var _ = Describe("Test shuffle an array", func() {
	var (
		nums  = []int{1, 2, 3}
		array = Constructor(nums)
	)

	Describe("Test shuffle an array", func() {
		Context("Array created", func() {
			It("should not be empty", func() {
				Expect(len(array.Original) > 0)
			})
			It("should be same length as params", func() {
				Expect(len(array.Original)).Should(Equal(len(nums)))
			})
			It("original and shuffle result should be same element", func() {
				array.Shuffle()
				Expect(equal(array.Original, array.ShuffleArray))
			})
			It("original and shuffle result should be same after reset", func() {
				array.Reset()
				Expect(array.Original, array.ShuffleArray)
			})
		})
	})
})

func equal(a []int, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)
	return reflect.DeepEqual(a, b)
}
