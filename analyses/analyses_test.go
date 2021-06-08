package analyses

import (
	"fmt"
	"testing"
	"time"

	"github.com/ion-channel/ionic/rulesets"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestAnalysis(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("New Analysis Summary", func() {
		g.It("should return a new analysis summary", func() {
			expectedAnalysisID := "analysizing"

			a := &Analysis{
				Summary: Summary{ID: expectedAnalysisID},
			}
			ar := &rulesets.AppliedRulesetSummary{}

			s := NewSummary(a, ar)
			Expect(s.ID).To(Equal(expectedAnalysisID))
			Expect(s.AnalysisID).To(Equal(expectedAnalysisID))
			Expect(s.RulesetName).To(Equal("N/A"))
			Expect(s.Passed).To(Equal(false))
			Expect(s.Risk).To(Equal("high"))
		})

		g.It("should return N/A for ruleset name when it is not available", func() {
			a := &Analysis{}
			ar := &rulesets.AppliedRulesetSummary{
				RuleEvaluationSummary: &rulesets.RuleEvaluationSummary{},
			}

			s := NewSummary(a, ar)
			Expect(s.RulesetName).To(Equal("N/A"))
			Expect(s.Passed).To(Equal(false))
			Expect(s.Risk).To(Equal("high"))
		})

		g.It("should return string in JSON", func() {
			createdAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			updatedAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)

			text := "sometext"

			a := Analysis{
				Summary: Summary{
					ID:            "someid",
					TeamID:        "someteamid",
					ProjectID:     "someproject",
					Name:          "somename",
					Text:          &text,
					Type:          "sometype",
					Source:        "somesource",
					Branch:        "somebranch",
					Description:   "somedesc",
					Status:        "somestatus",
					RulesetID:     "somerulesetid",
					CreatedAt:     createdAt,
					UpdatedAt:     updatedAt,
					Duration:      4.242,
					TriggerHash:   "sometriggerhas",
					TriggerText:   "sometriggertext",
					TriggerAuthor: "sometriggerauthor",
				},
				ScanSummaries: nil,
				Public:        true,
			}
			Expect(fmt.Sprintf("%v", a)).To(Equal(`{"id":"someid","analysis_id":"","team_id":"someteamid","project_id":"someproject","name":"somename","text":"sometext","type":"sometype","source":"somesource","branch":"somebranch","description":"somedesc","risk":"","summary":"","passed":false,"ruleset_id":"somerulesetid","ruleset_name":"","status":"somestatus","created_at":"2018-07-07T13:42:47.651387237Z","updated_at":"2018-07-07T13:42:47.651387237Z","duration":4.242,"trigger_hash":"sometriggerhas","trigger_text":"sometriggertext","trigger_author":"sometriggerauthor","trigger":"","scan_summaries":null,"public":true}`))
		})
	})
}
