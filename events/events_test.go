package events

import (
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestEvents(t *testing.T) {
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
	})
}

const (
	SampleValidEventJSON = `{"vulnerability":{"updates": ["CVE-2014-2734"]}}`
)
