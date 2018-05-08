package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestAliases(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Alias", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should add an alias", func() {
			server.AddPath("/v1/project/addAlias").
				SetMethods("POST").
				SetPayload([]byte(SampleValidAliasedProject)).
				SetStatus(http.StatusOK)

			alias, err := client.AddAlias("f9bca953-80ac-46c4-b195-d37f3bc4f498", "ateamid", "name", "version", "token")
			Expect(err).To(BeNil())
			Expect(alias.ID).To(Equal("334c183d-4d37-4515-84c4-0d0ed0fb8db0"))
			Expect(alias.Name).To(Equal("name"))
			Expect(alias.Version).To(Equal("version"))
		})

	})
}

const (
	SampleValidAliasedProject = `{"data":{"id":"334c183d-4d37-4515-84c4-0d0ed0fb8db0","version":"version","name":"name"}}`
)
