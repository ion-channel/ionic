package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestDependencies(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Dependencies", func() {
		var server *bogus.Bogus
		var h, p string
		var client *IonClient

		g.BeforeEach(func() {
			server = bogus.New()
			h, p = server.HostPort()
			client, _ = New(fmt.Sprintf("http://%v:%v", h, p))
		})

		g.It("should get the latest version of a dependency", func() {
			server.AddPath("/v1/dependency/getLatestVersionForDependency").
				SetMethods("GET").
				SetPayload([]byte(sampleLatestVersionResponse)).
				SetStatus(http.StatusOK)
			dep, err := client.GetLatestVersionForDependency("bundler", RubyEcosystem, "atoken")

			Expect(err).To(BeNil())
			Expect(dep.Version).To(Equal("1.16.3"))

			hrs := server.HitRecords()
			Expect(len(hrs)).To(Equal(1))
		})

		g.It("should get the latest version of a dependency", func() {
			server.AddPath("/v1/dependency/getLatestVersionsForDependency").
				SetMethods("GET").
				SetPayload([]byte(sampleLatestVersionsResponse)).
				SetStatus(http.StatusOK)
			deps, err := client.GetLatestVersionsForDependency("bundler", RubyEcosystem, "atoken")

			Expect(err).To(BeNil())
			Expect(deps[0].Version).To(Equal("1.16.3"))

			hrs := server.HitRecords()
			Expect(len(hrs)).To(Equal(1))
		})
	})
}

const (
	sampleLatestVersionResponse  = `{"data":{"version":"1.16.3"}}`
	sampleLatestVersionsResponse = `{"data":["1.16.3","1.16.2"]}`
)
