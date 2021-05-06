package ionic

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestAliases(t *testing.T) {
	g := goblin.Goblin(t)
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

			aao := AddAliasOptions{
				Name:      "name",
				ProjectID: "f9bca953-80ac-46c4-b195-d37f3bc4f498",
				TeamID:    "ateamid",
				Version:   "version",
			}

			alias, err := client.AddAlias(aao, "token")
			Expect(err).To(BeNil())
			Expect(alias.ID).To(Equal("334c183d-4d37-4515-84c4-0d0ed0fb8db0"))
			Expect(alias.Name).To(Equal("name"))
			Expect(alias.Version).To(Equal("version"))

			Expect(server.Hits()).To(Equal(1))
			hr := server.HitRecords()[0]
			Expect(hr.Query.Get("project_id")).To(Equal("f9bca953-80ac-46c4-b195-d37f3bc4f498"))
		})

	})
}

const (
	SampleValidAliasedProject = `{"data":{"id":"334c183d-4d37-4515-84c4-0d0ed0fb8db0","version":"version","name":"name"}}`
)
