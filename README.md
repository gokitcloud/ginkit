[![Build Status](https://github.com/gokitcloud/ginkit/workflows/Run%20Tests/badge.svg?branch=master)](https://github.com/gokitcloud/ginkit/actions?query=branch%3Amaster)
[![codecov](https://codecov.io/gh/gokitcloud/ginkit/branch/master/graph/badge.svg)](https://codecov.io/gh/gokitcloud/ginkit)
[![Go Report Card](https://goreportcard.com/badge/github.com/gokitcloud/ginkit)](https://goreportcard.com/report/github.com/gokitcloud/ginkit)
[![GoDoc](https://pkg.go.dev/badge/github.com/gokitcloud/ginkit?status.svg)](https://pkg.go.dev/github.com/gokitcloud/ginkit?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/gokitcloud/ginkit/-/badge.svg)](https://sourcegraph.com/github.com/gokitcloud/ginkit?badge)
[![Open Source Helpers](https://www.codetriage.com/gokitcloud/ginkit/badges/users.svg)](https://www.codetriage.com/gokitcloud/ginkit)
[![Release](https://img.shields.io/github/release/gokitcloud/ginkit.svg?style=flat-square)](https://github.com/gokitcloud/ginkit/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gokitcloud/ginkit)](https://www.tickgit.com/browse?repo=github.com/gokitcloud/ginkit)

Ginkit is a web server written in [Go](https://go.dev/) extending [Gin](https://github.com/gokitcloud/ginkit/). It maintains a Gin-like API with completed boiler plate functionality to increase your productivity. If you love Gin, you will love Ginkit.

**The key features of Ginkit are:**

- TBD

## Getting started

### Prerequisites

- **[Go](https://go.dev/)**: any major [release](https://go.dev/doc/devel/release) since **1.18**.

### Getting Ginkit

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/gocloud/ginkit"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `gin` package:

```sh
$ go get -u github.com/gokitcloud/ginkit
```

### Running Gin

First you need to import Gin package for using Gin, one simplest example likes the follow `examples/example01/example.go`:

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gokitcloud/ginkit"
)

func main() {
	r := ginkit.Default()
	r.GET("/ping", gin.H{
		"message": "pong",
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

And use the Go command to run the demo:

```
# run example.go and visit 0.0.0.0:8080/ping on browser
$ go run example.go
```

### Learn more examples

#### Quick Start

Learn and practice more examples, please read the [Ginkit Quick Start](docs/doc.md) which includes API examples and builds tag.

#### Examples

A number of ready-to-run examples demonstrating various use cases of Ginkit in the [examples](examples) folder.

## Documentation

See [API documentation and descriptions](https://godoc.org/github.com/gokitcloud/ginkit) for package documentation.