package jgstr_test

import (
	"testing"

	jgstr "github.com/Jinglever/go-string"
	"github.com/bmizerany/assert"
)

func TestFormatGo(t *testing.T) {
	src := []byte(`package jgstr
	import "fmt"
	func demo() {
	fmt.Println("hello world")
	}`)
	dst, err := jgstr.FormatGo(src)
	if err != nil {
		t.Fatal(err)
	}
	eq := `package jgstr

import "fmt"

func demo() {
	fmt.Println("hello world")
}
`
	assert.Equal(t, eq, string(dst))
}
