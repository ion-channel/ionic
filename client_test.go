package ionic

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
		g.It("should return a new client", func() {
			s := "foosecret"
			u := "http://google.com"
			cli, err := New(s, u)

			Expect(err).To(BeNil())
			Expect(cli.bearerToken).To(Equal(s))
		})

		g.It("should return an error on a bad url", func() {
			s := "foosecret"
			u := "http://googl%8675309e\\house.com"
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
			p := &pagination.Pagination{Offset: o, Limit: l}
			cli, _ := New("foosecret", b)

			u := cli.createURL(e, nil, p)
			Expect(u.String()).To(Equal(fmt.Sprintf("%v/%v?limit=%v&offset=%v", b, e, l, o)))
		})

		g.It("should set a new token", func() {
			old := "footoken"
			new := "bartoken"

			cli, _ := New(old, "http://google.com")
			Expect(cli.bearerToken).To(Equal(old))
			cli.SetBearerToken(new)
			Expect(cli.bearerToken).To(Equal(new))
		})
	})
}

var client = &IonClient{
	baseURL:     nil,
	bearerToken: "footoken",
	client:      nil,
}

func ExamplePagination_customRange() {
	pages := &pagination.Pagination{Offset: 20, Limit: 100}

	vulns, err := client.GetVulnerabilities("ruby", "1.9.3", pages)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Vulnerabilities: %v\n", vulns)
}

func ExamplePagination_defaultRange() {
	// nil for pagination will use the default set by the API and may vary for each object
	vulns, err := client.GetVulnerabilities("ruby", "1.9.3", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Vulnerabilities: %v\n", vulns)
}
