package xconfig_test

import (
	"Linda/baselibs/abstractions/xconfig"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testXEnvSuite struct {
	suite.Suite
}

type testConf struct {
	Port    int    `xenv:"testenv_port" xdefault:"1111"`
	Host    string `xenv:"testenv_host" xdefault:"localhost"`
	SubConf testSubConf
}

type testSubConf struct {
	Url string `xenv:"testsubenv_url" xdefault:"non"`
}

func (s *testXEnvSuite) TestLoadFromEnv() {
	s.T().Setenv("testenv_port", "1234")
	s.T().Setenv("testsubenv_url", "https://test.url")

	c := xconfig.NewFromOSEnv[testConf]()
	s.Equal(1234, c.Port)
	s.Equal("localhost", c.Host)
	s.Equal("https://test.url", c.SubConf.Url)
}

func TestXenvSuiteMain(t *testing.T) {
	suite.Run(t, new(testXEnvSuite))
}
