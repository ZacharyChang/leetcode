package leetcode_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/zacharychang/leetcode/933.number-of-recent-calls"
)

var _ = Describe("933.number-of-recent-calls.number-of-recent-calls", func() {
	var (
		calls = Constructor()
	)

	Describe("Test Numbers of Recent Calls", func() {
		Context("Test Case 1", func() {
			It("should be 1", func() {
				Expect(calls.Ping(1)).To(Equal(1))
			})
			It("should be 2", func() {
				Expect(calls.Ping(100)).To(Equal(2))
			})
			It("should be 3", func() {
				Expect(calls.Ping(3001)).To(Equal(3))
			})
			It("should be 3", func() {
				Expect(calls.Ping(3002)).To(Equal(3))
			})
		})
	})
})
