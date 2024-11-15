package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func dotest(s string, expected []string) {
	actual := wave(s)
	if len(expected) == 0 {
		Expect(actual).To(BeEmpty(), "with words = \"%s\"", s)
	} else {
		Expect(actual).To(Equal(expected), "with words = \"%s\"", s)
	}
}

var _ = Describe("Example test cases:", func() {
	It("should return the correct values", func() {
		dotest(" x yz", []string{" X yz", " x Yz", " x yZ"})
		dotest("abc", []string{"Abc", "aBc", "abC"})
		dotest("abc", []string{"Abc", "aBc", "abC"})
		dotest(" ab  c", []string{" Ab  c", " aB  c", " ab  C"})
		dotest("", []string{})
		dotest("z", []string{"Z"})
		dotest("a a a a a", []string{"A a a a a", "a A a a a", "a a A a a", "a a a A a", "a a a a A"})
		dotest("aaaaa", []string{"Aaaaa", "aAaaa", "aaAaa", "aaaAa", "aaaaA"})
		dotest("                                                           ", []string{})
		dotest(" a  b     c  defghijk l  m no pqrs tuvwx yz     ", []string{
			" A  b     c  defghijk l  m no pqrs tuvwx yz     ",
			" a  B     c  defghijk l  m no pqrs tuvwx yz     ",
			" a  b     C  defghijk l  m no pqrs tuvwx yz     ",
			" a  b     c  Defghijk l  m no pqrs tuvwx yz     ",
			" a  b     c  dEfghijk l  m no pqrs tuvwx yz     ",
			" a  b     c  deFghijk l  m no pqrs tuvwx yz     ",
			" a  b     c  defGhijk l  m no pqrs tuvwx yz     ",
			" a  b     c  defgHijk l  m no pqrs tuvwx yz     ",
			" a  b     c  defghIjk l  m no pqrs tuvwx yz     ",
			" a  b     c  defghiJk l  m no pqrs tuvwx yz     ",
			" a  b     c  defghijK l  m no pqrs tuvwx yz     ",
			" a  b     c  defghijk L  m no pqrs tuvwx yz     ",
			" a  b     c  defghijk l  M no pqrs tuvwx yz     ",
			" a  b     c  defghijk l  m No pqrs tuvwx yz     ",
			" a  b     c  defghijk l  m nO pqrs tuvwx yz     ",
			" a  b     c  defghijk l  m no Pqrs tuvwx yz     ",
			" a  b     c  defghijk l  m no pQrs tuvwx yz     ",
			" a  b     c  defghijk l  m no pqRs tuvwx yz     ",
			" a  b     c  defghijk l  m no pqrS tuvwx yz     ",
			" a  b     c  defghijk l  m no pqrs Tuvwx yz     ",
			" a  b     c  defghijk l  m no pqrs tUvwx yz     ",
			" a  b     c  defghijk l  m no pqrs tuVwx yz     ",
			" a  b     c  defghijk l  m no pqrs tuvWx yz     ",
			" a  b     c  defghijk l  m no pqrs tuvwX yz     ",
			" a  b     c  defghijk l  m no pqrs tuvwx Yz     ",
			" a  b     c  defghijk l  m no pqrs tuvwx yZ     "})
	})
})
