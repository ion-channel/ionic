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
		g.It("should return true if an event contains the specified vulnerability", func() {
			e := &Event{EventVulnerability{Updates: []string{"CVE-2014-2734"}}}
			Expect(e.contains("baz")).To(BeFalse())
			Expect(e.contains("CVE-2014-2734")).To(BeTrue())
		})

		g.It("should append two events", func() {
			var left Event
			json.Unmarshal([]byte(SampleValidEventJSON), &left)
			Expect(len(left.Vulns.Updates)).To(Equal(1))

			right := left
			right.Vulns.Updates = []string{"foo"}
			left.Append(right)

			Expect(len(left.Vulns.Updates)).To(Equal(2))
		})
	})
}

const (
	SampleValidEventJSON = `{"vulns":{"updates": ["CVE-2014-2734"]}}`
)
