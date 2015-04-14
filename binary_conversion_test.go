package binary_conversion

import (
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type BinaryConversionSuite struct{}

var _ = Suite(&BinaryConversionSuite{})

func (s *BinaryConversionSuite) TestUintToBoolArray(c *C) {
	var bits []bool
	bits = UintToBoolArray(214, 8)
	c.Assert(len(bits), Equals, 8)
	expectedBits := []bool{false, true, true, false, true, false, true, true}
	c.Assert(len(bits), Equals, len(expectedBits))
	for i, b := range bits {
		c.Assert(b, Equals, expectedBits[i])
	}
	c.Assert(BoolArrayToUint(bits), Equals, uint64(214))
}
