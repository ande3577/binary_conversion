package binary_conversion

import "math"

func UintToBoolArray(value uint64, bitCount int) []bool {
	bits := make([]bool, bitCount)
	for i := 0; i < bitCount; i++ {
		if (value & (1 << uint(i))) != 0 {
			bits[i] = true
		}
	}
	return bits
}

func BoolArrayToUint(bits []bool) uint64 {
	var value uint64 = 0
	for i, b := range bits {
		if b {
			value |= (1 << uint(i))
		}
	}
	return value
}

func MaxUint(bitCount uint) uint64 {
	return (uint64(1) << bitCount) - 1
}

func IntToBoolArray(value int64, bitCount int) []bool {
	if value < 0 {
		bits := make([]bool, bitCount)
		unsignedValue := uint64(-value)
		unsignedValue++

		for i := 0; i < bitCount; i++ {
			if (value & (1 << uint(i))) != 0 {
				bits[i] = true
			}
		}

		return bits
	} else {
		return UintToBoolArray(uint64(value), bitCount)
	}
}

func BoolArrayToInt(bits []bool) int64 {
	if len(bits) > 0 && bits[len(bits)-1] {
		var unsignedValue uint64 = 0
		for i, b := range bits {
			if !b {
				unsignedValue |= (1 << uint(i))
			}
		}
		unsignedValue += 1
		return -int64(unsignedValue)
	} else {
		return int64(BoolArrayToUint(bits))
	}
}

func MaxInt(bitCount uint) int64 {
	return (int64(1) << (bitCount - 1)) - 1
}

func MinInt(bitCount uint) int64 {
	return -(int64(1) << (bitCount - 1))
}

func BoolArrayToByteArray(bits []bool) []byte {
	j := uint(0)
	by := byte(0)
	bytes := make([]byte, (len(bits)+7)/8)
	for i, b := range bits {
		if b {
			by |= (1 << j)
		}
		j++
		if j == 8 {
			j = 0
			bytes[i/8] = by
			by = byte(0)
		}
	}
	return bytes
}

func StringToBoolArray(str string, bitsLength int) []bool {
	bits := make([]bool, bitsLength)

	j := uint(0)
	for i := 0; i < bitsLength; i++ {
		if i/8 < len(str) {
			ch := str[i/8]
			if ch&(1<<j) != 0 {
				bits[i] = true
			}
			j++
			if j == 8 {
				j = 0
			}
		}
	}

	return bits
}

func BoolArrayToString(bits []bool) string {
	str := ""
	var ch rune
	j := uint(0)
	for _, b := range bits {
		if b {
			ch |= (1 << j)
		}
		j++
		if j == 8 {
			j = 0
			if ch != 0 {
				str += string(ch)
			}
			ch = 0
		}
	}
	return str[:len(str)]
}

func Float32ToBoolArray(f float32) []bool {
	u32 := math.Float32bits(f)
	return UintToBoolArray(uint64(u32), 32)
}

func BoolArrayToFloat32(bits []bool) float32 {
	u32 := uint32(BoolArrayToUint(bits))
	return math.Float32frombits(u32)
}

func Float64ToBoolArray(f float64) []bool {
	u64 := math.Float64bits(f)
	return UintToBoolArray(u64, 64)
}

func BoolArrayToFloat64(bits []bool) float64 {
	u64 := BoolArrayToUint(bits)
	return math.Float64frombits(u64)
}
