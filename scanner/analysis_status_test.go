package scanner

import (
	"testing"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestEvaluation(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("AnalysisStatus", func() {

		g.It("should provide a simple function for determining done status", func() {
			a := &AnalysisStatus{
				Status: AnalysisStatusErrored,
			}

			Expect(a.Done()).To(BeTrue())
		})

		g.It("should provide a simple function for determining not done status", func() {
			a := &AnalysisStatus{
				Status: AnalysisStatusAnalyzing,
			}

			Expect(a.Done()).To(BeFalse())
		})
	})
}
