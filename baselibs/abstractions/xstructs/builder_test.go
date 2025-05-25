package xstructs_test

import (
	"Linda/baselibs/abstractions/xstructs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AClass struct {
	xstructs.DepStructHelper
	testIniter
}

type BClass struct {
	xstructs.DepStructHelper
	testIniter
}

type CClass struct {
	xstructs.DepStructHelper
	testIniter
}

type testIniter struct {
	inited bool
}

func (ti *testIniter) Init() error {
	ti.inited = true
	return nil
}

type testBuilderSuites struct {
	suite.Suite
}

func (s *testBuilderSuites) TestCyclingDependency() {
	builder := xstructs.NewDepStructsBuilder()
	ac := &AClass{
		DepStructHelper: xstructs.DepStructHelper{
			NameStr: "a",
		}}
	bc := &BClass{
		DepStructHelper: xstructs.DepStructHelper{
			NameStr: "b",
		}}
	cc := &CClass{
		DepStructHelper: xstructs.DepStructHelper{
			NameStr: "c",
		}}
	ac.AddDeps(bc)
	bc.AddDeps(cc)
	cc.AddDeps(ac)
	builder.AddDepStructs(ac, bc, cc)
	s.NotNil(builder.Build())
}

func (s *testBuilderSuites) TestNormalScenario() {
	builder := xstructs.NewDepStructsBuilder()
	ac := &AClass{
		DepStructHelper: xstructs.DepStructHelper{
			NameStr: "a",
		}}
	bc := &BClass{
		DepStructHelper: xstructs.DepStructHelper{
			NameStr: "b",
		}}
	cc := &CClass{
		DepStructHelper: xstructs.DepStructHelper{
			NameStr: "c",
		}}

	ac.AddDeps(bc)
	bc.AddDeps(cc)
	builder.AddDepStructs(ac, bc, cc)
	s.Nil(builder.Build())

	s.True(ac.inited)
	s.True(bc.inited)
	s.True(cc.inited)
}

func TestBuilder_Main(t *testing.T) {
	suite.Run(t, new(testBuilderSuites))
}
