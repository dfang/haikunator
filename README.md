# Haikunator

[![Build Status](https://img.shields.io/travis/dfang/haikunator.svg?style=flat-square)](https://travis-ci.org/dfang/haikunator?branch=master)
[![Coverage Status](https://img.shields.io/coveralls/dfang/haikunator.svg?style=flat-square)](https://coveralls.io/r/dfang/haikunator)
[![Go Report Card](https://goreportcard.com/badge/github.com/dfang/haikunator?style=flat-square)](https://goreportcard.com/report/github.com/dfang/haikunator)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/github.com/dfang/haikunator)

Generate Heroku-like random names to use in your go applications.

this is FORKED from [atrox](https://github.com/atrox/haikunatorgo) and add words from this [fork](https://github.com/TykTechnologies/go-haikunator)

## Install haikunator Command

```
go install github.com/dfang/haikunator/cmd/haikunator
```

## Usage in go code

Haikunator is pretty simple.

```go
package main

import (
  "github.com/dfang/haikunator"
)

func main() {
  haikunator := haikunator.New()

  // default usage
  haikunator.Haikunate() // => "wispy-dust-1337"

  // custom length (default=4)
  haikunator.TokenLength = 9
  haikunator.Haikunate() // => "patient-king-887265"

  // use hex instead of numbers
  haikunator.TokenHex = true
  haikunator.Haikunate() // => "purple-breeze-98e1"

  // use custom chars instead of numbers/hex
  haikunator.TokenChars = "HAIKUNATE"
  haikunator.Haikunate() // => "summer-atom-IHEA"

  // don't include a token
  haikunator.TokenLength = 0
  haikunator.Haikunate() // => "cold-wildflower"

  // use a different delimiter
  haikunator.Delimiter = "."
  haikunator.Haikunate() // => "restless.sea.7976"

  // no token, space delimiter
  haikunator.TokenLength = 0
  haikunator.Delimiter = " "
  haikunator.Haikunate() // => "delicate haze"

  // no token, empty delimiter
  haikunator.TokenLength = 0
  haikunator.Delimiter = ""
  haikunator.Haikunate() // => "billowingleaf"

  // custom nouns and/or adjectives
  haikunator.Adjectives = []string{"red", "green", "blue"}
  haikunator.Nouns = []string{"reindeer"}
  haikunator.Haikunate() // => "blue-reindeer-4252"
}
```

## Options

The following options are available:

```go
Haikunator{
  Adjectives:  []string{"custom", "adjectives"},
  Nouns:       []string{"custom", "nouns"},
  Delimiter:   "-",
  TokenLength: 4,
  TokenHex:    false,
  TokenChars:  "0123456789",
  Random:      rand.New(rand.NewSource(time.Now().UnixNano())),
}
```
*If ```TokenHex``` is true, it overrides any tokens specified in ```TokenChars```*

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs](https://github.com/atrox/haikunatorgo/issues)
- Fix bugs and [submit pull requests](https://github.com/atrox/haikunatorgo/pulls)
- Write, clarify, or fix documentation
- Suggest or add new features

## Other Languages

Haikunator is also available in other languages. Check them out:

- Node: https://github.com/Atrox/haikunatorjs
- .NET: https://github.com/Atrox/haikunator.net
- Python: https://github.com/Atrox/haikunatorpy
- PHP: https://github.com/Atrox/haikunatorphp
- Java: https://github.com/Atrox/haikunatorjava
- Dart: https://github.com/Atrox/haikunatordart
- Ruby: https://github.com/usmanbashir/haikunator
- Rust: https://github.com/nishanths/rust-haikunator
