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
	bits := UintToBoolArray(214, 8)
	c.Assert(len(bits), Equals, 8)
	expectedBits := []bool{false, true, true, false, true, false, true, true}
	c.Assert(len(bits), Equals, len(expectedBits))
	for i, b := range bits {
		c.Assert(b, Equals, expectedBits[i])
	}
	c.Assert(BoolArrayToUint(bits), Equals, uint64(214))
}

func (s *BinaryConversionSuite) TestUintOfLengthZero(c *C) {
	bits := UintToBoolArray(0, 0)
	c.Assert(len(bits), Equals, 0)
	c.Assert(BoolArrayToUint(bits), Equals, uint64(0))
}

func (s *BinaryConversionSuite) TestIntToBoolArray(c *C) {
	bits := IntToBoolArray(117, 8)
	c.Assert(len(bits), Equals, 8)
	expectedBits := []bool{true, false, true, false, true, true, true, false}
	c.Assert(len(bits), Equals, len(expectedBits))
	for i, b := range bits {
		c.Assert(b, Equals, expectedBits[i])
	}
	c.Assert(BoolArrayToInt(bits), Equals, int64(117))
}

func (s *BinaryConversionSuite) TestNegativeIntToBoolArray(c *C) {
	bits := IntToBoolArray(-113, 8)
	c.Assert(len(bits), Equals, 8)
	expectedBits := []bool{true, true, true, true, false, false, false, true}
	c.Assert(len(bits), Equals, len(expectedBits))
	for i, b := range bits {
		c.Assert(b, Equals, expectedBits[i])
	}
	c.Assert(BoolArrayToInt(bits), Equals, int64(-113))
}

func (s *BinaryConversionSuite) TestIntOfLengthZero(c *C) {
	bits := IntToBoolArray(0, 0)
	c.Assert(len(bits), Equals, 0)
	c.Assert(BoolArrayToInt(bits), Equals, int64(0))
}

func (s *BinaryConversionSuite) TestBitArrayToByteArray(c *C) {
	bits := UintToBoolArray(uint64(0x1234), 16)
	bytes := BoolArrayToByteArray(bits)
	c.Assert(len(bytes), Equals, 2)
	c.Assert(bytes[0], Equals, byte(0x34))
	c.Assert(bytes[1], Equals, byte(0x12))
}
