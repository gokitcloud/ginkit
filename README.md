# Ginkit Golang Gin Web Framework

<img align="right" width="159px" src="https://raw.githubusercontent.com/gokitcloud/logo/main/web.png">

[![Build Status](https://github.com/gokitcloud/ginkit/workflows/Run%20Tests/badge.svg?branch=default)](https://github.com/gokitcloud/ginkit/actions?query=branch%3Adefault)
[![Go Report Card](https://goreportcard.com/badge/github.com/gokitcloud/ginkit)](https://goreportcard.com/report/github.com/gokitcloud/ginkit)
[![GoDoc](https://pkg.go.dev/badge/github.com/gokitcloud/ginkit?status.svg)](https://pkg.go.dev/github.com/gokitcloud/ginkit?tab=doc)
[![Release](https://img.shields.io/github/release/gokitcloud/ginkit.svg?style=flat-square)](https://github.com/gokitcloud/ginkit/releases)

We love [Gin](https://github.com/gokitcloud/ginkit/)! Ginkit is written in [Go](https://go.dev/) and extends the Gin Web Framework with essential features and enhancements to make your life better! If you love Gin, you will love Ginkit.

**The key features of Ginkit are:**

- easy
- saml auth
- token auth
- template and layout built in
- reverse proxy built in
- simple data wrappers
- auto handling of response data
- oauto support built in

## Getting started

### Prerequisites

- **[Go](https://go.dev/)**: any major [release](https://go.dev/doc/devel/release) since **1.18**.

### Getting Ginkit

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/gokitcloud/ginkit"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `ginkit` package:

```sh
$ go get -u github.com/gokitcloud/ginkit
```

### Running Ginkit

First you need to import Ginkit package for using Gin, and set up the simple example below:  
_from [example01](examples/example01/example.go)_

```go
package main

import (
	"github.com/gokitcloud/ginkit"
)

func main() {
	r := ginkit.Default()
	r.GET("/ping", ginkit.H{
		"message": "pong",
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

And use the Go command to run the demo:

```
# run example.go from the examples/example01 directory
# and visit 0.0.0.0:8080/ping on browser
$ go run example.go
```

### Learn more examples

#### Quick Start

Learn and practice more examples, please read the [Ginkit Quick Start](docs/doc.md) which includes API examples and builds tag.

#### Examples

A number of ready-to-run examples demonstrating various use cases of Ginkit in the [examples](examples) folder.

## Documentation

See [API documentation and descriptions](https://godoc.org/github.com/gokitcloud/ginkit) for package documentation.
