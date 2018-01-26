package ionic

import (
	"encoding/json"
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
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get a project", func() {
			server.AddPath("/v1/project/getProject").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProject)).
				SetStatus(http.StatusOK)

			project, err := client.GetProject("334c183d-4d37-4515-84c4-0d0ed0fb8db0", "bef86653-1926-4990-8ef8-5f26cd59d6fc", "")
			Expect(err).To(BeNil())
			Expect(project.ID).To(Equal("334c183d-4d37-4515-84c4-0d0ed0fb8db0"))
			Expect(project.Name).To(Equal("Statler"))
		})

		g.It("should get a raw project", func() {
			server.AddPath("/v1/project/getProject").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProject)).
				SetStatus(http.StatusOK)

			raw, err := client.GetRawProject("334c183d-4d37-4515-84c4-0d0ed0fb8db0", "bef86653-1926-4990-8ef8-5f26cd59d6fc", "")
			Expect(err).To(BeNil())
			Expect(raw).To(Equal(json.RawMessage(SampleValidRawProject)))
		})

		g.It("should get all projects", func() {
			server.AddPath("/v1/project/getProjects").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProjects)).
				SetStatus(http.StatusOK)

			projects, err := client.GetProjects("bef86653-1926-4990-8ef8-5f26cd59d6fc", "", nil)
			Expect(err).To(BeNil())
			Expect(len(projects)).To(Equal(1))
			Expect(projects[0].ID).To(Equal("334c183d-4d37-4515-84c4-0d0ed0fb8db0"))
			Expect(projects[0].Name).To(Equal("Statler"))
		})
	})
}

const (
	SampleValidProject    = `{"data":{"active":true,"aliases":[],"branch":"master","chat_channel":"foo","created_at":"2016-08-29T17:38:40.401Z","deploy_key":null,"description":"Statler Travis CI testing","id":"334c183d-4d37-4515-84c4-0d0ed0fb8db0","key_fingerprint":"","name":"Statler","password":null,"poc_email":"","poc_email_hash":"","poc_name":"","poc_name_hash":"","ruleset_id":"f7583ed9-c939-4b51-a865-394cc8ddcffa","should_monitor":false,"source":"git@github.com:ion-channel/statler.git","tags":[],"team_id":"bef86653-1926-4990-8ef8-5f26cd59d6fc","type":"git","updated_at":"2017-05-22T18:00:54.982Z","username":null}}`
	SampleValidRawProject = `{"active":true,"aliases":[],"branch":"master","chat_channel":"foo","created_at":"2016-08-29T17:38:40.401Z","deploy_key":null,"description":"Statler Travis CI testing","id":"334c183d-4d37-4515-84c4-0d0ed0fb8db0","key_fingerprint":"","name":"Statler","password":null,"poc_email":"","poc_email_hash":"","poc_name":"","poc_name_hash":"","ruleset_id":"f7583ed9-c939-4b51-a865-394cc8ddcffa","should_monitor":false,"source":"git@github.com:ion-channel/statler.git","tags":[],"team_id":"bef86653-1926-4990-8ef8-5f26cd59d6fc","type":"git","updated_at":"2017-05-22T18:00:54.982Z","username":null}`
	SampleValidProjects   = `{"data":[{"active":true,"aliases":[],"branch":"master","chat_channel":"foo","created_at":"2016-08-29T17:38:40.401Z","deploy_key":null,"description":"Statler Travis CI testing","id":"334c183d-4d37-4515-84c4-0d0ed0fb8db0","key_fingerprint":"","name":"Statler","password":null,"poc_email":"","poc_email_hash":"","poc_name":"","poc_name_hash":"","ruleset_id":"f7583ed9-c939-4b51-a865-394cc8ddcffa","should_monitor":false,"source":"git@github.com:ion-channel/statler.git","tags":[],"team_id":"bef86653-1926-4990-8ef8-5f26cd59d6fc","type":"git","updated_at":"2017-05-22T18:00:54.982Z","username":null}]}`
)
