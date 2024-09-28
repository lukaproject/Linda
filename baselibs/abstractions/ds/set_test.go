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
	for i := 0; i < 100; i++ {
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

func Test_testSetSuiteMain(t *testing.T) {
	s := new(testSetSuite)
	suite.Run(t, s)
}
