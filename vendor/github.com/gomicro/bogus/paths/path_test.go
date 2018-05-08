package paths

import (
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/penname"
	. "github.com/onsi/gomega"
)

func TestBogus(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Paths", func() {
		g.Describe("Payload & Status", func() {
			g.It("should set a payload", func() {
				payload := []byte("some payload")
				p := New().
					SetPayload(payload)

				Expect(string(p.payload)).To(Equal(string(payload)))
			})

			g.It("should set a status", func() {
				status := http.StatusTeapot
				p := New().
					SetStatus(status)

				Expect(p.status).To(Equal(status))
			})
		})

		g.Describe("Setting Headers", func() {
			g.It("should allow setting headers for paths", func() {
				headers := map[string]string{
					"Content-Type": "plain/text",
				}
				p := New().
					SetHeaders(headers)

				Expect(p.headers["Content-Type"]).To(Equal("plain/text"))
				Expect(p.headers).To(Equal(headers))
			})
		})

		g.Describe("Setting Methods", func() {
			g.It("should allow setting methods for paths", func() {
				p := New().
					SetMethods("GET")

				Expect(len(p.methods)).To(Equal(1))
				Expect(p.methods[0]).To(Equal("GET"))

				p.SetMethods("get", "post", "put")

				Expect(len(p.methods)).To(Equal(3))
				Expect(p.methods[0]).To(Equal("GET"))
				Expect(p.methods[1]).To(Equal("POST"))
				Expect(p.methods[2]).To(Equal("PUT"))
			})
		})

		g.Describe("Checking Methods", func() {
			g.It("should return true if a path has a method", func() {
				p := New().
					SetMethods("GET")

				ok := p.hasMethod("GET")
				Expect(ok).To(BeTrue())

				p.SetMethods("GET", "PUT")

				ok = p.hasMethod("PUT")
				Expect(ok).To(BeTrue())

				ok = p.hasMethod("GET")
				Expect(ok).To(BeTrue())
			})

			g.It("should return false if a path does not have a method", func() {
				p := New().
					SetMethods("GET", "PUT")

				ok := p.hasMethod("PUT")
				Expect(ok).To(BeTrue())

				ok = p.hasMethod("GET")
				Expect(ok).To(BeTrue())

				ok = p.hasMethod("POST")
				Expect(ok).To(BeFalse())
			})

			g.It("should return false if no methods are set", func() {
				p := New()

				ok := p.hasMethod("GET")
				Expect(ok).To(BeFalse())
			})
		})

		g.Describe("Handling Methods", func() {
			g.It("shouldn't allow putting to a get path", func() {
				p := New().
					SetMethods("GET")
				w := penname.New()
				r := &http.Request{
					Method: "PUT",
				}

				p.HandleRequest(w, r)
				Expect(string(w.WrittenHeaders)).To(Equal("Header: 403"))
			})

			g.It("should allow putting to a put path", func() {
				p := New().
					SetMethods("PUT")
				w := penname.New()
				r := &http.Request{
					Method: "PUT",
				}

				p.HandleRequest(w, r)
				Expect(string(w.WrittenHeaders)).To(Equal("Header: 200"))
			})
		})
	})
}
