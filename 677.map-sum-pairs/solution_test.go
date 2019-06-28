package leetcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map Sum Pairs Test - Prefix Hashtable solution", func() {

	Describe("Test Start", func() {
		Context("Test Case 1", func() {
			mapSum := Constructor()
			It("Should return 3", func() {
				mapSum.Insert("apple", 3)
				Expect(mapSum.Sum("ap")).To(Equal(3))
			})
			It("Should return 5", func() {
				mapSum.Insert("app", 2)
				Expect(mapSum.Sum("ap")).To(Equal(5))
			})
		})
		Context("Test Case 2", func() {
			mapSum := Constructor()
			It("Should return 3", func() {
				mapSum.Insert("aa", 3)
				Expect(mapSum.Sum("a")).To(Equal(3))
			})
			It("Should return 5", func() {
				mapSum.Insert("aa", 2)
				Expect(mapSum.Sum("a")).To(Equal(2))
			})
		})
		Context("Test Case 3", func() {
			mapSum := Constructor()
			It("Should return 3", func() {
				mapSum.Insert("apple", 3)
				Expect(mapSum.Sum("apple")).To(Equal(3))
			})
			It("Should return 2", func() {
				mapSum.Insert("app", 2)
				Expect(mapSum.Sum("ap")).To(Equal(5))
			})
		})
	})
})
