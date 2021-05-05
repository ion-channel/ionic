package events

import (
	"encoding/json"
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestUserEvents(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("User Events", func() {
		g.It("should unmarshal a user event action", func() {
			var ue UserEvent
			err := json.Unmarshal([]byte(SampleValidUserEvent), &ue)

			Expect(err).To(BeNil())
			Expect(ue.Action).To(Equal(UserEventAction("forgot_password")))
		})

		g.It("should return an error for an invalid action", func() {
			var ue UserEvent
			err := json.Unmarshal([]byte(SampleInvalidUserEvent), &ue)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("invalid user event action"))
		})

		g.It("should return an invalid error for a blank action", func() {
			var ue UserEvent
			err := json.Unmarshal([]byte(SampleBlankUserEvent), &ue)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("invalid user event action"))
		})
	})
}

const (
	SampleValidUserEvent   = `{"user":{"name":"foo"}, "action":"forgot_password"}`
	SampleInvalidUserEvent = `{"user":{"name":"foo"}, "action":"foo_action"}`
	SampleBlankUserEvent   = `{"user":{"name":"foo"}, "action":""}`
)
