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
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get a team", func() {
			server.AddPath("/v1/teams/getTeam").
				SetMethods("GET").
				SetPayload([]byte(SampleValidTeam)).
				SetStatus(http.StatusOK)

			team, err := client.GetTeam("cd98e4e1-6926-4989-8ef8-f326cd5956fc", "atoken")
			Expect(err).To(BeNil())
			Expect(team.ID).To(Equal("cd98e4e1-6926-4989-8ef8-f326cd5956fc"))
			Expect(team.Name).To(Equal("ion-channel"))
		})

		g.It("should get teams", func() {
			server.AddPath("/v1/teams/getTeams").
				SetMethods("GET").
				SetPayload([]byte(fmt.Sprintf(`{"data":[%v,%v]}`, SampleValidTeam, SampleValidTeam))).
				SetStatus(http.StatusOK)

			ts, err := client.GetTeams("atoken")
			Expect(err).To(BeNil())
			Expect(len(ts)).To(Equal(2))
			rec := server.HitRecords()[len(server.HitRecords())-1]
			Expect(rec.Body).To(Equal([]byte("")))
			Expect(rec.Path).To(Equal("/v1/teams/getTeams"))
			Expect(rec.Header.Get("Authorization")).To(Equal("Bearer atoken"))
		})

		g.It("should create a team", func() {
			server.AddPath("/v1/teams/createTeam").
				SetMethods("POST").
				SetPayload([]byte(SampleCreateTeam)).
				SetStatus(http.StatusOK)

			opts := CreateTeamOptions{
				Name:     "test-team",
				POCName:  "Ion",
				POCEmail: "test@iontest.com",
			}
			team, err := client.CreateTeam(opts, "atoken")
			Expect(err).To(BeNil())
			Expect(team.ID).To(Equal("5c4a8a84-efa0-4357-91f6-9f9e95f7dd1a"))
			Expect(team.Name).To(Equal("test-team"))
			Expect(team.POCName).To(Equal("Ion"))
			Expect(team.POCEmail).To(Equal("test@iontest.com"))
		})
	})
}

const (
	SampleValidTeam  = `{"data":{"id":"cd98e4e1-6926-4989-8ef8-f326cd5956fc","created_at":"2016-09-09T22:06:49.487Z","updated_at":"2016-09-09T22:06:49.487Z","name":"ion-channel","sys_admin":true,"poc_name":"","poc_email":"","poc_name_hash":"","poc_email_hash":""}}`
	SampleCreateTeam = `{"data":{"id":"5c4a8a84-efa0-4357-91f6-9f9e95f7dd1a","created_at":"2018-01-05T23:59:58.160Z","updated_at":"2018-01-05T23:59:58.160Z","name":"test-team","sys_admin":false,"deleted_at":null,"poc_name":"Ion","poc_email":"test@iontest.com","poc_name_hash":"","poc_email_hash":"","delivery_location":"","access_key":"","secret_key":"","delivery_region":""}}`
)
