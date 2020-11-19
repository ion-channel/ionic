package events

import (
	"encoding/json"
	"testing"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestProjectEvents(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Project Events", func() {
		g.It("should unmarshal a project event action", func() {
			var ue ProjectEvent
			err := json.Unmarshal([]byte(SampleValidProjectEvent), &ue)

			Expect(err).To(BeNil())
			Expect(ue.Action).To(Equal(ProjectEventAction("project_added")))
		})

		g.It("should return an error for an invalid action", func() {
			var ue ProjectEvent
			err := json.Unmarshal([]byte(SampleInvalidProjectEvent), &ue)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("invalid project event action"))
		})
	})
}

const (
	SampleValidProjectEvent   = `{"project":"apache/spark", "action":"project_added"}`
	SampleInvalidProjectEvent = `{"project":"apache/spark", "action":"foo_action"}`
)
