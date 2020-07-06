package digests

import (
	"encoding/json"
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestBuildsystemsDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Buildsystems", func() {
		g.It("should say so when no builds are detected", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{},
			}

			ds, err := buildsystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))
			Expect(string(ds[0].Data)).To(Equal(`{"chars":"none detected"}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})

		g.It("should produce a chars digests when one build is present", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{
					Compilers: []scans.Compiler{
						scans.Compiler{
							Name: "gcc",
						},
					},
				},
			}

			ds, err := buildsystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))
			Expect(ds[0].Title).To(Equal("compiler"))
			Expect(string(ds[0].Data)).To(Equal(`{"chars":"gcc"}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})

		g.It("should produce a count digests when more than one build is present", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{
					Compilers: []scans.Compiler{
						scans.Compiler{
							Name: "gcc",
						},
						scans.Compiler{
							Name: "ruby",
						},
						scans.Compiler{
							Name: "python",
						},
						scans.Compiler{
							Name: "rust",
						},
					},
				},
			}

			ds, err := buildsystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))
			Expect(string(ds[0].Data)).To(Equal(`{"count":4}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})

		g.It("should say so when no images are detected", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{},
			}

			ds, err := buildsystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))
			Expect(string(ds[1].Data)).To(Equal(`{"chars":"none detected"}`))
			Expect(ds[1].Pending).To(BeFalse())
			Expect(ds[1].Errored).To(BeFalse())
		})

		g.It("should produce a chars digests when one image is present", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{
					Dockerfile: scans.Dockerfile{
						Images: []scans.Image{
							scans.Image{
								Name: "foo",
							},
						},
					},
				},
			}

			ds, err := buildsystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))
			Expect(ds[1].Title).To(Equal("container image"))
			Expect(string(ds[1].Data)).To(Equal(`{"chars":"foo"}`))
			Expect(ds[1].Pending).To(BeFalse())
			Expect(ds[1].Errored).To(BeFalse())
		})

		g.It("should produce a count digests when more than one container image is present", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{
					Dockerfile: scans.Dockerfile{
						Images: []scans.Image{
							scans.Image{
								Name: "foo",
							},
							scans.Image{
								Name: "bar",
							},
							scans.Image{
								Name: "baz",
							},
						},
					},
				},
			}

			ds, err := buildsystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))
			Expect(string(ds[1].Data)).To(Equal(`{"count":3}`))
			Expect(ds[1].Pending).To(BeFalse())
			Expect(ds[1].Errored).To(BeFalse())
		})
	})
}

func TestCreateCompilerDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("CompilerDigests", func() {
		g.It("should return none detected when no compiler data", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{},
			}

			d, err := createCompilerDigests(s, e)
			Expect(err).To(BeNil())
			Expect(d).NotTo(BeNil())
			Expect(d.Data).To(Equal(json.RawMessage("{\"chars\":\"none detected\"}")))
		})

		g.It("should return 1 compiler with single compiler", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{
					Compilers: []scans.Compiler{
						scans.Compiler{
							Name: "gcc",
						},
					},
				},
			}

			d, err := createCompilerDigests(s, e)
			Expect(err).To(BeNil())
			Expect(err).To(BeNil())
			Expect(d).NotTo(BeNil())
			Expect(d.Data).To(Equal(json.RawMessage("{\"chars\":\"gcc\"}")))
		})

		g.It("should return compiler count of 4", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{
					Compilers: []scans.Compiler{
						scans.Compiler{
							Name: "gcc",
						},
						scans.Compiler{
							Name: "ruby",
						},
						scans.Compiler{
							Name: "python",
						},
						scans.Compiler{
							Name: "rust",
						},
					},
				},
			}

			d, err := createCompilerDigests(s, e)
			Expect(err).To(BeNil())
			Expect(err).To(BeNil())
			Expect(d).NotTo(BeNil())
			Expect(d.Data).To(Equal(json.RawMessage("{\"count\":4}")))
		})
	})
}

func TestCreateImagesDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Container Image Digests", func() {
		g.It("should return none detected when no container image data", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{},
			}

			d, err := createImagesDigests(s, e)
			Expect(err).To(BeNil())
			Expect(d).NotTo(BeNil())
			Expect(d.Data).To(Equal(json.RawMessage(`{"chars":"none detected"}`)))
		})

		g.It("should return 1 result with single container image", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{
					Dockerfile: scans.Dockerfile{
						Images: []scans.Image{
							scans.Image{
								Name: "foo",
							},
						},
					},
				},
			}

			d, err := createImagesDigests(s, e)
			Expect(err).To(BeNil())
			Expect(err).To(BeNil())
			Expect(d).NotTo(BeNil())
			Expect(d.Data).To(Equal(json.RawMessage(`{"chars":"foo"}`)))
		})

		g.It("should return container image count of 4", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{
					Dockerfile: scans.Dockerfile{
						Images: []scans.Image{
							scans.Image{
								Name: "foo",
							},
							scans.Image{
								Name: "bar",
							},
							scans.Image{
								Name: "baz",
							},
						},
					},
				},
			}

			d, err := createImagesDigests(s, e)
			Expect(err).To(BeNil())
			Expect(err).To(BeNil())
			Expect(d).NotTo(BeNil())
			Expect(d.Data).To(Equal(json.RawMessage(`{"count":3}`)))
		})
	})
}
