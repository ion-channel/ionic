package bogus

import (
	"bytes"
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
			server.Start()
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
		})
	})
}
