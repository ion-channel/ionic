package digests

import (
	"fmt"
	"sort"
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestDigest(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Digest", func() {
		g.Describe("Sorting", func() {
			g.It("should be sortable by index", func() {
				ds := []Digest{
					Digest{Index: 1},
					Digest{Index: 3},
					Digest{Index: 2},
					Digest{Index: 0},
				}

				Expect(ds[0].Index).To(Equal(1))
				Expect(ds[1].Index).To(Equal(3))
				Expect(ds[2].Index).To(Equal(2))
				Expect(ds[3].Index).To(Equal(0))

				sort.Slice(ds, func(i, j int) bool { return ds[i].Index < ds[j].Index })

				Expect(ds[0].Index).To(Equal(0))
				Expect(ds[1].Index).To(Equal(1))
				Expect(ds[2].Index).To(Equal(2))
				Expect(ds[3].Index).To(Equal(3))
			})
		})

		g.Describe("Strings", func() {
			g.It("should return string in JSON", func() {
				d := Digest{
					Index:          12,
					Title:          "sometitle",
					Data:           nil,
					ScanID:         "somescanid",
					RuleID:         "someruleid",
					RulesetID:      "somerulesetid",
					Evaluated:      true,
					Pending:        false,
					Passed:         true,
					PassedMessage:  "somepassedmessage",
					Warning:        false,
					WarningMessage: "somewarningmessage",
					Errored:        false,
					ErroredMessage: "someerroredmessage",
					singularTitle:  "somesingulartitle",
					pluralTitle:    "somepluraltitle",
				}
				Expect(fmt.Sprintf("%v", d)).To(Equal(`{"index":12,"title":"sometitle","data":null,"source_data":null,"scan_id":"somescanid","rule_id":"someruleid","ruleset_id":"somerulesetid","evaluated":true,"pending":false,"passed":true,"passed_message":"somepassedmessage","warning":false,"warning_message":"somewarningmessage","errored":false,"errored_message":"someerroredmessage"}`))
			})
		})

		g.Describe("States", func() {
			g.It("should be marked as pending when no status is provided", func() {
				ds := NewDigest(nil, 0, "", "")
				Expect(ds).NotTo(BeNil())
				Expect(ds.Pending).To(BeTrue())
			})

			g.It("should show an error if present", func() {
				s := &scanner.ScanStatus{
					Status:  scanner.AnalysisStatusErrored,
					Message: "failed to perform the scan for a reason",
				}
				ds := NewDigest(s, 0, "", "")
				Expect(ds).NotTo(BeNil())
				Expect(ds.Pending).To(BeFalse())
				Expect(ds.Errored).To(BeTrue())
				Expect(ds.ErroredMessage).To(Equal("failed to perform the scan for a reason"))

				s = &scanner.ScanStatus{
					Status:  scanner.AnalysisStatusFinished,
					Message: "completed scan",
				}
				ds = NewDigest(s, 0, "", "")
				Expect(ds).NotTo(BeNil())
				Expect(ds.Pending).To(BeFalse())
				Expect(ds.Errored).To(BeTrue())
				Expect(ds.ErroredMessage).To(Equal("evaluation not received"))
			})

			g.It("should show an error when an invalid count is appended", func() {
				s := &scanner.ScanStatus{
					Status:  "errored",
					Message: "failed to perform the scan for a reason",
				}
				ds := NewDigest(s, 0, "", "")
				e := scans.NewEval()

				ds.AppendEval(e, "count", -1)
				Expect(ds).NotTo(BeNil())
				Expect(ds.Pending).To(BeFalse())
				Expect(ds.Errored).To(BeTrue())
				Expect(ds.ErroredMessage).To(Equal("failed to perform the scan for a reason"))
			})

			g.It("should not have an error", func() {
				s := &scanner.ScanStatus{
					Status:  "finished",
					Message: "completed scan",
				}
				ds := NewDigest(s, 0, "", "")
				e := scans.NewEval()

				ds.AppendEval(e, "count", 1)
				Expect(ds).NotTo(BeNil())
				Expect(ds.Pending).To(BeFalse())
				Expect(ds.Errored).To(BeFalse())
				Expect(ds.ErroredMessage).To(Equal(""))
			})
		})

		g.Describe("Pluralization", func() {
			g.It("should return a singular title when appropriate", func() {
				ds := Digest{
					singularTitle: "vulnerability",
					pluralTitle:   "vulnerabilities",
				}
				e := scans.NewEval()

				err := ds.AppendEval(e, "count", 1)
				Expect(err).To(BeNil())
				Expect(ds.Title).To(Equal("vulnerability"))

				err = ds.AppendEval(e, "list", []string{"something"})
				Expect(err).To(BeNil())
				Expect(ds.Title).To(Equal("vulnerability"))

				err = ds.AppendEval(e, "chars", "true")
				Expect(err).To(BeNil())
				Expect(ds.Title).To(Equal("vulnerabilities"))

				err = ds.AppendEval(e, "bool", true)
				Expect(err).To(BeNil())
				Expect(ds.Title).To(Equal("vulnerabilities"))

				err = ds.AppendEval(e, "percent", float64(10))
				Expect(err).To(BeNil())
				Expect(ds.Title).To(Equal("vulnerabilities"))
			})

			g.It("should return a plural title when appropriate", func() {
				ds := Digest{
					singularTitle: "vulnerability",
					pluralTitle:   "vulnerabilities",
				}
				e := scans.NewEval()

				err := ds.AppendEval(e, "count", 10)
				Expect(err).To(BeNil())
				Expect(ds.Title).To(Equal("vulnerabilities"))

				err = ds.AppendEval(e, "list", []string{"something", "another"})
				Expect(err).To(BeNil())
				Expect(ds.Title).To(Equal("vulnerabilities"))

				err = ds.AppendEval(e, "chars", "true")
				Expect(err).To(BeNil())
				Expect(ds.Title).To(Equal("vulnerabilities"))

				err = ds.AppendEval(e, "bool", true)
				Expect(err).To(BeNil())
				Expect(ds.Title).To(Equal("vulnerabilities"))

				err = ds.AppendEval(e, "percent", float64(10))
				Expect(err).To(BeNil())
				Expect(ds.Title).To(Equal("vulnerabilities"))
			})
		})

		g.Describe("Evaluation", func() {
			g.It("should return an error for an unsupported digest type", func() {
				ds := Digest{}
				e := scans.NewEval()

				err := ds.AppendEval(e, "badtype", 10)
				Expect(err).To(Equal(ErrUnsupportedType))
			})

			g.It("should return an error when the value doesn't match the type", func() {
				ds := Digest{}
				e := scans.NewEval()

				err := ds.AppendEval(e, "bool", "not a bool")
				Expect(err).To(Equal(ErrFailedValueAssertion))

				err = ds.AppendEval(e, "count", "not a count")
				Expect(err).To(Equal(ErrFailedValueAssertion))

				err = ds.AppendEval(e, "list", "not a list")
				Expect(err).To(Equal(ErrFailedValueAssertion))

				err = ds.AppendEval(e, "percent", "not a percent")
				Expect(err).To(Equal(ErrFailedValueAssertion))

				err = ds.AppendEval(e, "chars", true)
				Expect(err).To(Equal(ErrFailedValueAssertion))
			})

			g.It("should append added evaluation info", func() {
				ds := &Digest{}
				e := scans.NewEval()

				e.ID = "someevalid"
				e.RuleID = "someruleid"
				e.RulesetID = "somerulesetid"
				e.Type = "someeval"
				e.Passed = true
				e.Description = "we evaluated a thing"

				err := ds.AppendEval(e, "count", 10)
				Expect(err).To(BeNil())
				Expect(ds.ScanID).To(Equal("someevalid"))
				Expect(ds.RuleID).To(Equal("someruleid"))
				Expect(ds.RulesetID).To(Equal("somerulesetid"))
				Expect(ds.Evaluated).To(BeTrue())
				Expect(ds.Passed).To(BeTrue())
			})
		})
	})
}
