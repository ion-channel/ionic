# Bogus
[![Build Status](https://travis-ci.org/gomicro/bogus.svg)](https://travis-ci.org/gomicro/bogus)
[![Go Reportcard](https://goreportcard.com/badge/github.com/gomicro/bogus)](https://goreportcard.com/report/github.com/gomicro/bogus)
[![GoDoc](https://godoc.org/github.com/gomicro/bogus?status.svg)](https://godoc.org/github.com/gomicro/bogus)
[![License](https://img.shields.io/github/license/gomicro/bogus.svg)](https://github.com/gomicro/bogus/blob/master/LICENSE.md)
[![Release](https://img.shields.io/github/release/gomicro/bogus.svg)](https://github.com/gomicro/bogus/releases/latest)

Bogus simplifies the creation of a mocked http server using the `net/http/httptest` package.  It allows the creation of one to many endpoints with unique responses.  The interactions of each endpoint are recorded for assertions.

# Requirements
Golang version 1.6 or higher

# Installation

```
go get github.com/gomicro/bogus
```

# Usage
See the [examples](https://godoc.org/github.com/gomicro/bogus#pkg-examples) within the docs for ways to use the library.

# Versioning
The library will be versioned in accordance with [Semver 2.0.0](http://semver.org).  See the [releases](https://github.com/gomicro/bogus/releases) section for the latest version.  Until version 1.0.0 the libary is considered to be unstable.

It is always highly recommended to vendor the version you are using.

# License
See [LICENSE.md](./LICENSE.md) for more information.
