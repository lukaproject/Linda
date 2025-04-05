package defaultor_test

import (
	"Linda/baselibs/abstractions/defaultor"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testX struct {
	P            int       `xdefault:"2"`
	G            string    `xdefault:"g"`
	StrSlice     []string  `xdefault:"test1,test2,test3"`
	IntSlice     []int     `xdefault:"1,2,400"`
	Float32Slice []float32 `xdefault:"1.1,2.2,400.4"`
	Float64Slice []float64 `xdefault:"1.1,2.2,400.4"`
}

type testV struct {
	A      int    `xdefault:"1"`
	B      string `xdefault:"0"`
	C      string `xdefault:"xxx"`
	XPtr   *testX
	X      testX
	DFloat float32 `xdefault:"3.35"`
	EBool  bool    `xdefault:"true"`
}

func TestDefaultValue(t *testing.T) {
	v := defaultor.New[testV]()
	assert.Equal(t, 1, v.A)
	assert.Equal(t, "0", v.B)
	assert.Equal(t, "xxx", v.C)
	assert.NotNil(t, v.XPtr)
	assert.Equal(t, 2, v.XPtr.P)
	assert.Equal(t, "g", v.XPtr.G)
	assert.Equal(t, 2, v.X.P)
	assert.Equal(t, "g", v.X.G)
	assert.Equal(t, 3, len(v.XPtr.StrSlice))
	assert.EqualValues(t, []string{"test1", "test2", "test3"}, v.XPtr.StrSlice)
	assert.Equal(t, 3, len(v.X.StrSlice))
	assert.EqualValues(t, []string{"test1", "test2", "test3"}, v.X.StrSlice)
	assert.EqualValues(t, []int{1, 2, 400}, v.X.IntSlice)
	eps := 1e-3
	assert.InEpsilonSlice(t, []float32{1.1, 2.2, 400.4}, v.X.Float32Slice, eps)
	assert.InEpsilonSlice(t, []float64{1.1, 2.2, 400.4}, v.X.Float64Slice, eps)
	assert.InEpsilon(t, 3.35, v.DFloat, eps)
	assert.True(t, v.EBool)
}
