# go-highlight [![GoDoc](https://godoc.org/github.com/d4l3k/go-highlight?status.svg)](https://godoc.org/github.com/d4l3k/go-highlight) [![Build Status](https://travis-ci.org/d4l3k/go-highlight.svg?branch=master)](https://travis-ci.org/d4l3k/go-highlight)

A Go (Golang) code syntax highlighting library. It uses automatically converted
[highlight.js](https://github.com/isagalaev/highlight.js) language definitions.

## Usage

```go
package main

import "github.com/d4l3k/go-highlight"

func main() {
  highlight.Highlight("go", `
    package main

    import "fmt"

    func main() {
      fmt.Println("Duck!")
    }
  `)
  /*
    <keyword>package</keyword> main

    <keyword>import</keyword> <string>"fmt"</string>

    <keyword>func</keyword> main() {
      fmt.Println(<string>"Duck!"</string>)
    }
  */
}
```

## Copyright

The code written by Tristan Rice is licensed under the MIT license.

The language definitions are ported from
[highlight.js](https://github.com/isagalaev/highlight.js) which is licensed
under the BSD licence.
