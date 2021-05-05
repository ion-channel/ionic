package vulnerabilities

import (
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestVulnerabilities(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Vulnerabilities expansion", func() {
		g.It("should return a new cvss version 3", func() {
			cvssv3 := NewV3FromShorthand("CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:U/C:N/I:N/A:L")

			Expect(cvssv3.AttackVector).To(Equal("network"))
			Expect(cvssv3.AttackComplexity).To(Equal("low"))
			Expect(cvssv3.PrivilegesRequired).To(Equal("none"))
			Expect(cvssv3.UserInteraction).To(Equal("required"))
			Expect(cvssv3.Scope).To(Equal("unchanged"))
			Expect(cvssv3.ConfidentialityImpact).To(Equal("none"))
			Expect(cvssv3.IntegrityImpact).To(Equal("none"))
			Expect(cvssv3.AvailabilityImpact).To(Equal("low"))

			// Additional AccessVector cases
			cvssv3 = NewV3FromShorthand("CVSS:3.0/AV:A/AC:L/PR:N/UI:R/S:U/C:N/I:N/A:L")
			Expect(cvssv3.AttackVector).To(Equal("adjacent"))

			cvssv3 = NewV3FromShorthand("CVSS:3.0/AV:L/AC:L/PR:N/UI:R/S:U/C:N/I:N/A:L")
			Expect(cvssv3.AttackVector).To(Equal("local"))

			cvssv3 = NewV3FromShorthand("CVSS:3.0/AV:P/AC:L/PR:N/UI:R/S:U/C:N/I:N/A:L")
			Expect(cvssv3.AttackVector).To(Equal("physical"))

			// Additional UserInteraction cases
			cvssv3 = NewV3FromShorthand("CVSS:3.0/AV:P/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:L")
			Expect(cvssv3.UserInteraction).To(Equal("none"))

			// Additional Scope cases
			cvssv3 = NewV3FromShorthand("CVSS:3.0/AV:P/AC:L/PR:N/UI:N/S:C/C:N/I:N/A:L")
			Expect(cvssv3.Scope).To(Equal("changed"))
		})

		g.It("should parse low high and none shorthands", func() {
			table := []struct {
				Shorthand string
				Longhand  string
			}{
				{"L", "low"},
				{"H", "high"},
				{"N", "none"},
				{"default", "default"},
			}

			for _, row := range table {
				Expect(parseLowHighNone(row.Shorthand)).To(Equal(row.Longhand))
			}
		})
	})
}
