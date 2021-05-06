package ionic

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestSessions(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Sessions", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get a session for valid credentials", func() {
			server.AddPath("/v1/sessions/login").
				SetMethods("POST").
				SetPayload([]byte(SampleSessionResponse)).
				SetStatus(http.StatusOK)

			session, err := client.Login("ion", "secretpass")
			Expect(err).To(BeNil())
			Expect(session.BearerToken).To(Equal("supersecretkey"))
			Expect(session.User.Email).To(Equal("ion@ion.ion"))
		})
	})
}

const (
	SampleSessionResponse = `{"data":{"jwt":"supersecretkey","user":{"id":"userid","created_at":"2016-08-17T21:07:29.697Z","updated_at":"2017-04-27T22:11:00.404Z","username":"ion","email":"ion@ion.ion","chat_handle":null,"last_active_at":"2017-04-27T22:11:00.396Z","sys_admin":true,"teams":{"adminteamid":"admin"}}},"meta":{"copyright":"Copyright 2016 Ion Channel Corporation","authors":["kitplummer","Olio Apps"],"version":"v1"},"links":{"self":"https://janice.ionchannel.testing/v1/sessions/login","created":"https://janice.ionchannel.testing/v1/sessions/login"},"timestamps":{"created":"2017-04-27T22:11:10.546+00:00","updated":"2017-04-27T22:11:10.546+00:00"}}`
)
