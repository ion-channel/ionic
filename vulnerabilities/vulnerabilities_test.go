package vulnerabilities

import (
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestVulnerabilities(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Vulnerabilities expansion", func() {
		g.It("should return a new cvss version 3", func() {
      cvssv3 := NewFromShorthandV3("CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:U/C:N/I:N/A:L")

      Expect(cvssv3.AccessVector).To(Equal("network"))
      Expect(cvssv3.AccessComplexity).To(Equal("low"))
      Expect(cvssv3.PrivilegesRequired).To(Equal("none"))
      Expect(cvssv3.UserInteraction).To(Equal("required"))
      Expect(cvssv3.Scope).To(Equal("unchanged"))
      Expect(cvssv3.ConfidentialityImpact).To(Equal("none"))
      Expect(cvssv3.IntegrityImpact).To(Equal("none"))
      Expect(cvssv3.AvailabilityImpact).To(Equal("low"))
		})
	})
}
