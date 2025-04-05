package ds_test

import (
	"Linda/baselibs/abstractions/ds"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testSetSuite struct {
	suite.Suite
}

func (s *testSetSuite) TestListByChan() {
	testSet := make(ds.Set[int])
	expectedList := make([]int, 0)
	for i := range 100 {
		testSet.Insert(i)
		expectedList = append(expectedList, i)
	}
	ch := make(chan int, 10)

	go testSet.ListByChan(ch)

	actualList := make([]int, 0)
	for {
		v, ok := <-ch
		if !ok {
			break
		} else {
			actualList = append(actualList, v)
		}
	}
	s.ElementsMatch(expectedList, actualList)
}

func (s *testSetSuite) TestGeneralScenario() {
	testSet := make(ds.Set[int])
	for i := range 100 {
		testSet.Insert(i)
	}
	s.Equal(100, testSet.Len())
	for i := 0; i < 100; i += 2 {
		testSet.Remove(i)
	}
	s.Equal(50, testSet.Len())
	s.True(testSet.Exist(77))
}

func Test_testSetSuiteMain(t *testing.T) {
	s := new(testSetSuite)
	suite.Run(t, s)
}
