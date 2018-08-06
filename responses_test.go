package ionic

import (
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/penname"
	. "github.com/onsi/gomega"
)

func TestResponses(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Response", func() {
		g.Describe("Construction", func() {
			g.It("should return a new response with the defaults set", func() {
				d := data{Name: "foo"}

				r, err := NewResponse(d, Meta{}, http.StatusOK)
				Expect(err).To(BeNil())
				Expect(r.Meta.Copyright).To(ContainSubstring("Selection Pressure LLC"))
				Expect(len(r.Meta.Authors)).To(Equal(1))
				Expect(r.Meta.Authors[0]).To(Equal("Ion Channel Dev Team"))
				Expect(r.Meta.Version).To(Equal("v1"))
				Expect(string(r.Data)).To(ContainSubstring("\"name\":\"foo\""))
			})
		})

		g.Describe("Writing", func() {
			var mw *penname.PenName
			g.BeforeEach(func() {
				mw = penname.New()
			})

			g.It("should write a response", func() {
				d := data{Name: "foo"}
				r, _ := NewResponse(d, Meta{}, http.StatusOK)

				r.WriteResponse(mw)

				Expect(string(mw.WrittenHeaders())).To(ContainSubstring("Header: 200"))
				Expect(string(mw.Written())).To(ContainSubstring(`"data":{"name":"foo"}`))
				Expect(string(mw.Written())).To(ContainSubstring(`"copyright":"Copyright 2018 Selection Pressure LLC www.selectpress.net"`))
				Expect(string(mw.Written())).To(ContainSubstring(`"authors":["Ion Channel Dev Team"]`))
				Expect(string(mw.Written())).To(ContainSubstring(`"version":"v1"`))
				Expect(string(mw.Written())).To(ContainSubstring(`"total_count":0,"offset":0`))
			})
		})
	})

	g.Describe("Error Response", func() {
		g.Describe("Construction", func() {
			g.It("should return a new error response", func() {
				msg := "foo error"
				fields := []string{"bar"}
				status := http.StatusUnauthorized

				er := NewErrorResponse(msg, fields, status)
				Expect(er.Message).To(Equal(msg))
				Expect(len(er.Fields)).To(Equal(len(fields)))
				Expect(er.Code).To(Equal(status))
			})
		})

		g.Describe("Writing", func() {
			var mw *penname.PenName
			g.BeforeEach(func() {
				mw = penname.New()
			})

			g.It("should write an error response", func() {
				er := NewErrorResponse("something went wrong", nil, http.StatusUnauthorized)

				er.WriteResponse(mw)

				Expect(string(mw.WrittenHeaders())).To(ContainSubstring("Header: 401"))
				Expect(string(mw.Written())).To(ContainSubstring(`"message":"something went wrong"`))
				Expect(string(mw.Written())).To(ContainSubstring(`"code":401`))
			})
		})
	})
}

type data struct {
	Name string `json:"name"`
}
