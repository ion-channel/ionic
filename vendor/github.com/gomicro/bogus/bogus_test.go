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

func TestBogus(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Bogus Functions", func() {
		var server *Bogus
		var host, port string

		g.BeforeEach(func() {
			server = New()
			host, port = server.HostPort()
		})

		g.AfterEach(func() {
			server.Close()
		})

		g.It("should return the host and port", func() {
			Expect(host).To(Equal("127.0.0.1"))
			Expect(port).ToNot(Equal(""))
		})

		g.It("should count hits against the server", func() {
			resp, err := http.Get("http://" + net.JoinHostPort(host, port))
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()
			Expect(server.Hits()).To(Equal(1))
		})

		g.It("should track full hit records", func() {
			resp, err := http.Get("http://" + net.JoinHostPort(host, port))
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			resp, err = http.Post(
				"http://"+net.JoinHostPort(host, port)+"/foo/bar",
				"application/octet-stream",
				bytes.NewBuffer([]byte("post body")))
			Expect(err).NotTo(HaveOccurred())

			req, err := http.NewRequest(
				"PUT",
				"http://"+net.JoinHostPort(host, port)+"/foo/bar/cool?foo=bar&baz=fiz",
				bytes.NewBuffer([]byte("put body")))
			Expect(err).NotTo(HaveOccurred())
			req.Header.Add("X-CustomHeader01", "custom value 01")
			client := http.Client{}
			resp, err = client.Do(req)
			Expect(err).NotTo(HaveOccurred())

			Expect(server.HitRecords()).To(HaveLen(3))

			firstHit := server.HitRecords()[0]
			Expect(firstHit.Verb).To(Equal("GET"))
			Expect(firstHit.Path).To(Equal("/"))
			Expect(firstHit.Query).To(HaveLen(0))
			Expect(firstHit.Body).To(HaveLen(0))

			secondHit := server.HitRecords()[1]
			Expect(secondHit.Verb).To(Equal("POST"))
			Expect(secondHit.Path).To(Equal("/foo/bar"))
			Expect(secondHit.Query).To(HaveLen(0))
			Expect(string(secondHit.Body)).To(Equal("post body"))

			thirdHit := server.HitRecords()[2]
			Expect(thirdHit.Verb).To(Equal("PUT"))
			Expect(thirdHit.Path).To(Equal("/foo/bar/cool"))
			Expect(thirdHit.Query).To(HaveLen(2))
			Expect(thirdHit.Query["foo"]).To(HaveLen(1))
			Expect(thirdHit.Query["foo"][0]).To(Equal("bar"))
			Expect(thirdHit.Query["baz"]).To(HaveLen(1))
			Expect(thirdHit.Query["baz"][0]).To(Equal("fiz"))
			Expect(string(thirdHit.Body)).To(Equal("put body"))
			Expect(thirdHit.Header.Get("X-CustomHeader01")).To(Equal("custom value 01"))
		})

		g.It("should allow adding a new path", func() {
			p1 := "some other payload"
			s1 := http.StatusOK
			server.AddPath("/").
				SetMethods("GET").
				SetPayload([]byte(p1)).
				SetStatus(s1)

			resp, err := http.Get("http://" + net.JoinHostPort(host, port))
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			Expect(resp.StatusCode).To(Equal(s1))
			Expect(server.Hits()).To(Equal(1))
		})

		g.It("should return unique payloads per path", func() {
			p1 := "some other payload"
			server.AddPath("/").
				SetMethods("GET").
				SetPayload([]byte(p1))

			p2 := "foobar"
			server.AddPath("/foo/bar").
				SetMethods("GET").
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

		g.It("should return 404 on an unregistered path", func() {
			server.AddPath("/foo/bar").
				SetPayload([]byte("this is /foo/bar payload"))

			resp, err := http.Get("http://" + net.JoinHostPort(host, port) + "/foo/bar/baz")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(body)).To(Equal("Not Found"))
		})

		g.It("should return unique headers per path", func() {
			headers1 := map[string]string{"Content-Type": "plain/text"}
			server.AddPath("/").
				SetMethods("GET").
				SetHeaders(headers1)

			headers2 := map[string]string{"Content-Type": "application/json"}
			server.AddPath("/foo/bar").
				SetMethods("GET").
				SetHeaders(headers2)

			resp, err := http.Get("http://" + net.JoinHostPort(host, port))
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			Expect(resp.Header.Get("Content-Type")).To(Equal("plain/text"))

			resp, err = http.Get("http://" + net.JoinHostPort(host, port) + "/foo/bar")
			Expect(err).NotTo(HaveOccurred())

			Expect(resp.Header.Get("Content-Type")).To(Equal("application/json"))
		})
	})
}
