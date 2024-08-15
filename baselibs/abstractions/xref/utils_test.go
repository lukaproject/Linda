package xref_test

import (
	"Linda/baselibs/abstractions/xref"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testXrefUtilsSuite struct {
	suite.Suite
}

func (s *testXrefUtilsSuite) TestWalk() {
	type TestStruct2 struct {
		X string `xenv:"inner_x"`
	}
	type TestStruct struct {
		A   int `xenv:"linda_test_rp"`
		B   int
		C   string
		TS2 TestStruct2
	}

	testStruct := &TestStruct{}
	xref.WalkValues(testStruct, func(fieldName string, tags reflect.StructTag, t reflect.Type, v reflect.Value) {
		tagValue, ok := tags.Lookup("xenv")
		if ok {
			if fieldName == "A" {
				s.Equal("linda_test_rp", tagValue)
			}
			if fieldName == "X" {
				s.Equal("inner_x", tagValue)
			}
		}
		if v.CanSet() && v.Kind() != reflect.Struct {
			if v.Kind() == reflect.Int {
				v.SetInt(100)
			}
			if v.Kind() == reflect.String {
				v.SetString("test")
			}
		}
	})
	s.Equal(100, testStruct.A)
	s.Equal(100, testStruct.B)
	s.Equal("test", testStruct.C)
	s.Equal("test", testStruct.TS2.X)
}

func (s *testXrefUtilsSuite) TestKind() {
	p0 := []string{}
	var p1 []int = nil
	t0 := reflect.TypeOf(p0)
	t1 := reflect.TypeOf(p1)
	s.T().Log(t0.Kind())
	s.T().Log(t1.Kind())
}

func TestXrefUtils(t *testing.T) {
	suite.Run(t, new(testXrefUtilsSuite))
}
