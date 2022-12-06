package jgstr_test

import (
	"testing"

	jgstr "github.com/Jinglever/go-string"
	"github.com/bmizerany/assert"
)

func TestUcfirst(t *testing.T) {
	assert.Equal(t, "Jinglever", jgstr.Ucfirst("jinglever"))
}
