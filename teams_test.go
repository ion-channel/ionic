package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestTeams(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Teams", func() {
		server := bogus.New()
		server.Start()
		h, p := server.HostPort()
		client, _ := New("", fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get a team", func() {
			server.AddPath("/v1/teams/getTeam").
				SetMethods("GET").
				SetPayload([]byte(SampleValidTeam)).
				SetStatus(http.StatusOK)

			team, err := client.GetTeam("cd98e4e1-6926-4989-8ef8-f326cd5956fc")
			Expect(err).To(BeNil())
			Expect(team.ID).To(Equal("cd98e4e1-6926-4989-8ef8-f326cd5956fc"))
			Expect(team.Name).To(Equal("ion-channel"))
		})
	})
}

const (
	SampleValidTeam = `{"data":{"id":"cd98e4e1-6926-4989-8ef8-f326cd5956fc","created_at":"2016-09-09T22:06:49.487Z","updated_at":"2016-09-09T22:06:49.487Z","name":"ion-channel","sys_admin":true,"poc_name":"","poc_email":"","poc_name_hash":"","poc_email_hash":""}}`
)
