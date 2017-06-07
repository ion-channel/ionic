package ionic

import (
	"encoding/json"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestEvent(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Event", func() {
		g.Describe("Vulnerability Events", func() {
			g.It("should return true if an event contains the specified vulnerability", func() {
				e := &Event{Vulnerability: &VulnerabilityEvent{Updates: []string{"CVE-2014-2734"}}}
				Expect(e.contains("baz")).To(BeFalse())
				Expect(e.contains("CVE-2014-2734")).To(BeTrue())
			})

			g.It("should return false if an event contains no vulnerabilities", func() {
				e := &Event{Vulnerability: &VulnerabilityEvent{Updates: nil}}
				Expect(e.contains("baz")).To(BeFalse())
				Expect(e.contains("CVE-2014-2734")).To(BeFalse())
			})

			g.It("should append two events", func() {
				left := Event{
					Vulnerability: &VulnerabilityEvent{
						Updates: []string{"bar"},
					},
				}
				right := Event{
					Vulnerability: &VulnerabilityEvent{
						Updates: []string{"foo"},
					},
				}

				left.Append(right)
				Expect(len(left.Vulnerability.Updates)).To(Equal(2))
			})
		})

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
		})
	})
}

const (
	SampleValidEventJSON   = `{"vulnerability":{"updates": ["CVE-2014-2734"]}}`
	SampleValidUserEvent   = `{"user":{"name":"foo"}, "action":"forgot_password"}`
	SampleInvalidUserEvent = `{"user":{"name":"foo"}, "action":"foo_action"}`
)
