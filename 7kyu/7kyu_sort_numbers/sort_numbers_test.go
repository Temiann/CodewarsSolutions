package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample Test Cases", func() {
	It("should work for sample tests", func() {
		Expect(Expect(SortNumbers([]int{1, 2, 10, 50, 5})).To(Equal([]int{1, 2, 5, 10, 50})))
		Expect(Expect(SortNumbers([]int{})).To(Equal([]int{})))
	})
})
