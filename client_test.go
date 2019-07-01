package ionic

import (
	"fmt"
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
			u := "http://google.com"
			cli, err := New(u)

			Expect(err).To(BeNil())
			Expect(cli).NotTo(BeNil())
		})

		g.It("should return an error on a bad url", func() {
			u := "http://googl%8675309e\\house.com"
			cli, err := New(u)

			Expect(err).NotTo(BeNil())
			Expect(cli).To(BeNil())
		})
	})
}

var client = &IonClient{
	baseURL: nil,
	client:  nil,
}

func ExampleIonClient_customPaginationRange() {
	pages := &pagination.Pagination{Offset: 20, Limit: 100}

	vulns, err := client.GetVulnerabilities("ruby", "1.9.3", "sometoken", pages)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Vulnerabilities: %v\n", vulns)
}

func ExampleIonClient_defaultPaginationRange() {
	// nil for pagination will use the default set by the API and may vary for each object
	vulns, err := client.GetVulnerabilities("ruby", "1.9.3", "sometoken", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Vulnerabilities: %v\n", vulns)
}
