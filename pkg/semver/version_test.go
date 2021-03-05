package semver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type version struct {
	major, minor, patch int
}

func (v version) Major() int {
	return v.major
}

func (v version) Minor() int {
	return v.minor
}

func (v version) Patch() int {
	return v.patch
}

type VersionTestSuite struct {
	suite.Suite
}

func TestVersionTestSuite(t *testing.T) {
	suite.Run(t, new(VersionTestSuite))
}

func (s *VersionTestSuite) TestCompare() {
	tests := []struct {
		v1, v2 Version
		result int
	}{
		{v1: version{major: 0}, v2: version{major: 0}, result: 0},
		{v1: version{major: 0}, v2: version{major: 1}, result: -1},
		{v1: version{major: 1}, v2: version{major: 0}, result: 1},
		{v1: version{major: 1}, v2: version{major: 1}, result: 0},

		{v1: version{major: 0, minor: 0}, v2: version{major: 0, minor: 1}, result: -1},
		{v1: version{major: 0, minor: 1}, v2: version{major: 0, minor: 0}, result: 1},
		{v1: version{major: 1, minor: 0}, v2: version{major: 1, minor: 1}, result: -1},
		{v1: version{major: 1, minor: 1}, v2: version{major: 1, minor: 0}, result: 1},

		{v1: version{major: 0, minor: 0, patch: 0}, v2: version{major: 0, minor: 0, patch: 1}, result: -1},
		{v1: version{major: 0, minor: 0, patch: 1}, v2: version{major: 0, minor: 0, patch: 0}, result: 1},
		{v1: version{major: 1, minor: 0, patch: 0}, v2: version{major: 1, minor: 0, patch: 1}, result: -1},
		{v1: version{major: 1, minor: 0, patch: 1}, v2: version{major: 1, minor: 0, patch: 0}, result: 1},
	}

	for _, test := range tests {
		result := Compare(test.v1, test.v2)
		assert.Equal(s.T(), test.result, result, "The result of compare version %v and %v should be %d", test.v1, test.v2, test.result)
	}
}
