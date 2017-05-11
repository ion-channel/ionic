package bogus

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestPaths(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Paths", func() {
		var server *Bogus
		var host, port string

		g.BeforeEach(func() {
			server = New()
			server.Start()
			host, port = server.HostPort()
		})

		g.Describe("Root Path", func() {
			g.It("should allow setting the payload for the root path", func() {
				p := "some payload"
				server.SetPayload([]byte(p))

				resp, err := http.Get("http://" + net.JoinHostPort(host, port))
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(body)).To(Equal(p))
				Expect(server.Hits()).To(Equal(1))
			})

			g.It("should allow setting the return status for the root path", func() {
				s := http.StatusTeapot
				server.SetStatus(s)

				resp, err := http.Get("http://" + net.JoinHostPort(host, port))
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(s))
				Expect(server.Hits()).To(Equal(1))
			})

			g.It("should return the root payload for all paths if it is the only registered path", func() {
				p := "this is / payload"
				server.SetPayload([]byte(p))

				resp, err := http.Get("http://" + net.JoinHostPort(host, port))
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(body).To(Equal([]byte(p)))

				resp, err = http.Get("http://" + net.JoinHostPort(host, port) + "/foo/bar")
				Expect(err).NotTo(HaveOccurred())
				body, err = ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				resp.Body.Close()
				Expect(body).To(Equal([]byte(p)))
			})
		})

		g.Describe("Additional Paths", func() {
			g.It("should allow adding a new path", func() {
				p1 := "some other payload"
				s1 := http.StatusOK
				server.SetPayload([]byte(p1))
				server.SetStatus(s1)

				s2 := http.StatusCreated
				server.AddPath("/foo/bar").
					SetStatus(s2)

				resp, err := http.Get("http://" + net.JoinHostPort(host, port))
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(s1))
				Expect(server.Hits()).To(Equal(1))

				resp, err = http.Get("http://" + net.JoinHostPort(host, port) + "/foo/bar")
				Expect(err).NotTo(HaveOccurred())

				Expect(resp.StatusCode).To(Equal(s2))
				Expect(server.Hits()).To(Equal(2))
			})

			g.It("should return unique payloads per path", func() {
				p1 := "some other payload"
				server.SetPayload([]byte(p1))

				p2 := "foobar"
				server.AddPath("/foo/bar").
					SetPayload([]byte(p2))

				resp, err := http.Get("http://" + net.JoinHostPort(host, port))
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(body)).To(Equal(p1))
				Expect(server.Hits()).To(Equal(1))

				resp, err = http.Get("http://" + net.JoinHostPort(host, port) + "/foo/bar")
				Expect(err).NotTo(HaveOccurred())

				body, err = ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(body)).To(Equal(p2))
				Expect(server.Hits()).To(Equal(2))
			})

			g.It("should return the number of times a path has been hit", func() {
				p1 := "some other payload"
				server.SetPayload([]byte(p1))

				p2 := "foobar"
				server.AddPath("/foo/bar").
					SetPayload([]byte(p2))

				resp, err := http.Get("http://" + net.JoinHostPort(host, port))
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				Expect(server.Hits()).To(Equal(1))
				Expect(server.paths["/"].Hits()).To(Equal(1))

				resp, err = http.Get("http://" + net.JoinHostPort(host, port) + "/foo/bar")
				Expect(err).NotTo(HaveOccurred())

				Expect(server.Hits()).To(Equal(2))
				Expect(server.paths["/"].Hits()).To(Equal(1))
				Expect(server.paths["/foo/bar"].Hits()).To(Equal(1))
			})

			g.It("should respect registered paths when more than the root path is registered", func() {
				payload1 := "this is / payload"
				payload2 := "this is /foo/bar payload"

				server.SetPayload([]byte(payload1))
				server.AddPath("/foo/bar").
					SetPayload([]byte(payload2))

				resp, err := http.Get("http://" + net.JoinHostPort(host, port))
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(body)).To(Equal(payload1))

				resp, err = http.Get("http://" + net.JoinHostPort(host, port) + "/foo/bar")
				Expect(err).NotTo(HaveOccurred())

				body, err = ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(body)).To(Equal(payload2))
			})

			g.It("should return 404 on an unregistered path when there is more than one registered", func() {
				payload1 := "this is / payload"
				payload2 := "this is /foo/bar payload"

				server.SetPayload([]byte(payload1))
				server.AddPath("/foo/bar").
					SetPayload([]byte(payload2))

				resp, err := http.Get("http://" + net.JoinHostPort(host, port))
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(body)).To(Equal(payload1))

				resp, err = http.Get("http://" + net.JoinHostPort(host, port) + "/foo/bar/baz")
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusNotFound))

				body, err = ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(body)).To(Equal("Not Found"))
			})

			g.It("should return 404 for root path when not registerd and additional path is registered", func() {
				p := "root is not registered"
				server.AddPath("/no/root").
					SetPayload([]byte(p))

				resp, err := http.Get("http://" + net.JoinHostPort(host, port))
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(body)).To(Equal("Not Found"))
				Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
				Expect(server.Hits()).To(Equal(1))

				resp, err = http.Get("http://" + net.JoinHostPort(host, port) + "/no/root")
				Expect(err).NotTo(HaveOccurred())

				body, err = ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(body)).To(Equal(p))
				Expect(server.Hits()).To(Equal(2))
			})
		})

		g.Describe("Path Methods", func() {
			g.It("should allow setting methods for paths", func() {
				p := server.AddPath("/bar").
					SetPayload([]byte("foo")).
					SetMethods("GET")

				Expect(len(p.methods)).To(Equal(1))
				Expect(p.methods[0]).To(Equal("GET"))

				p = server.AddPath("/barbar").
					SetPayload([]byte("drinks")).
					SetMethods("get", "post", "put")

				Expect(len(p.methods)).To(Equal(3))
				Expect(p.methods[0]).To(Equal("GET"))
				Expect(p.methods[1]).To(Equal("POST"))
				Expect(p.methods[2]).To(Equal("PUT"))
			})

			g.It("should return true if a path has a method", func() {
				p := server.AddPath("/bar").
					SetMethods("GET")

				ok := p.hasMethod("GET")
				Expect(ok).To(BeTrue())

				p.SetMethods("GET", "PUT")

				ok = p.hasMethod("PUT")
				Expect(ok).To(BeTrue())

				ok = p.hasMethod("GET")
				Expect(ok).To(BeTrue())
			})

			g.It("should return true if a path has no methods and GET is checked for", func() {
				p := server.AddPath("/bar")

				ok := p.hasMethod("GET")
				Expect(ok).To(BeTrue())
			})

			g.It("should return false if a path does not have a method", func() {
				p := server.AddPath("/bar").
					SetMethods("GET", "PUT")

				ok := p.hasMethod("PUT")
				Expect(ok).To(BeTrue())

				ok = p.hasMethod("GET")
				Expect(ok).To(BeTrue())

				ok = p.hasMethod("POST")
				Expect(ok).To(BeFalse())
			})

			g.It("should return false if a path has a method other than GET and GET is checked for", func() {
				p := server.AddPath("/bar").
					SetMethods("PUT")

				ok := p.hasMethod("GET")
				Expect(ok).To(BeFalse())
			})

			g.It("shouldn't allow putting to a get path", func() {
				p := "foo"
				postData := "freakazoid"
				server.AddPath("/spacebar").
					SetPayload([]byte(p)).
					SetStatus(http.StatusOK).
					SetMethods("GET")

				req, err := http.NewRequest("PUT", "http://"+net.JoinHostPort(host, port)+"/spacebar", bytes.NewReader([]byte(postData)))
				Expect(err).NotTo(HaveOccurred())

				client := &http.Client{}
				resp, err := client.Do(req)
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(http.StatusForbidden))
			})

			g.It("should allow putting to a put path", func() {
				p := "foo"
				postData := "live long and prosper"
				server.AddPath("/force").
					SetPayload([]byte(p)).
					SetStatus(http.StatusOK).
					SetMethods("PUT")

				req, err := http.NewRequest("PUT", "http://"+net.JoinHostPort(host, port)+"/force", bytes.NewReader([]byte(postData)))
				Expect(err).NotTo(HaveOccurred())

				client := &http.Client{}
				resp, err := client.Do(req)
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(http.StatusCreated))
			})
		})
	})
}
