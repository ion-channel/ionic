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
	})
}

const (
	SampleUsersForEventResponse = `{"data":{"users":[{"email":"ion@iontest.io","username":"ion"}]},"meta":{"copyright":"Copyright 2016 Ion Channel Corporation","authors":["kitplummer","Olio Apps"],"version":"v1"},"links":{"self":"https://janice.ionchannel.testing/v1/users/subscribedForEvent","created":"https://janice.ionchannel.testing/v1/users/subscribedForEvent"},"timestamps":{"created":"2017-04-18T18:56:39.076+00:00","updated":"2017-04-18T18:56:39.076+00:00"}}`
)
