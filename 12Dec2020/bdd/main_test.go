package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sum", func() {
	var (
		p, q, m, n, sum1, sum2 int
	)

	By(testEnvironment)
	BeforeEach(func() {
		p, q, sum1 = 5, 6, 11
		// Putting wrong value of sum2 intentionally
		m, n, sum2 = 8, 7, 16
	})
	Context("By providing two digits as input", func() {
		It("should return sum of the two digits", func() {
			addition_of_p_and_q := Sum(p, q)
			Expect(addition_of_p_and_q).Should(Equal(sum1))
		})
		It("should not return the sum provided", func() {
			addition_of_m_and_n := Sum(m, n)
			Expect(addition_of_m_and_n).ShouldNot(Equal(sum2))
		})
	})
})
