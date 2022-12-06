package jgstr_test

import (
	"testing"

	jgstr "github.com/Jinglever/go-string"
	"github.com/bmizerany/assert"
)

func TestJsonEncode(t *testing.T) {
	assert.Equal(t, `{"a":1}`, jgstr.JsonEncode(map[string]int{"a": 1}))
}
