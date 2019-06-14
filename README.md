[![CircleCI](https://circleci.com/gh/spatialcurrent/go-flat/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-flat/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-flat)](https://goreportcard.com/report/spatialcurrent/go-flat)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-flat?status.svg)](https://godoc.org/github.com/spatialcurrent/go-flat) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-flat/blob/master/LICENSE)

# go-flat

# Description

**go-flat** is a simple library that provides functions to recursively flatten a slice of slices.

# Usage

**Go**

You can import **go-flat** as a library with:

```go
import (
  "github.com/spatialcurrent/go-flat/pkg/flat"
)
```

See [flat](https://godoc.org/github.com/spatialcurrent/go-flat/pkg/flat) in GoDoc for information on how to use Go API.  See the tests for ways to use this library.

# Examples

See [examples](https://godoc.org/github.com/spatialcurrent/go-flat/pkg/flat) in GoDoc.

# Testing

Run test using `make test` or (`bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-flat/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
