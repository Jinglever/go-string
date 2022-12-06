package jgstr

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMD5(t *testing.T) {
	assert.Equal(t, "5eb63bbbe01eeed093cb22bb8f5acdc3", MD5("hello world"))
}
