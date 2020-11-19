package users

import (
	"fmt"
	"testing"
	"time"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestUser(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("User Object Validation", func() {

		g.It("should return string in JSON", func() {
			createdAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			updatedAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			lastActive := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)

			u := User{
				ID:                "someid",
				Email:             "some_email",
				Username:          "some_user",
				ChatHandle:        "some_chat_handle",
				CreatedAt:         createdAt,
				UpdatedAt:         updatedAt,
				LastActive:        lastActive,
				ExternallyManaged: true,
				Metadata:          nil,
				SysAdmin:          true,
				System:            false,
				Teams:             nil,
			}

			Expect(fmt.Sprintf("%v", u)).To(Equal(`{"id":"someid","email":"some_email","username":"some_user","chat_handle":"some_chat_handle","created_at":"2018-07-07T13:42:47.651387237Z","updated_at":"2018-07-07T13:42:47.651387237Z","last_active_at":"2018-07-07T13:42:47.651387237Z","externally_managed":true,"metadata":null,"sys_admin":true,"system":false,"teams":null}`))

		})

		g.It("should check if in team", func() {
			createdAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			updatedAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			lastActive := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)

			u := User{
				ID:                "someid",
				Email:             "some_email",
				Username:          "some_user",
				ChatHandle:        "some_chat_handle",
				CreatedAt:         createdAt,
				UpdatedAt:         updatedAt,
				LastActive:        lastActive,
				ExternallyManaged: true,
				Metadata:          nil,
				SysAdmin:          true,
				System:            false,
				Teams:             make(map[string]string),
			}

			u.Teams["inteam"] = "member"
			Expect(u.IsMemberOfTeam("inteam")).To(Equal(true))
			Expect(u.IsMemberOfTeam("notinteam")).To(Equal(false))
		})

		g.It("should check if an admin in team", func() {
			createdAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			updatedAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			lastActive := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)

			u := User{
				ID:                "someid",
				Email:             "some_email",
				Username:          "some_user",
				ChatHandle:        "some_chat_handle",
				CreatedAt:         createdAt,
				UpdatedAt:         updatedAt,
				LastActive:        lastActive,
				ExternallyManaged: true,
				Metadata:          nil,
				SysAdmin:          true,
				System:            false,
				Teams:             make(map[string]string),
			}

			u.Teams["inteam"] = "admin"
			Expect(u.IsMemberOfTeam("inteam")).To(Equal(true))
			Expect(u.IsAdminOfTeam("inteam")).To(Equal(true))
			Expect(u.IsAdminOfTeam("notinteam")).To(Equal(false))
		})

	})
}
