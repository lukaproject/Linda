package gen_test

import (
	"Linda/baselibs/testcommon/gen"
	"strings"
	"testing"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type strGenTestSuite struct {
	suite.Suite
}

func (s *strGenTestSuite) TestStrGenAlpha() {
	charset := gen.CharsetLowerCase + gen.CharsetUpperCase
	str := xerr.Must(gen.StrGenerate(charset, 60, 100))
	for _, r := range str {
		s.True(strings.ContainsRune(charset, r))
	}
	s.True(60 <= len(str) && len(str) <= 100)
}

func (s *strGenTestSuite) TestStrGenFixedLen() {
	charset := gen.CharsetLowerCase + gen.CharsetUpperCase
	str := xerr.Must(gen.StrGenerate(charset, 75, 75))
	for _, r := range str {
		s.True(strings.ContainsRune(charset, r))
	}
	s.Len(str, 75)
}

func TestStrGenMain(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(strGenTestSuite))
}
