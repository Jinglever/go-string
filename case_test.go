package jgstr

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestUcfirst(t *testing.T) {
	assert.Equal(t, "Jinglever", Ucfirst("jinglever"))
}
