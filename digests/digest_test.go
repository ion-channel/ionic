package digests

import (
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
		g.Describe("Constructor", func() {
			g.It("should return an error for an unsupported digest type", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				ds, err := NewDigest(0, "A title", "badtype", 10, e, s)
				Expect(err).To(Equal(ErrUnsupportedType))
				Expect(ds).To(BeNil())
			})

			g.It("should return an error when the value doesn't match the type", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				ds, err := NewDigest(0, "Another title", "bool", "not a bool", e, s)
				Expect(err).To(Equal(ErrFailedValueAssertion))
				Expect(ds).To(BeNil())

				s = &scanner.ScanStatus{}
				e = scans.NewEval()
				ds, err = NewDigest(1, "Another title", "count", "not a count", e, s)
				Expect(err).To(Equal(ErrFailedValueAssertion))
				Expect(ds).To(BeNil())

				s = &scanner.ScanStatus{}
				e = scans.NewEval()
				ds, err = NewDigest(2, "Another title", "list", "not a list", e, s)
				Expect(err).To(Equal(ErrFailedValueAssertion))
				Expect(ds).To(BeNil())

				s = &scanner.ScanStatus{}
				e = scans.NewEval()
				ds, err = NewDigest(3, "Another title", "percent", "not a percent", e, s)
				Expect(err).To(Equal(ErrFailedValueAssertion))
				Expect(ds).To(BeNil())

				s = &scanner.ScanStatus{}
				e = scans.NewEval()
				ds, err = NewDigest(3, "Another title", "chars", true, e, s)
				Expect(err).To(Equal(ErrFailedValueAssertion))
				Expect(ds).To(BeNil())
			})
		})

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

		g.Describe("States", func() {
			g.It("should be marked as pending when no status is provided", func() {
				e := scans.NewEval()
				ds, err := NewDigest(0, "Another title", "chars", "some thing", e, nil)
				Expect(err).To(BeNil())
				Expect(ds).NotTo(BeNil())
				Expect(ds.Pending).To(BeTrue())
			})

			g.It("should show an error if present", func() {
				s := &scanner.ScanStatus{
					Status:  "errored",
					Message: "failed to perform the scan for a reason",
				}
				e := scans.NewEval()
				ds, err := NewDigest(0, "Another title", "chars", "some thing", e, s)
				Expect(err).To(BeNil())
				Expect(ds).NotTo(BeNil())
				Expect(ds.Pending).To(BeFalse())
				Expect(ds.Errored).To(BeTrue())
				Expect(ds.ErroredMessage).To(Equal("failed to perform the scan for a reason"))

				s = &scanner.ScanStatus{
					Status:  "finished",
					Message: "completed scan",
				}
				e = scans.NewEval()
				ds, err = NewDigest(0, "Another title", "chars", "some thing", e, s)
				Expect(err).To(BeNil())
				Expect(ds).NotTo(BeNil())
				Expect(ds.Pending).To(BeFalse())
				Expect(ds.Errored).To(BeFalse())
				Expect(ds.ErroredMessage).To(Equal(""))
			})
		})
	})
}
