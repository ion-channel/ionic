package users

import (
	"context"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestTokens(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Users", func() {
		g.Describe("Parsing", func() {
			g.It("should find a user in a context", func() {
				u := &User{
					ID:                "arandouserid",
					Email:             "fantastic@drwho.net",
					ExternallyManaged: false,
					SysAdmin:          false,
					Teams: map[string]string{
						"coolteam":    "user",
						"awesometeam": "admin",
					},
				}

				ctx := context.Background()
				ctx = u.WithContext(ctx)

				newU, err := FromContext(ctx)
				Expect(err).To(BeNil())
				Expect(newU.ID).To(Equal("arandouserid"))
				Expect(newU.Email).To(Equal("fantastic@drwho.net"))
				Expect(len(newU.Teams)).To(Equal(2))
				Expect(newU.Teams["coolteam"]).To(Equal("user"))
			})

			g.It("should return an error if it cannot find the user in the context", func() {
				ctx := context.Background()

				u, err := FromContext(ctx)
				Expect(err).NotTo(BeNil())
				Expect(u).To(BeNil())
			})
		})
	})
}
