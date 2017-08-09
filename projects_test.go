package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestProjects(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Projects", func() {
		server := bogus.New()
		server.Start()
		h, p := server.HostPort()
		client, _ := New("", fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get a ruleset", func() {
			server.AddPath("/v1/projects/getProject").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProject)).
				SetStatus(http.StatusOK)

			project, err := client.GetProject("334c183d-4d37-4515-84c4-0d0ed0fb8db0", "bef86653-1926-4990-8ef8-5f26cd59d6fc")
			Expect(err).To(BeNil())
			Expect(project.ID).To(Equal("334c183d-4d37-4515-84c4-0d0ed0fb8db0"))
			Expect(project.Name).To(Equal("Statler"))
		})
	})
}

const (
	SampleValidProject = `{"data":{"active":true,"aliases":[],"branch":"master","chat_channel":"foo","created_at":"2016-08-29T17:38:40.401Z","deploy_key":null,"description":"Statler Travis CI testing","id":"334c183d-4d37-4515-84c4-0d0ed0fb8db0","key_fingerprint":"","name":"Statler","password":null,"poc_email":"","poc_email_hash":"","poc_name":"","poc_name_hash":"","ruleset_id":"f7583ed9-c939-4b51-a865-394cc8ddcffa","should_monitor":false,"source":"git@github.com:ion-channel/statler.git","tags":[],"team_id":"bef86653-1926-4990-8ef8-5f26cd59d6fc","type":"git","updated_at":"2017-05-22T18:00:54.982Z","username":null}}`
)
