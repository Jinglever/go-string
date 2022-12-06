package jgstr_test

import (
	"testing"

	jgstr "github.com/Jinglever/go-string"
	"github.com/bmizerany/assert"
)

func TestMD5(t *testing.T) {
	assert.Equal(t, "5eb63bbbe01eeed093cb22bb8f5acdc3", jgstr.MD5("hello world"))
}
