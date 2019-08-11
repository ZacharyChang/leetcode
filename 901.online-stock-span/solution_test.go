package leetcode_test

import (
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"

	. "github.com/zacharychang/leetcode/901.online-stock-span"
)

var _ = Describe("901.online-stock-span", func() {
	var (
		stock = Constructor()
	)

	Describe("Test stock span", func() {
		Context("Test Case 1", func() {
			It("should be 1", func() {
				Expect(stock.Next(100)).To(Equal(1))
			})
			It("should be 1", func() {
				Expect(stock.Next(80)).To(Equal(1))
			})
			It("should be 1", func() {
				Expect(stock.Next(60)).To(Equal(1))
			})
			It("should be 2", func() {
				Expect(stock.Next(70)).To(Equal(2))
			})
			It("should be 1", func() {
				Expect(stock.Next(60)).To(Equal(1))
			})
			It("should be 4", func() {
				Expect(stock.Next(75)).To(Equal(4))
			})
			It("should be 6", func() {
				Expect(stock.Next(85)).To(Equal(6))
			})
		})
	})
})
