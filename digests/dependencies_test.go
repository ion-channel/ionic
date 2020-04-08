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
			s.Status = scanner.ScanStatusFinished
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

			ds, err := dependencyDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(4))
			Expect(ds[0].Title).To(Equal("dependencies outdated"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":2}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())

			Expect(ds[1].Title).To(Equal("dependency no version specified"))
			Expect(string(ds[1].Data)).To(Equal(`{"count":1}`))
			Expect(ds[1].Warning).To(BeTrue())
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

		g.It("should produce no version digest with relevent data", func() {
			s := &scanner.ScanStatus{}
			s.Status = scanner.ScanStatusFinished
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "dependency",
				Data: scans.DependencyResults{
					Dependencies: []scans.Dependency{
						scans.Dependency{
							Name:        "ExpectNoVersion",
							Requirement: "",
						},
						scans.Dependency{
							Name:        "ExpectVersion",
							Requirement: "1.1.1",
						},
					},
					Meta: scans.DependencyMeta{
						FirstDegreeCount:     2,
						NoVersionCount:       1,
						TotalUniqueCount:     115,
						UpdateAvailableCount: 2,
					},
				},
			}

			ds, err := dependencyDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(4))

			Expect(ds[1].Title).To(Equal("dependency no version specified"))
			Expect(string(ds[1].Data)).To(Equal(`{"count":1}`))
			Expect(string(ds[1].SourceData)).To(Equal(`{"type":"dependency","data":[{"latest_version":"","org":"","name":"ExpectNoVersion","type":"","package":"","version":"","scope":"","requirement":""}]}`))
			Expect(ds[1].Warning).To(BeTrue())
			Expect(ds[1].Pending).To(BeFalse())
			Expect(ds[1].Errored).To(BeFalse())
		})

		g.It("should have no warning with transitive dependencies", func() {
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

			ds, err := dependencyDigests(s, e)
			Expect(err).To(BeNil())
			Expect(ds[3].Warning).To(BeFalse())
			Expect(ds[3].WarningMessage).To(BeEmpty())
			Expect(string(ds[3].Data)).To(ContainSubstring("count\":113"))
			Expect(string(ds[3].Title)).To(Equal("transitive dependencies"))
		})

		g.It("should have a warning with transitive dependencies", func() {
			s := &scanner.ScanStatus{}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "dependency",
				Data: scans.DependencyResults{
					Dependencies: nil,
					Meta: scans.DependencyMeta{
						FirstDegreeCount:     2,
						NoVersionCount:       1,
						TotalUniqueCount:     2,
						UpdateAvailableCount: 2,
					},
				},
			}

			ds, err := dependencyDigests(s, e)
			Expect(err).To(BeNil())
			Expect(ds[3].Warning).To(BeTrue())
			Expect(ds[3].WarningMessage).To(Equal("no transitive dependencies found"))
			Expect(string(ds[3].Data)).To(ContainSubstring("count\":0"))
			Expect(string(ds[3].Title)).To(Equal("transitive dependencies"))
		})

		g.It("should have no warning with direct dependencies", func() {
			s := &scanner.ScanStatus{}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "dependency",
				Data: scans.DependencyResults{
					Dependencies: nil,
					Meta: scans.DependencyMeta{
						FirstDegreeCount:     13,
						NoVersionCount:       1,
						TotalUniqueCount:     115,
						UpdateAvailableCount: 2,
					},
				},
			}

			ds, err := dependencyDigests(s, e)
			Expect(err).To(BeNil())
			Expect(ds[2].Warning).To(BeFalse())
			Expect(ds[2].WarningMessage).To(BeEmpty())
			Expect(string(ds[2].Data)).To(ContainSubstring("count\":13"))
			Expect(string(ds[2].Title)).To(Equal("direct dependencies"))
		})

		g.It("should have a warning with direct dependencies", func() {
			s := &scanner.ScanStatus{}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "dependency",
				Data: scans.DependencyResults{
					Dependencies: nil,
					Meta: scans.DependencyMeta{
						FirstDegreeCount:     0,
						NoVersionCount:       1,
						TotalUniqueCount:     2,
						UpdateAvailableCount: 2,
					},
				},
			}

			ds, err := dependencyDigests(s, e)
			Expect(err).To(BeNil())
			Expect(ds[2].Warning).To(BeTrue())
			Expect(ds[2].WarningMessage).To(Equal("no direct dependencies found"))
			Expect(string(ds[2].Data)).To(ContainSubstring("count\":0"))
			Expect(string(ds[2].Title)).To(Equal("direct dependencies"))
		})
	})

}
