package digests

import (
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestDependenciesDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Dependencies", func() {
		g.It("should produce digests", func() {
			s := &scanner.ScanStatus{}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "dependency",
				Data: scans.DependencyResults{
					Dependencies: nil,
					Meta: scans.DependencyMeta{
						FirstDegreeCount:     2,
						NoVersionCount:       1,
						TotalUniqueCount:     115,
						UpdateAvailableCount: 2,
					},
				},
			}

			ds, err := dependencyDigests(e, s)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(4))
			Expect(ds[0].Title).To(Equal("dependencies outdated"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":2}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())

			Expect(ds[1].Title).To(Equal("dependency no version specified"))
			Expect(string(ds[1].Data)).To(Equal(`{"count":1}`))
			Expect(ds[1].Pending).To(BeFalse())
			Expect(ds[1].Errored).To(BeFalse())

			Expect(ds[2].Title).To(Equal("direct dependencies"))
			Expect(string(ds[2].Data)).To(Equal(`{"count":2}`))
			Expect(ds[2].Pending).To(BeFalse())
			Expect(ds[2].Errored).To(BeFalse())

			Expect(ds[3].Title).To(Equal("transitive dependencies"))
			Expect(string(ds[3].Data)).To(Equal(`{"count":113}`))
			Expect(ds[3].Pending).To(BeFalse())
			Expect(ds[3].Errored).To(BeFalse())
		})
	})
}
