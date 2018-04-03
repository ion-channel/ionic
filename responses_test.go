package ionic

import (
	"net/http"
	"strconv"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestResponses(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Response", func() {
		g.It("should return a new response", func() {
			data := struct {
				Foo string `json:"foo"`
			}{
				Foo: "Bar",
			}
			meta := Meta{
				Copyright: "not yos",
				Authors:   []string{"us"},
			}
			status := http.StatusOK

			b, s := NewResponse(data, meta, status)
			Expect(string(b)).To(ContainSubstring("\"data\":"))
			Expect(string(b)).To(ContainSubstring("\"meta\":"))
			Expect(string(b)).To(ContainSubstring("\"foo\":"))
			Expect(string(b)).To(ContainSubstring("Bar"))
			Expect(string(b)).To(ContainSubstring("\"copyright\":"))
			Expect(string(b)).To(ContainSubstring("not yos"))
			Expect(s).To(Equal(status))
		})
		g.It("should make an ion response object", func(){
			data := struct {
				Foo string `json:"foo"`
			}{
				Foo: "Bar",
			}
			meta := Meta{
				Copyright: "not yos",
				Authors:   []string{"us"},
				Version:   "vtest",
			}
			response, err := makeIonResponse(data, meta)
			Expect(err).NotTo(HaveOccurred())
			Expect(response).NotTo(BeNil())
			Expect(string(response.Data)).To(Equal(`{"foo":"Bar"}`))
			Expect(response.Meta.Copyright).To(Equal("not yos"))
			Expect(response.Meta.Authors).To(HaveLen(1))
			Expect(response.Meta.Authors[0]).To(Equal("us"))
			Expect(response.Meta.Version).To(Equal("vtest"))
		})
		g.It("should add default Copyright value if missing", func(){
			data := struct {
				Foo string `json:"foo"`
			}{
				Foo: "Bar",
			}
			meta := Meta{
				Authors:   []string{"us"},
				Version:   "vtest",
			}
			response, err := makeIonResponse(data, meta)
			Expect(err).NotTo(HaveOccurred())
			Expect(response).NotTo(BeNil())
			Expect(string(response.Data)).To(Equal(`{"foo":"Bar"}`))
			Expect(response.Meta.Copyright).To(Equal("Copyright 2018 Selection Pressure LLC www.selectpress.net"))
			Expect(response.Meta.Authors).To(HaveLen(1))
			Expect(response.Meta.Authors[0]).To(Equal("us"))
			Expect(response.Meta.Version).To(Equal("vtest"))
		})
		g.It("should add default Authors value if missing", func(){
			data := struct {
				Foo string `json:"foo"`
			}{
				Foo: "Bar",
			}
			meta := Meta{
				Copyright: "not yos",
				Version:   "vtest",
			}
			response, err := makeIonResponse(data, meta)
			Expect(err).NotTo(HaveOccurred())
			Expect(response).NotTo(BeNil())
			Expect(string(response.Data)).To(Equal(`{"foo":"Bar"}`))
			Expect(response.Meta.Copyright).To(Equal("not yos"))
			Expect(response.Meta.Authors).To(HaveLen(1))
			Expect(response.Meta.Authors[0]).To(Equal("Ion Channel Dev Team"))
			Expect(response.Meta.Version).To(Equal("vtest"))
		})
		g.It("should add default Version value if missing", func(){
			data := struct {
				Foo string `json:"foo"`
			}{
				Foo: "Bar",
			}
			meta := Meta{
				Copyright: "not yos",
				Authors:   []string{"us"},
			}
			response, err := makeIonResponse(data, meta)
			Expect(err).NotTo(HaveOccurred())
			Expect(response).NotTo(BeNil())
			Expect(string(response.Data)).To(Equal(`{"foo":"Bar"}`))
			Expect(response.Meta.Copyright).To(Equal("not yos"))
			Expect(response.Meta.Authors).To(HaveLen(1))
			Expect(response.Meta.Authors[0]).To(Equal("us"))
			Expect(response.Meta.Version).To(Equal("v1"))
		})
	})

	g.Describe("Error Response", func() {
		g.It("should return a new error response", func() {
			msg := "foo error"
			fields := []string{"bar"}
			status := http.StatusUnauthorized

			b, s := NewErrorResponse(msg, fields, status)
			Expect(string(b)).To(ContainSubstring(msg))
			Expect(string(b)).To(ContainSubstring(fields[0]))
			Expect(string(b)).To(ContainSubstring(strconv.Itoa(status)))
			Expect(s).To(Equal(status))
		})
	})
}
