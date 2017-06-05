package gopad

import "bytes"

type direction bool

// LeftDir is used to specify text will be padded on the left side.
const LeftDir direction = false

// RightDir is used to specify text will be padded on the right side.
const RightDir direction = true

// Pad is used to pad a given string using the specified character to satisfy the difference between the text length and the length specified.
func Pad(text string, dir direction, length int, char byte) string {
	startLength := len(text)
	diff := length - startLength

	// If no difference exists, return the text as it satisfies the length.
	if diff <= 0 {
		return text
	}

	// Writing to a buffer to improve performance.
	buff := bytes.NewBuffer(nil)
	if dir == RightDir {
		buff.WriteString(text)
	}

	for i := 0; i < diff; i++ {
		buff.WriteByte(char)
	}

	if dir == LeftDir {
		buff.WriteString(text)
	}

	return buff.String()
}

// Left will pad the left of the string with spaces to meet the specified length.
func Left(text string, length int) string {
	return Pad(text, LeftDir, length, ' ')
}

// LeftChar will pad the left of the string with the specified character to meet
// the specified length.
func LeftChar(text string, length int, char byte) string {
	return Pad(text, LeftDir, length, char)
}

// Right will pad the right of the string with spaces to meet the
// specified length.
func Right(text string, length int) string {
	return Pad(text, RightDir, length, ' ')
}

// RIghtChar will pad the right of the string with the specified character to
// meet the specified length.
func RightChar(text string, length int, char byte) string {
	return Pad(text, RightDir, length, char)
}
