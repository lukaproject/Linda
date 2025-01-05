package comm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SliceAssignment(t *testing.T) {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	b := a
	a = make([]int, 0)
	t.Logf("b=%v", b)
	t.Logf("a=%v", a)
	assert.Len(t, b, 10)
	assert.Len(t, a, 0)
}
