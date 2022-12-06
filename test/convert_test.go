package jgstr_test

import (
	"testing"

	jgstr "github.com/Jinglever/go-string"
	"github.com/bmizerany/assert"
)

func TestCamelToSnake(t *testing.T) {
	assert.Equal(t, "jinglever_com", jgstr.CamelToSnake("JingleverCom"))
}

func TestFloatVal(t *testing.T) {
	assert.Equal(t, 1.0, jgstr.FloatVal("1"))
	assert.Equal(t, 0.0, jgstr.FloatVal("a"))
}
