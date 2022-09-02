**This (hopefully temporary) fork supports Hamlib v4.**

# Hamlib binding for Golang
[![Go Report Card](https://goreportcard.com/badge/github.com/xylo04/goHamlib)](https://goreportcard.com/report/github.com/xylo04/goHamlib)
[![Build Status](https://travis-ci.org/xylo04/goHamlib.svg?branch=master)](https://travis-ci.org/xylo04/goHamlib)
[![Coverage Status](https://coveralls.io/repos/github/xylo04/goHamlib/badge.svg?branch=master)](https://coveralls.io/github/xylo04/goHamlib?branch=master)

This is a [golang](https://golang.org) binding for
[Hamlib](http://hamlib.org). The binding has been hand written in order
to provide a golang idiomatic API but staying also close as possible to
[hamlib's C API](http://hamlib.sourceforge.net/manuals/3.0.1/index.html).
Except of Hamlib (C), goHamlib has no other external dependencies.

goHamlib is compatible with both, Hamlib 4.x.

You might also want to checkout [gorigctl](https://github.com/dh1tw/gorigctl)
which is a drop-in replacement for hamlib's rigctl, although it comes with a
cli GUI and a modern, open network interface (implementing the [shackbus](https://shackbus.org) specification using
[protocol buffers](https://developers.google.com/protocol-buffers/) and
[MQTT](http://mqtt.org)).

## Implementation status

Currently, goHamlib covers about 80% of the hamlib API. Basically (the rarely used) Channel API is still missing.

## Supported Platforms

goHamlib has been tested on the following platforms:

- AMD64
- i386
- ARMv7 (e.g. Raspberry Pi 2)
- ARMv8 (e.g. Raspberry Pi 3)

and the following operating Systems:

- Linux (Ubuntu, Raspian, Armbian)
- MacOS (Sierra)
- Windows 10

## License

This library is published under the permissive [MIT](https://choosealicense.com/licenses/mit/) license. Although please keep in
mind that Hamlib is published under the
[LGPL license](https://choosealicense.com/licenses/lgpl-3.0/).

## Installation & Compilation

goHamlib depends on Hamlib which can be either installed through your OS
packet manager or easily build from source. If you want the latest patches
and support for the latest rigs, you might want to compile the latest
[stable release](https://sourceforge.net/projects/hamlib/files/hamlib)
or directly from the [git repository](https://github.com/n0nb/hamlib).

Make sure that you have installed the [latest release](https://golang.org/dl/) of Go and set the GOPATH environment variable.

### Linux Dependencies

```bash
$ sudo apt-get install -y libhamlib2 libhamlib-dev
```

### MacOS Dependencies

You can install the hamlib package through [homebrew](https://brew.sh)

```bash
$ brew install hamlib
```

### Compilation

```bash
$ go get github.com/dl1igc/goHamlib
```

## Tests

Most of goHamlib's API is unit tested. In order to execute the unit tests,
run: 

```bash
$ cd $GOPATH/src/github.com/dl1igc/goHamlib
$ go test
```

## Documentation

goHamlib's API is documented at [go.dev](https://pkg.go.dev/github.com/xylo04/goHamlib)

## Example

Checkout the [dummyrig_test.go](https://github.com/xylo04/goHamlib/blob/master/dummyrig_test.go) in this
repository to see how to use goHamlib.
