# fn

[![Build Status](https://travis-ci.org/bcongdon/fn.svg?branch=master)](https://travis-ci.org/bcongdon/fn)
[![GoDoc](https://godoc.org/github.com/bcongdon/fn?status.svg)](https://godoc.org/github.com/bcongdon/fn)
[![Go Report Card](https://goreportcard.com/badge/github.com/bcongdon/fn)](https://goreportcard.com/report/github.com/bcongdon/fn)

Golang library for generating (mostly) unique date-sortable filenames


## Installation

Via `go get`:

```
go get github.com/bcongdon/fn
```

Via Go modules: Add the following to your `go.mod` file

```
github.com/bcongdon/fn v0.0.1
```

## Usage

As per inconvergent's [fn](https://github.com/inconvergent/fn), the formatting of file names is as follows:

```
yyyymmdd-hhmmss-gitsha-procdatetimesha
```

The first 2 chunks are the current date (day and time), the 3rd chunk is the hash of the current git commit, and the 4th chunk is the hash of the current process ID and time.

Optionally, a prefix/postfix can be added to the names by setting the `Prefix` and `Postfix` fields of `Fn`, respectively.

```go
fNamer := fn.New()

// Basic Name
fmt.Println(fNamer.Name())
// "200260220-072532-392d644-6193ecb1"

// Name w/ Prefix
fNamer.Prefix = "foo"
fmt.Println(fNamer.Name())
// "foo-200260220-072532-392d644-4dffcc5b"

// Name w/ Postfix
fNamer.Postfix = "bar"
fNamer.Prefix = ""
fmt.Println(fNamer.Name())
// "200260220-072532-392d644-c25334f9-bar"

// Name w/ file extension
fNamer.Postfix = ""
fmt.Println(fNamer.NameWithFileType("png"))
// "200260220-072532-392d644-a165c554.png"
```

A full example lives in the `examples/` directory.

## Attribution

This library is more-or-less a direct port of [inconvergent](https://inconvergent.net/)'s [fn](https://github.com/inconvergent/fn) Python package.