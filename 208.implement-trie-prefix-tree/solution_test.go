package leetcode_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/zacharychang/leetcode/208.implement-trie-prefix-tree"
)

var _ = Describe("208.implement-trie-prefix-tree", func() {
	var (
		trie = Constructor()
	)

	Describe("Test Trie implements", func() {
		Context("Set namespace to test", func() {
			It("should be true", func() {
				trie.Insert("apple")
				Expect(trie.Search("apple")).To(Equal(true))
			})
			It("should be false", func() {
				Expect(trie.Search("app")).To(Equal(false))
			})
			It("should be true", func() {
				Expect(trie.StartsWith("app")).To(Equal(true))
			})
			It("should be true", func() {
				trie.Insert("app")
				Expect(trie.Search("app")).To(Equal(true))
			})
		})
	})

})
