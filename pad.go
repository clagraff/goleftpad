/*
MIT License

Copyright (c) 2017 Curtis La Graff

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// pad provides basic functionallity for padding sides of a string.
package pad

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

// RightChar will pad the right of the string with the specified character to
// meet the specified length.
func RightChar(text string, length int, char byte) string {
	return Pad(text, RightDir, length, char)
}
