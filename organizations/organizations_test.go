package organizations

import (
	"fmt"
	"testing"
	"time"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestTeam(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Organization Object Validation", func() {

		g.It("should return string in JSON", func() {
			createdAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			updatedAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			deletedAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)

			org := Organization{
				ID:               "someid",
				CreatedAt:        createdAt,
				UpdatedAt:        updatedAt,
				DeletedAt:        deletedAt,
				Name:             "somename",
			}
			Expect(fmt.Sprintf("%v", org)).To(Equal(`{"id":"someid","created_at":"2018-07-07T13:42:47.651387237Z","updated_at":"2018-07-07T13:42:47.651387237Z","deleted_at":"2018-07-07T13:42:47.651387237Z","name":"somename"}`))
		})
	})
}
