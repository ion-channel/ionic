package teams

import (
	"fmt"
	"testing"
	"time"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestTeam(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Team Object Validation", func() {

		g.It("should return string in JSON", func() {
			createdAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			updatedAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			deletedAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)

			t := Team{
				ID:               "someid",
				CreatedAt:        createdAt,
				UpdatedAt:        updatedAt,
				DeletedAt:        deletedAt,
				Name:             "somename",
				Delivering:       false,
				SysAdmin:         true,
				POCName:          "youknowit",
				POCEmail:         "you@know.it",
				DefaultDeployKey: "secrets here",
			}
			Expect(fmt.Sprintf("%v", t)).To(Equal(`{"id":"someid","created_at":"2018-07-07T13:42:47.651387237Z","updated_at":"2018-07-07T13:42:47.651387237Z","deleted_at":"2018-07-07T13:42:47.651387237Z","name":"somename","delivering":false,"sys_admin":true,"poc_name":"youknowit","poc_email":"you@know.it","default_deploy_key":"secrets here"}`))
		})
	})
}
