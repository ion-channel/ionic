package ionic

import (
	"fmt"
	"net/url"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestClient(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Client", func() {
		g.It("should return a new client", func() {
			s := "foosecret"
			u := "http://google.com"
			cli, err := New(s, u)

			Expect(err).To(BeNil())
			Expect(cli.bearerToken).To(Equal(s))
		})

		g.It("should return an error on a bad url", func() {
			s := "foosecret"
			u := "http\\foo://google.com"
			cli, err := New(s, u)

			Expect(err).NotTo(BeNil())
			Expect(cli).To(BeNil())
		})

		g.It("should create a url with params nil", func() {
			e := "some/random/endpoint"
			b := "http://google.com"
			cli, _ := New("foosecret", b)

			u := cli.createURL(e, nil, nil)
			Expect(u.String()).To(Equal(fmt.Sprintf("%v/%v", b, e)))
		})

		g.It("should create a url with empty params", func() {
			e := "some/random/endpoint"
			b := "http://google.com"
			cli, _ := New("foosecret", b)
			p := &url.Values{}

			u := cli.createURL(e, p, nil)
			Expect(u.String()).To(Equal(fmt.Sprintf("%v/%v", b, e)))
		})

		g.It("should create a url with params", func() {
			e := "some/random/endpoint"
			b := "http://google.com"
			cli, _ := New("foosecret", b)
			p := &url.Values{}
			p.Set("foo", "bar")

			u := cli.createURL(e, p, nil)
			Expect(u.String()).To(Equal(fmt.Sprintf("%v/%v?%v", b, e, p.Encode())))
		})

		g.It("should create a url with pagination params", func() {
			e := "some/random/endpoint"
			b := "http://google.com"
			o := 21
			l := 100
			p := &Pagination{Offset: o, Limit: l}
			cli, _ := New("foosecret", b)

			u := cli.createURL(e, nil, p)
			Expect(u.String()).To(Equal(fmt.Sprintf("%v/%v?limit=%v&offset=%v", b, e, l, o)))
		})
	})
}
