package agents

import (
	"Linda/baselibs/codes/errno"
	"Linda/baselibs/testcommon/testenv"
	"testing"

	"github.com/stretchr/testify/suite"
)

// This is an internal struct test, not be exposed
// in public.

type nodeStatesTestSuite struct {
	testenv.TestBase
}

func (s *nodeStatesTestSuite) TestJoinSuccess() {
	testBagName := "test-bag"
	state := newNodeStates()
	err := state.Join(testBagName)
	s.Nil(err)
	s.Equal(emptyBagName, state.GetBagName())
	s.True(state.IsOnGoingStates())
	s.Equal(testBagName, state.BagName)
	state.JoinFinished(testBagName)
	s.Equal(testBagName, state.GetBagName())
}

func (s *nodeStatesTestSuite) TestMultiTimeJoinSuceess() {
	testBagName := "test-bag"
	state := newNodeStates()
	err := state.Join(testBagName)
	s.Nil(err)
	s.Equal(emptyBagName, state.GetBagName())
	s.True(state.IsOnGoingStates())
	s.Equal(testBagName, state.BagName)
	s.Equal(emptyBagName, state.GetBagName())
	s.True(state.IsOnGoingStates())
	s.Equal(emptyBagName, state.GetBagName())
	s.True(state.IsOnGoingStates())
	state.JoinFinished(testBagName)
	s.Equal(testBagName, state.GetBagName())
}

func (s *nodeStatesTestSuite) TestJoinWithDifferentBagNameFailed() {
	testBagName := "test-bag"
	testAnotherBagName := "test-another-bag"
	state := newNodeStates()
	s.Nil(state.Join(testBagName))
	s.Equal(
		errno.ErrNodeBelongsToAnotherBag,
		state.Join(testAnotherBagName))
}

func (s *nodeStatesTestSuite) TestFreeNodeSuccess() {
	testBagName := "test-bag"
	state := newNodeStates()
	s.Nil(state.Join(testBagName))
	state.JoinFinished(testBagName)
	s.True(state.IsSteadyStates())
	state.Free()
	s.True(state.IsOnGoingStates())
	state.FreeFinished()
	s.True(state.IsSteadyStates())
	s.Equal(emptyBagName, state.GetBagName())
}

func TestNodeStatesTestSuiteMain(t *testing.T) {
	suite.Run(t, new(nodeStatesTestSuite))
}
