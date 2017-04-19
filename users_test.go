package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestUsers(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Users", func() {
		server := bogus.New()
		server.Start()
		h, p := server.HostPort()
		client, _ := New("secret", fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get users for an event", func() {
			server.AddPath("/v1/users/subscribedForEvent").
				SetMethods("POST").
				SetPayload([]byte(SampleUsersForEventResponse)).
				SetStatus(http.StatusOK)
			e := Event{}

			users, err := client.GetUsersSubscribedForEvent(e)
			Expect(err).To(BeNil())
			Expect(len(users)).To(Equal(1))
			Expect(users[0].Email).To(Equal("ion@iontest.io"))
		})

		g.It("should get a users self", func() {
			server.AddPath("/v1/users/getSelf").
				SetMethods("GET").
				SetPayload([]byte(SampleSelfResponse)).
				SetStatus(http.StatusOK)

			me, err := client.GetSelf()
			Expect(err).To(BeNil())
			Expect(me.Email).To(Equal("admin@ion.io"))
			Expect(me.Username).To(Equal("ion"))
			Expect(me.SysAdmin).To(Equal(true))
		})
	})
}

const (
	SampleUsersForEventResponse = `{"data":{"users":[{"email":"ion@iontest.io","username":"ion"}]},"meta":{"copyright":"Copyright 2016 Ion Channel Corporation","authors":["kitplummer","Olio Apps"],"version":"v1"},"links":{"self":"https://janice.ionchannel.testing/v1/users/subscribedForEvent","created":"https://janice.ionchannel.testing/v1/users/subscribedForEvent"},"timestamps":{"created":"2017-04-18T18:56:39.076+00:00","updated":"2017-04-18T18:56:39.076+00:00"}}`
	SampleSelfResponse          = `{"data":{"id":"foobarid","created_at":"2016-08-17T21:07:29.697Z","updated_at":"2017-04-13T20:37:22.943Z","username":"ion","email":"admin@ion.io","chat_handle":null,"sys_admin":true,"teams":{"someteamid":"teamname"},"metadata":{}},"meta":{"copyright":"Copyright 2016 Ion Channel Corporation","authors":["kitplummer","Olio Apps"],"version":"v1"},"links":{"self":"https://janice.ionchannel.testing/v1/users/getSelf","created":"https://janice.ionchannel.testing/v1/users/getSelf"},"timestamps":{"created":"2017-04-18T23:45:05.928+00:00","updated":"2017-04-18T23:45:05.928+00:00"}}`
)
