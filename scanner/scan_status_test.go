package scanner

import (
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestScanStatus(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("ScanStatus", func() {
		g.It("should return true if a scan status errored", func() {
			s := &ScanStatus{
				Status: "ErRoReD",
			}

			Expect(s.Errored()).To(BeTrue())
		})

		g.It("should return false if a scan status has not errored", func() {
			s := &ScanStatus{
				Status: "anythingelse",
			}

			Expect(s.Errored()).To(BeFalse())
		})
	})
}
