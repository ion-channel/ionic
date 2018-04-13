package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestTags(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Tags", func() {
		server := bogus.New()
		server.Start()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get a tag", func() {
			server.AddPath("/v1/tag/getTag").
				SetMethods("GET").
				SetPayload([]byte(SampleValidTag)).
				SetStatus(http.StatusOK)

			tag, err := client.GetTag("27A06D70-A8AA-4532-ABBF-C83ADF49F855", "A3EB1676-C91A-4FCF-AE1D-9F887C0D4B66", "atoken")
			Expect(err).To(BeNil())
			Expect(tag.ID).To(Equal("27A06D70-A8AA-4532-ABBF-C83ADF49F855"))
			Expect(tag.Name).To(Equal("Jenkins"))
		})

	})
}

const (
	SampleValidTag = `{"data":{"id":"27A06D70-A8AA-4532-ABBF-C83ADF49F855","team_id":"A3EB1676-C91A-4FCF-AE1D-9F887C0D4B66","name":"Jenkins","description":null,"created_at":"2017-07-10T18:39:53.914Z","updated_at":"2017-07-10T18:39:53.914Z"}}`
)
