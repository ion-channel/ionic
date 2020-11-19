package responses

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	goblin "github.com/franela/goblin"
	"github.com/gomicro/penname"
	. "github.com/onsi/gomega"
)

func TestResponses(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Response", func() {
		g.Describe("Construction", func() {
			g.It("should return a new response with the defaults set", func() {
				d := data{Name: "foo"}
				r, err := NewResponse(d, Meta{}, http.StatusOK)
				Expect(err).To(BeNil())
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
				Expect(string(mw.Written())).To(ContainSubstring(`"total_count":0,"offset":0`))
				Expect(string(mw.Written())).ToNot(ContainSubstring(`"last_updated"`))
			})

			g.It("should write a response with the appropriate Meta fields", func() {
				d := data{Name: "foo"}
				responseWithBlankMeta, _ := NewResponse(d, Meta{}, http.StatusOK)

				responseWithBlankMeta.WriteResponse(mw)

				Expect(string(mw.Written())).ToNot(ContainSubstring(`"last_updated"`))
				Expect(string(mw.Written())).ToNot(ContainSubstring(`"limit"`))

				d = data{Name: "foo"}
				time := time.Now()
				responseWithIncludedMeta, _ := NewResponse(d, Meta{LastUpdate: &time, Limit: 1}, http.StatusOK)

				responseWithIncludedMeta.WriteResponse(mw)

				Expect(string(mw.Written())).To(ContainSubstring(`"last_update"`))
				Expect(string(mw.Written())).To(ContainSubstring(`"limit"`))
			})

		})
	})

	g.Describe("Error Response", func() {
		g.Describe("Construction", func() {
			g.It("should return a new error response", func() {
				msg := "foo error"
				fields := map[string]string{
					"bar": "its foo-ed up",
				}
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

		g.Describe("Fields", func() {
			g.It("should format the fields nicely", func() {
				fields := map[string]string{
					"source": "required field",
					"type":   "invalid type",
					"name":   "required field",
				}
				er := NewErrorResponse("something went wrong", fields, http.StatusUnauthorized)

				str := fmt.Sprint(er.Fields)
				Expect(str[0:1]).To(Equal("["))
				Expect(str[len(str)-1:]).To(Equal("]"))
				Expect(str).To(ContainSubstring("source: required field"))
				Expect(str).To(ContainSubstring("type: invalid type"))
				Expect(str).To(ContainSubstring("name: required field"))
			})
		})
	})
}

type data struct {
	Name string `json:"name"`
}
