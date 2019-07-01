package requests

import (
	"fmt"
	"net/url"
	"testing"

	. "github.com/franela/goblin"
	"github.com/ion-channel/ionic/pagination"
	. "github.com/onsi/gomega"
)

func TestClient(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Client", func() {
		g.It("should create a url with params nil", func() {
			e := "some/random/endpoint"
			b, _ := url.Parse("http://google.com")

			u := createURL(b, e, nil, nil)
			Expect(u.String()).To(Equal(fmt.Sprintf("%v/%v", b, e)))
		})

		g.It("should create a url with empty params", func() {
			e := "some/random/endpoint"
			b, _ := url.Parse("http://google.com")
			p := &url.Values{}

			u := createURL(b, e, p, nil)
			Expect(u.String()).To(Equal(fmt.Sprintf("%v/%v", b, e)))
		})

		g.It("should create a url with params", func() {
			e := "some/random/endpoint"
			b, _ := url.Parse("http://google.com")
			p := &url.Values{}
			p.Set("foo", "bar")

			u := createURL(b, e, p, nil)
			Expect(u.String()).To(Equal(fmt.Sprintf("%v/%v?%v", b, e, p.Encode())))
		})

		g.It("should create a url with pagination params", func() {
			e := "some/random/endpoint"
			b, _ := url.Parse("http://google.com")
			o := 21
			l := 100
			p := &pagination.Pagination{Offset: o, Limit: l}

			u := createURL(b, e, nil, p)
			Expect(u.String()).To(Equal(fmt.Sprintf("%v/%v?limit=%v&offset=%v", b, e, l, o)))
		})
	})
}
