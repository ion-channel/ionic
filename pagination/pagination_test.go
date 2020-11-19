package pagination

import (
	"net/http"
	"net/url"
	"testing"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestPagination(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Pagination", func() {
		g.Describe("Instantiation", func() {
			g.It("should return a new pagination object", func() {
				o := 20
				l := 10
				p := New(o, l)

				Expect(p.Limit).To(Equal(l))
				Expect(p.Offset).To(Equal(o))
			})

			g.It("should set limit to 10 if not positive", func() {
				o := 20
				l := -100
				p := New(o, l)

				Expect(p.Limit).NotTo(Equal(l))
				Expect(p.Limit).To(Equal(10))
				Expect(p.Offset).To(Equal(o))
			})

			g.It("should prevent offset from being negative", func() {
				o := -1
				l := 10
				p := New(o, l)

				Expect(p.Limit).To(Equal(l))
				Expect(p.Offset).NotTo(Equal(o))
				Expect(p.Offset).To(Equal(0))
			})
		})

		g.Describe("URL Params", func() {
			g.It("should append the url pagination params", func() {
				params := &url.Values{}
				params.Set("foo", "bar")
				p := Pagination{Limit: 10, Offset: 20}

				p.AddParams(params)
				Expect(params.Get("foo")).To(Equal("bar"))
				Expect(params.Get("limit")).To(Equal("10"))
				Expect(params.Get("offset")).To(Equal("20"))
			})
		})

		g.Describe("Incrementing", func() {
			g.It("should increment the offset up", func() {
				l := 10
				o := 15
				p := Pagination{Limit: l, Offset: o}

				p.Up()
				Expect(p.Offset).To(Equal(o + l))
				Expect(p.Limit).To(Equal(l))
			})

			g.It("should ignore incrementing up if all items", func() {
				l := AllItems.Limit
				o := AllItems.Offset

				AllItems.Up()
				Expect(AllItems.Limit).To(Equal(l))
				Expect(AllItems.Offset).To(Equal(o))
			})

			g.It("should increment the offset down", func() {
				l := 10
				o := 15
				p := Pagination{Limit: l, Offset: o}

				p.Down()
				Expect(p.Offset).To(Equal(o - l))
				Expect(p.Limit).To(Equal(l))
			})

			g.It("should not increment past 0", func() {
				l := 10
				o := 15
				p := Pagination{Limit: l, Offset: o}

				p.Down()
				Expect(p.Offset).To(Equal(o - l))
				Expect(p.Limit).To(Equal(l))
				p.Down()
				Expect(p.Offset).To(Equal(0))
				Expect(p.Limit).To(Equal(l))
			})

			g.It("should ignore incrementing up if all items", func() {
				l := AllItems.Limit
				o := AllItems.Offset

				AllItems.Down()
				Expect(AllItems.Limit).To(Equal(l))
				Expect(AllItems.Offset).To(Equal(o))
			})
		})

		g.Describe("SQLification", func() {
			g.It("should convert to a sql string", func() {
				p := Pagination{Limit: 10, Offset: 15}

				sql := p.SQL()
				Expect(sql).To(ContainSubstring("OFFSET 15"))
				Expect(sql).To(ContainSubstring("LIMIT 10"))
			})

			g.It("should leave limit off the sql string if 0", func() {
				p := Pagination{Limit: 0, Offset: 15}

				sql := p.SQL()
				Expect(sql).To(ContainSubstring("OFFSET 15"))
				Expect(sql).NotTo(ContainSubstring("LIMIT"))
			})

			g.It("should be blank when all items", func() {
				sql := AllItems.SQL()
				Expect(sql).To(Equal(""))
			})
		})

		g.Describe("Parsing", func() {
			g.It("should parse offset and limit from query params", func() {
				u, _ := url.Parse("http://localhost/something?offset=10&limit=100")
				req := &http.Request{
					URL: u,
				}

				p := ParseFromRequest(req)
				Expect(p.Offset).To(Equal(10))
				Expect(p.Limit).To(Equal(100))
			})

			g.It("should parse pagination params from the headers", func() {
				req := &http.Request{
					Header: http.Header{},
				}
				req.Header.Set("Offset", "100")
				req.Header.Set("Limit", "50")

				p := ParseFromRequest(req)
				Expect(p.Offset).To(Equal(100))
				Expect(p.Limit).To(Equal(50))
			})

			g.It("should return the default pagination values when nothing is given", func() {
				p := ParseFromRequest(&http.Request{})
				Expect(p.Offset).To(Equal(DefaultOffset))
				Expect(p.Limit).To(Equal(DefaultLimit))
			})

			g.It("should assume no offset when the offset is not in the request", func() {
				req := &http.Request{
					Header: http.Header{},
				}
				req.Header.Set("Limit", "50")

				p := ParseFromRequest(req)
				Expect(p.Offset).To(Equal(0))
				Expect(p.Limit).To(Equal(50))
			})

			g.It("should use a default when offset is not a number", func() {
				req := &http.Request{
					Header: http.Header{},
				}
				req.Header.Set("Offset", "destroy")
				req.Header.Set("Limit", "51")

				p := ParseFromRequest(req)
				Expect(p.Offset).To(Equal(0))
				Expect(p.Limit).To(Equal(51))
			})

			g.It("should assume the default limit when the limit is not in the request", func() {
				req := &http.Request{
					Header: http.Header{},
				}
				req.Header.Set("Offset", "100")

				p := ParseFromRequest(req)
				Expect(p.Offset).To(Equal(100))
				Expect(p.Limit).To(Equal(10))
			})

			g.It("should use a default when limit is not a number", func() {
				req := &http.Request{
					Header: http.Header{},
				}
				req.Header.Set("Offset", "101")
				req.Header.Set("Limit", "war_doctor")

				p := ParseFromRequest(req)
				Expect(p.Offset).To(Equal(101))
				Expect(p.Limit).To(Equal(10))
			})
		})
	})
}
