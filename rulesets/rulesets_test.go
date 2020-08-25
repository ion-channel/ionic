package rulesets

import (
	"fmt"
	"testing"
	"time"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestTeam(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Ruleset Object Validation", func() {

		g.It("should return string in JSON", func() {
			createdAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			updatedAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)

			r := RuleSet{
				ID:          "someid",
				TeamID:      "some_teamID",
				Name:        "somename",
				Description: "somedescription",
				RuleIDs:     nil,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
				Rules:       nil,
				Deprecated:  false,
			}
			Expect(fmt.Sprintf("%v", r)).To(Equal(`{"id":"someid","team_id":"some_teamID","name":"somename","description":"somedescription","rule_ids":null,"created_at":"2018-07-07T13:42:47.651387237Z","updated_at":"2018-07-07T13:42:47.651387237Z","rules":null,"has_deprecated_rules":false}`))

		})
	})
}
