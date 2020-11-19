package ionic

import (
	"fmt"
	"net/http"
	"testing"

	goblin "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestUsers(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Users", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get a users self", func() {
			server.AddPath("/v1/users/getSelf").
				SetMethods("GET").
				SetPayload([]byte(SampleSelfResponse)).
				SetStatus(http.StatusOK)

			me, err := client.GetSelf("atoken")
			Expect(err).To(BeNil())
			Expect(me.Email).To(Equal("admin@ion.io"))
			Expect(me.Username).To(Equal("ion"))
			Expect(me.SysAdmin).To(Equal(true))
		})

		g.It("should get a user", func() {
			server.AddPath("/v1/users/getUser").
				SetMethods("GET").
				SetPayload([]byte(SampleSelfResponse)).
				SetStatus(http.StatusOK)

			me, err := client.GetUser("user", "atoken")
			Expect(err).To(BeNil())
			Expect(server.HitRecords()[len(server.HitRecords())-1].Query.Get("id")).To(Equal("user"))
			Expect(me.Email).To(Equal("admin@ion.io"))
			Expect(me.Username).To(Equal("ion"))
			Expect(me.SysAdmin).To(Equal(true))
		})

		g.It("should get all users", func() {
			server.AddPath("/v1/users/getUsers").
				SetMethods("GET").
				SetPayload([]byte(fmt.Sprintf(`{"data": [%v,%v]}`, SampleSelfResponse, SampleSelfResponse))).
				SetStatus(http.StatusOK)

			us, err := client.GetUsers("atoken")
			Expect(err).To(BeNil())
			Expect(len(us)).To(Equal(2))
		})

		g.It("should create a user", func() {
			server.AddPath("/v1/users/createUser").
				SetMethods("POST").
				SetPayload([]byte(SampleCreatedUser)).
				SetStatus(http.StatusOK)

			email := "testuser@test.com"
			username := "tester"
			u, err := client.CreateUser(email, username, "123456", "token")
			Expect(err).To(BeNil())
			Expect(u.Email).To(Equal(email))
			Expect(u.Username).To(Equal(username))
			Expect(u.SysAdmin).To(Equal(false))
		})

		g.It("should get user names", func() {
			server.AddPath("/v1/users/getUserNames").
				SetMethods("POST").
				SetPayload([]byte(SampleUserNames)).
				SetStatus(http.StatusOK)

			un, err := client.GetUserNames([]string{"276bbec3-cc77-44b9-a46d-c7760947ec9d", "B061D58B-FDFD-46BF-A766-2D38DE3B1D7B"}, "teamID", "sometoken")
			Expect(err).To(BeNil())
			Expect(len(un)).To(Equal(2))
			Expect(un[0].ID).To(Equal("276bbec3-cc77-44b9-a46d-c7760947ec9d"))
			Expect(un[0].Name).To(Equal("user1"))
			Expect(un[1].ID).To(Equal("B061D58B-FDFD-46BF-A766-2D38DE3B1D7B"))
			Expect(un[1].Name).To(Equal("user2"))
		})
	})
}

const (
	SampleSelfResponse = `{"data":{"id":"foobarid","created_at":"2016-08-17T21:07:29.697Z","updated_at":"2017-04-13T20:37:22.943Z","username":"ion","email":"admin@ion.io","chat_handle":null,"sys_admin":true,"teams":{"someteamid":"teamname"},"metadata":{}},"meta":{"copyright":"Copyright 2016 Ion Channel Corporation","authors":["kitplummer","Olio Apps"],"version":"v1"},"links":{"self":"https://janice.ionchannel.testing/v1/users/getSelf","created":"https://janice.ionchannel.testing/v1/users/getSelf"},"timestamps":{"created":"2017-04-18T23:45:05.928+00:00","updated":"2017-04-18T23:45:05.928+00:00"}}`
	SampleCreatedUser  = `{"data":{"id":"463843dd-cb9e-486a-8787-64a7e8523378","created_at":"2018-01-05T17:53:41.430Z","updated_at":"2018-01-05T17:53:41.430Z","username":"tester","email":"testuser@test.com","chat_handle":null,"last_active_at":"2017-04-26T16:37:44.372Z","externally_managed":false,"sys_admin":false,"teams":{}}}`
	SampleUserNames    = `{"data":[{"id":"276bbec3-cc77-44b9-a46d-c7760947ec9d","name":"user1"},{"id":"B061D58B-FDFD-46BF-A766-2D38DE3B1D7B","name":"user2"}],"meta":{"total_count":2,"offset":0,"last_update":"0001-01-01T00:00:00Z"}}`
)
