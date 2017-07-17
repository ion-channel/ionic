package events

import (
	"encoding/json"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestAnalysisEvents(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Analysis Events", func() {
		g.It("should unmarshal a analysis event action", func() {
			var ue AnalysisEvent
			err := json.Unmarshal([]byte(SampleValidAnalysisEvent), &ue)

			Expect(err).To(BeNil())
			Expect(ue.Action).To(Equal(AnalysisEventAction("analysis_failed")))
		})

		g.It("should return an error for an invalid action", func() {
			var ue AnalysisEvent
			err := json.Unmarshal([]byte(SampleInvalidAnalysisEvent), &ue)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("invalid analysis event action"))
		})
	})
}

const (
	SampleValidAnalysisEvent   = `{"analysis":"fooanalysis", "action":"analysis_failed"}`
	SampleInvalidAnalysisEvent = `{"analysis":"fooanalysis", "action":"foo_action"}`
)
