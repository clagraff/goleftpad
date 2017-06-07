# pad
Pad is a Golang package which provides functions to manage adding padding to 
the sides of strings. Like [Leftpad](https://www.npmjs.com/package/left-pad), 
but for Golang, and it supports right padding too!

[![CircleCI](https://circleci.com/gh/clagraff/pad/tree/master.svg?style=svg)](https://circleci.com/gh/clagraff/pad/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/clagraff/pad)](https://goreportcard.com/report/github.com/clagraff/pad)
[![Godoc](https://camo.githubusercontent.com/994c4dfb011a5952ca63171965089d673f8bc142/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f756e69706c616365732f636172626f6e3f7374617475732e737667)](https://godoc.org/github.com/clagraff/pad)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Installation
Just run in your terminal:

```
go get github.com/clagraff/pad
```

... and you are all set and ready to go!

## Examples
Here are a few examples of the package in action:

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

You can see more on the available functions, and how to use them, by reading the 
[![Godoc](https://camo.githubusercontent.com/994c4dfb011a5952ca63171965089d673f8bc142/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f756e69706c616365732f636172626f6e3f7374617475732e737667)](https://godoc.org/github.com/clagraff/pad) page for the package.


## Issues and Feedback
Feel free to submit issues, feature suggestions, or pull requests on Github
at: https://github.com/clagraff/pad

Updated tests, benchmarks, or improvements to the code are always welcome!

## Why?
Idk. I was bored and wanted something to do. So I made a left-pad implementation
in Golang.

Like it? Feel free to give it a star so I know it was appreciated. Thanks!
