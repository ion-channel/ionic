# Bogus
[![Build Status](https://travis-ci.org/gomicro/bogus.svg)](https://travis-ci.org/gomicro/bogus)
[![Coverage](http://gocover.io/_badge/github.com/gomicro/bogus)](http://gocover.io/github.com/gomicro/bogus)
[![Go Reportcard](https://goreportcard.com/badge/github.com/gomicro/bogus)](https://goreportcard.com/report/github.com/gomicro/bogus)
[![GoDoc](https://godoc.org/github.com/gomicro/bogus?status.png)](https://godoc.org/github.com/gomicro/bogus)

Bogus simplifies the creation of a mocked http server using the `net/http/httptest` package.  It allows the creation of one to many endpoints with unique responses.  The interactions of each endpoint are recorded for assertions.

# Usage

Setting a payload and status against a root path

```
import "github.com/gomicro/bogus"

...

	g.Describe("Tests needing a test server", func(){
		var server *bogus.Bogus

		g.BeforeEach(func(){
			server = bogus.New()
			server.SetPayload([]byte("some return payload"))
			server.SetStatus(200)
			server.Start()
		})

		g.It("should connect to a test server", func(){
			host, port := server.HostPort()

			...

			Expect(server.Hits()).To(Equal(1))
		})
	})
```

Setting a payload and status against a specific path

```
import "github.com/gomicro/bogus"

...

	g.Describe("Tests needing a test server", func(){
		var server *bogus.Bogus

		g.BeforeEach(func(){
			server = bogus.New()
			server.Start()
		})

		g.It("should connect to a test server", func(){
			server.AddPath("/foo/bar").
				SetPayload([]byte("some return payload")).
				SetStatus(200)
			host, port := server.HostPort()

			...

			Expect(server.Hits()).To(Equal(1))
		})
	})
```
