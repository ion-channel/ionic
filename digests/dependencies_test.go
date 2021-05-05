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
							Dependencies: []scans.Dependency{
								scans.Dependency{
									Name:        "ExpectNoVersion",
									Requirement: ">=0.0",
								},
							},
						},
						scans.Dependency{
							Name:        "ExpectVersion",
							Requirement: "1.1.1",
							Dependencies: []scans.Dependency{
								scans.Dependency{
									Name:        "ExpectNoVersion",
									Requirement: "",
								},
								scans.Dependency{
									Name:        "ExpectVersion",
									Requirement: "1.1.1",
									Dependencies: []scans.Dependency{
										scans.Dependency{
											Name:        "ExpectNoVersion",
											Requirement: ">= 0",
										},
									},
								},
							},
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
			Expect(string(ds[1].SourceData)).To(Equal(`{"type":"dependency","data":[{"latest_version":"","org":"","name":"ExpectNoVersion","type":"","package":"","version":"","scope":"","requirement":"","file":"","dependencies":null},{"latest_version":"","org":"","name":"ExpectVersion","type":"","package":"","version":"","scope":"","requirement":"1.1.1","file":"","dependencies":[{"latest_version":"","org":"","name":"ExpectNoVersion","type":"","package":"","version":"","scope":"","requirement":"\u003e=0.0","file":"","dependencies":null}]},{"latest_version":"","org":"","name":"ExpectVersion","type":"","package":"","version":"","scope":"","requirement":"1.1.1","file":"","dependencies":[{"latest_version":"","org":"","name":"ExpectNoVersion","type":"","package":"","version":"","scope":"","requirement":"","file":"","dependencies":null},{"latest_version":"","org":"","name":"ExpectVersion","type":"","package":"","version":"","scope":"","requirement":"1.1.1","file":"","dependencies":[{"latest_version":"","org":"","name":"ExpectNoVersion","type":"","package":"","version":"","scope":"","requirement":"\u003e= 0","file":"","dependencies":null}]}]}]}`))
			Expect(ds[1].Warning).To(BeTrue())
			Expect(ds[1].Pending).To(BeFalse())
			Expect(ds[1].Errored).To(BeFalse())
		})

		g.It("should produce outdated digest with relevant data and counts", func() {
			s := &scanner.ScanStatus{}
			s.Status = scanner.ScanStatusFinished
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "dependency",
				Data: scans.DependencyResults{
					Dependencies: []scans.Dependency{
						scans.Dependency{
							Name:          "UpToDate",
							Version:       "3.0.0",
							LatestVersion: "3.0.0",
							Requirement:   "",
						},
						scans.Dependency{
							Name:          "ExpectVersion1",
							Version:       "5.9.1",
							LatestVersion: "5.14.2",
							Requirement:   "5.1",
						},
						scans.Dependency{
							Name:          "ExpectVersion2",
							Version:       "0.9.9.6",
							LatestVersion: "0.13.1",
							Requirement:   "0",
						},
						scans.Dependency{
							Name:          "ExpectVersion3",
							Version:       "10",
							LatestVersion: "10",
							Requirement:   "10",
							Dependencies: []scans.Dependency{
								scans.Dependency{
									Name:          "ExpectVersionTDep1",
									Version:       "1.1.1",
									LatestVersion: "2.0.0",
									Requirement:   "1.1.1",
								},
								scans.Dependency{
									Name:          "ExpectVersionTDep2",
									Version:       "1.0.0",
									LatestVersion: "1.5.0",
									Requirement:   "1.0.0",
								},
							},
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

			Expect(ds[0].Title).To(Equal("dependencies outdated"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":2}`))

			// ds[0] is dependencies outdated
			ds0Source := `{"type":"dependency","data":[{"latest_version":"5.14.2","org":"","name":"ExpectVersion1","type":"","package":"","version":"5.9.1","scope":"","requirement":"5.1","file":"","dependency_counts":{"first_degree_count":0,"no_version_count":0,"total_unique_count":0,"update_available_count":0,"vulnerable_count":0},"outdated_version":{"major_behind":0,"minor_behind":5,"patch_behind":1},"dependencies":[]},{"latest_version":"0.13.1","org":"","name":"ExpectVersion2","type":"","package":"","version":"0.9.9.6","scope":"","requirement":"0","file":"","dependency_counts":{"first_degree_count":0,"no_version_count":0,"total_unique_count":0,"update_available_count":0,"vulnerable_count":0},"outdated_version":{"major_behind":0,"minor_behind":4,"patch_behind":0},"dependencies":[]},{"latest_version":"10","org":"","name":"ExpectVersion3","type":"","package":"","version":"10","scope":"","requirement":"10","file":"","dependency_counts":{"first_degree_count":0,"no_version_count":0,"total_unique_count":0,"update_available_count":2,"vulnerable_count":0},"outdated_version":{"major_behind":0,"minor_behind":0,"patch_behind":0},"dependencies":[{"latest_version":"2.0.0","org":"","name":"ExpectVersionTDep1","type":"","package":"","version":"1.1.1","scope":"","requirement":"1.1.1","file":"","outdated_version":{"major_behind":1,"minor_behind":0,"patch_behind":0},"dependencies":null},{"latest_version":"1.5.0","org":"","name":"ExpectVersionTDep2","type":"","package":"","version":"1.0.0","scope":"","requirement":"1.0.0","file":"","outdated_version":{"major_behind":0,"minor_behind":5,"patch_behind":0},"dependencies":null}]}]}`
			Expect(string(ds[0].SourceData)).To(Equal(ds0Source), "\nds0 source is %+v\n\nexpected is   %+v\n", string(ds[0].SourceData), ds0Source)

			// transitive deps
			ds1Source := `{"type":"dependency","data":[{"latest_version":"3.0.0","org":"","name":"UpToDate","type":"","package":"","version":"3.0.0","scope":"","requirement":"","file":"","dependencies":null}]}`
			Expect(string(ds[1].SourceData)).To(Equal(ds1Source), "\n------\nds1 source is %+v\n\nexpected is   %+v\n", string(ds[1].SourceData), ds1Source)

			// direct deps
			ds2Source := `{"type":"dependency","data":[{"latest_version":"3.0.0","org":"","name":"UpToDate","type":"","package":"","version":"3.0.0","scope":"","requirement":"","file":"","dependencies":null},{"latest_version":"5.14.2","org":"","name":"ExpectVersion1","type":"","package":"","version":"5.9.1","scope":"","requirement":"5.1","file":"","dependencies":null},{"latest_version":"0.13.1","org":"","name":"ExpectVersion2","type":"","package":"","version":"0.9.9.6","scope":"","requirement":"0","file":"","dependencies":null},{"latest_version":"10","org":"","name":"ExpectVersion3","type":"","package":"","version":"10","scope":"","requirement":"10","file":"","dependencies":null}]}`
			Expect(string(ds[2].SourceData)).To(Equal(ds2Source), "\n------\nds2 source is %+v\n\nexpected is   %+v\n", string(ds[2].SourceData), ds2Source)

			Expect(ds[0].Warning).To(BeFalse())
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})

		g.It("should produce transitive digest with relevent data", func() {
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
							Dependencies: []scans.Dependency{
								scans.Dependency{
									Name:        "ExpectVersion",
									Requirement: "1.1.1",
								},
							},
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

			Expect(ds[3].Title).To(Equal("transitive dependencies"))
			Expect(string(ds[3].Data)).To(Equal(`{"count":113}`))
			Expect(string(ds[3].SourceData)).To(Equal(`{"type":"dependency","data":[{"latest_version":"","org":"","name":"ExpectNoVersion","type":"","package":"","version":"","scope":"","requirement":"","file":"","dependencies":[{"latest_version":"","org":"","name":"ExpectVersion","type":"","package":"","version":"","scope":"","requirement":"1.1.1","file":"","dependencies":null}]}]}`))
			Expect(ds[3].Warning).To(BeFalse())
			Expect(ds[3].Pending).To(BeFalse())
			Expect(ds[3].Errored).To(BeFalse())
		})

		g.It("should produce direct digest with relevent data", func() {
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
							Dependencies: []scans.Dependency{
								scans.Dependency{
									Name:        "ExpectVersion",
									Requirement: "1.1.1",
								},
							},
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

			Expect(ds[2].Title).To(Equal("direct dependencies"))
			Expect(string(ds[2].Data)).To(Equal(`{"count":2}`))
			b := `{"type":"dependency","data":[{"latest_version":"","org":"","name":"ExpectNoVersion","type":"","package":"","version":"","scope":"","requirement":"","file":"","dependencies":null}]}`
			Expect(string(ds[2].SourceData)).To(Equal(b))
			Expect(ds[2].Warning).To(BeFalse())
			Expect(ds[2].Pending).To(BeFalse())
			Expect(ds[2].Errored).To(BeFalse())
		})

		g.It("should properly calculate versions behind", func() {
			d := scans.Dependency{
				Name:          "ExpectVersion",
				Version:       "1.1.1",
				LatestVersion: "2.0.0",
				Requirement:   "1.1.1",
			}

			// process out of date
			dep := od(&d)
			Expect(dep).ToNot(BeNil())
			Expect(dep.OutdatedMeta.MajorBehind).To(Equal(1))
			Expect(dep.OutdatedMeta.MinorBehind).To(Equal(0))
			Expect(dep.OutdatedMeta.PatchBehind).To(Equal(0))
		})

		g.It("should return a dependency unmodified if the version is bad", func() {
			d := scans.Dependency{
				Name:          "ExpectBadVersion",
				Version:       "1.1.1",
				LatestVersion: "foo1bar",
				Requirement:   "1.1.1",
			}

			// process out of date
			dep := od(&d)
			Expect(dep).ToNot(BeNil())
			Expect(dep.OutdatedMeta).To(BeNil())
		})

		g.It("should process a dependency with odd version lengths", func() {
			d := scans.Dependency{
				Name:          "ExpectGoodVersion",
				Version:       "1.1.1.1.1",
				LatestVersion: "2.1.1.1",
				Requirement:   "2.0",
			}

			// process out of date
			dep := od(&d)
			Expect(dep).ToNot(BeNil())
			Expect(dep.OutdatedMeta).ToNot(BeNil())
		})

		g.It("should return a dependency versions behind for a proper npm package", func() {
			d := scans.Dependency{
				Org:           "eslint",
				Type:          "npmjs",
				Name:          "eslint",
				Package:       "package",
				Version:       "5.16.0",
				LatestVersion: "7.12.1",
				Requirement:   "^5.8.0",
			}

			// process out of date
			dep := od(&d)
			Expect(dep).ToNot(BeNil())
			Expect(dep.OutdatedMeta).ToNot(BeNil())
			Expect(dep.OutdatedMeta.MajorBehind).To(Equal(2))
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
