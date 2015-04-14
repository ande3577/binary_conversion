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
