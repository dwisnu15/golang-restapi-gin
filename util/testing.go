package util

//yet to be used

import (
	"testing"
	"unicode/utf8"
)
//this is a benchmarking function
func RuneCount(s string) int {
	return len([]rune(s))
}

//this will be compared to RuneCount
func RuneCount2(s string) int {
	return utf8.RuneCountInString(s)
}

//run the benchmarking for RuneCount
func BenchmarkRuneCount(b *testing.B) {
	s := "Gophers are amazing 😁"
	for i := 0; i < b.N; i++ {
		RuneCount(s)
	}
}
//same but for RuneCount2
func BenchmarkRuneCount2(b *testing.B) {
	s := "Gophers are amazing 😁"
	for i := 0; i < b.N; i++ {
		RuneCount2(s)
	}
}
