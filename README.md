# goleftpad
A library to add padding to the sides of strings. You can leftpad, or even 
rightpad if you're crazy!

[![CircleCI](https://circleci.com/gh/clagraff/pad/tree/master.svg?style=svg)](https://circleci.com/gh/clagraff/pad/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/clagraff/pad)](https://goreportcard.com/report/github.com/clagraff/pad)
[![Godoc](https://camo.githubusercontent.com/994c4dfb011a5952ca63171965089d673f8bc142/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f756e69706c616365732f636172626f6e3f7374617475732e737667)](https://godoc.org/github.com/clagraff/pad)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Installation

```
go get github.com/clagraff/pad
```

and BOOM, all set and ready to go!

## Examples
Here are a few small examples of the package in action:

```go
func main() {
    fmt.Println(pad.Left("foo", 6))             // "   foo"
    fmt.Println(pad.LeftChar("foo", 6, '~'))    // "~~~foo"

    fmt.Println(pad.Right("fizz", 4))           // "fizz "
    fmt.Println(pad.RightChar("fizz", 4, '+'))  // "fizz+"

    fmt.Println(
        pad.Pad("left", pad.LeftDir, 8, '_'),   // "____left"
        pad.Pad("right", pad.RightDir, 8, "_"), // "right___"
    )
}
```

Simple!

## Issues and Feedback
Feel free to submit issues, feature suggestions, or pull requests on Github
at: https://github.com/clagraff/pad

## Why?
Idk. I was bored and wanted something to do. So I made a left-pad implementation
in Golang.
