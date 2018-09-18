package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestTeamUsers(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Teams", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get a team user", func() {
			server.AddPath("/v1/teamUsers/getTeamUser").
				SetMethods("GET").
				SetPayload([]byte(SampleValidTeamUser)).
				SetStatus(http.StatusOK)

			tu, err := client.GetTeamUser("team", "user", "atoken")
			Expect(err).To(BeNil())
			Expect(tu.TeamID).To(Equal("team"))
			Expect(tu.UserID).To(Equal("user"))
			Expect(tu.Role).To(Equal("paper"))
		})

		g.It("should create a team user", func() {
			server.AddPath("/v1/teamUsers/createTeamUser").
				SetMethods("POST").
				SetPayload([]byte(SampleCreateTeamUser)).
				SetStatus(http.StatusOK)

			opts := CreateTeamUserOptions{
				Role:   "admin",
				Status: "active",
				TeamID: "team",
				UserID: "user",
			}
			tu, err := client.CreateTeamUser(opts, "atoken")
			Expect(err).To(BeNil())
			Expect(tu.ID).To(Equal("teamuser"))
			Expect(tu.Role).To(Equal("admin"))
			Expect(tu.Status).To(Equal("active"))
			Expect(tu.TeamID).To(Equal("team"))
			Expect(tu.UserID).To(Equal("user"))
		})
	})
}

const (
	SampleValidTeamUser  = `{"data":{"id":"teamuser","created_at":"2016-09-09T22:06:49.487Z","updated_at":"2016-09-09T22:06:49.487Z","team_id":"team","user_id":"user","role":"paper"}}`
	SampleCreateTeamUser = `{"data":{"id":"teamuser","created_at":"2018-01-05T23:59:58.160Z","updated_at":"2018-01-05T23:59:58.160Z","team_id":"team","user_id":"user","role":"admin","status":"active"}}`
)
