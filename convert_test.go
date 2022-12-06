package jgstr

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCamelToSnake(t *testing.T) {
	assert.Equal(t, "jinglever_com", CamelToSnake("JingleverCom"))
}

func TestFloatVal(t *testing.T) {
	assert.Equal(t, 1.0, FloatVal("1"))
	assert.Equal(t, 0.0, FloatVal("a"))
}
