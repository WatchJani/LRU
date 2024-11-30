package main

import "testing"

func BenchmarkInset(b *testing.B) {
	b.StopTimer()

	f := new(DLL)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		f.Inset(i)
	}
}

func BenchmarkRead(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
