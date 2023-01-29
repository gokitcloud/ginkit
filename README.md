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

First you need to import Gin package for using Gin, one simplest example likes the follow `examples/example1/example.go`:

```go
package main

import (
	"log"

	"github.com/gokitcloud/ginkit"
)

func main() {
	e := ginkit.NewDefault()

	e.Router().GET("/test", ginkit.WrapDataFunc(test))

	err := e.Run("8080")
	if err != nil {
		log.Println(err)
	}
}

func test() (any, error) {
	return map[string]any{"foo": "bar"}, nil
}


package main

import (
  "net/http"

  "github.com/gokitcloud/ginkit"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
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

Learn and practice more examples, please read the [Gin Quick Start](docs/doc.md) which includes API examples and builds tag.

#### Examples

A number of ready-to-run examples demonstrating various use cases of Gin on the [Gin examples](https://github.com/gin-gonic/examples) repository.

## Documentation

See [API documentation and descriptions](https://godoc.org/github.com/gokitcloud/ginkit) for package.

All documentation is available on the Gin website.

- [English](https://gin-gonic.com/docs/)
- [简体中文](https://gin-gonic.com/zh-cn/docs/)
- [繁體中文](https://gin-gonic.com/zh-tw/docs/)
- [日本語](https://gin-gonic.com/ja/docs/)
- [Español](https://gin-gonic.com/es/docs/)
- [한국어](https://gin-gonic.com/ko-kr/docs/)
- [Turkish](https://gin-gonic.com/tr/docs/)
- [Persian](https://gin-gonic.com/fa/docs/)

### Articles about Gin

A curated list of awesome Gin framework.

- [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)

## Benchmarks

Gin uses a custom version of [HttpRouter](https://github.com/julienschmidt/httprouter), [see all benchmarks details](/BENCHMARKS.md).

| Benchmark name                 |       (1) |             (2) |          (3) |             (4) |
| ------------------------------ | --------: | --------------: | -----------: | --------------: |
| BenchmarkGin_GithubAll         | **43550** | **27364 ns/op** |   **0 B/op** | **0 allocs/op** |
| BenchmarkAce_GithubAll         |     40543 |     29670 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkAero_GithubAll        |     57632 |     20648 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkBear_GithubAll        |      9234 |    216179 ns/op |   86448 B/op |   943 allocs/op |
| BenchmarkBeego_GithubAll       |      7407 |    243496 ns/op |   71456 B/op |   609 allocs/op |
| BenchmarkBone_GithubAll        |       420 |   2922835 ns/op |  720160 B/op |  8620 allocs/op |
| BenchmarkChi_GithubAll         |      7620 |    238331 ns/op |   87696 B/op |   609 allocs/op |
| BenchmarkDenco_GithubAll       |     18355 |     64494 ns/op |   20224 B/op |   167 allocs/op |
| BenchmarkEcho_GithubAll        |     31251 |     38479 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkGocraftWeb_GithubAll  |      4117 |    300062 ns/op |  131656 B/op |  1686 allocs/op |
| BenchmarkGoji_GithubAll        |      3274 |    416158 ns/op |   56112 B/op |   334 allocs/op |
| BenchmarkGojiv2_GithubAll      |      1402 |    870518 ns/op |  352720 B/op |  4321 allocs/op |
| BenchmarkGoJsonRest_GithubAll  |      2976 |    401507 ns/op |  134371 B/op |  2737 allocs/op |
| BenchmarkGoRestful_GithubAll   |       410 |   2913158 ns/op |  910144 B/op |  2938 allocs/op |
| BenchmarkGorillaMux_GithubAll  |       346 |   3384987 ns/op |  251650 B/op |  1994 allocs/op |
| BenchmarkGowwwRouter_GithubAll |     10000 |    143025 ns/op |   72144 B/op |   501 allocs/op |
| BenchmarkHttpRouter_GithubAll  |     55938 |     21360 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkHttpTreeMux_GithubAll |     10000 |    153944 ns/op |   65856 B/op |   671 allocs/op |
| BenchmarkKocha_GithubAll       |     10000 |    106315 ns/op |   23304 B/op |   843 allocs/op |
| BenchmarkLARS_GithubAll        |     47779 |     25084 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkMacaron_GithubAll     |      3266 |    371907 ns/op |  149409 B/op |  1624 allocs/op |
| BenchmarkMartini_GithubAll     |       331 |   3444706 ns/op |  226551 B/op |  2325 allocs/op |
| BenchmarkPat_GithubAll         |       273 |   4381818 ns/op | 1483152 B/op | 26963 allocs/op |
| BenchmarkPossum_GithubAll      |     10000 |    164367 ns/op |   84448 B/op |   609 allocs/op |
| BenchmarkR2router_GithubAll    |     10000 |    160220 ns/op |   77328 B/op |   979 allocs/op |
| BenchmarkRivet_GithubAll       |     14625 |     82453 ns/op |   16272 B/op |   167 allocs/op |
| BenchmarkTango_GithubAll       |      6255 |    279611 ns/op |   63826 B/op |  1618 allocs/op |
| BenchmarkTigerTonic_GithubAll  |      2008 |    687874 ns/op |  193856 B/op |  4474 allocs/op |
| BenchmarkTraffic_GithubAll     |       355 |   3478508 ns/op |  820744 B/op | 14114 allocs/op |
| BenchmarkVulcan_GithubAll      |      6885 |    193333 ns/op |   19894 B/op |   609 allocs/op |

- (1): Total Repetitions achieved in constant time, higher means more confident result
- (2): Single Repetition Duration (ns/op), lower is better
- (3): Heap Memory (B/op), lower is better
- (4): Average Allocations per Repetition (allocs/op), lower is better

## Middlewares

You can find many useful Gin middlewares at [gin-contrib](https://github.com/gin-contrib).

## Users

Awesome project lists using [Gin](https://github.com/gokitcloud/ginkit) web framework.

- [gorush](https://github.com/appleboy/gorush): A push notification server written in Go.
- [fnproject](https://github.com/fnproject/fn): The container native, cloud agnostic serverless platform.
- [photoprism](https://github.com/photoprism/photoprism): Personal photo management powered by Go and Google TensorFlow.
- [lura](https://github.com/luraproject/lura): Ultra performant API Gateway with middlewares.
- [picfit](https://github.com/thoas/picfit): An image resizing server written in Go.
- [dkron](https://github.com/distribworks/dkron): Distributed, fault tolerant job scheduling system.

## Contributing

Gin is the work of hundreds of contributors. We appreciate your help!

Please see [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches and the contribution workflow.
