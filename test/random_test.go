package jgstr_test

import (
	"fmt"
	"testing"

	jgstr "github.com/Jinglever/go-string"
)

func TestApproach8(t *testing.T) {
	fmt.Println(jgstr.RandStr(10))
}

func BenchmarkApproach8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = jgstr.RandStr(10)
	}
}
