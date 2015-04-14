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
