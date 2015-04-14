package binary_conversion

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
