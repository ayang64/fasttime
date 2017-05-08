package main

import (
	"testing"
	"time"
)

func BenchmarkSysTime(t *testing.B) {
	for n := 0; n < t.N; n++ {
		time.Now()
	}
}

func BenchmarkFastTime(t *testing.B) {
	for n := 0; n < t.N; n++ {
		Now()
	}
}
