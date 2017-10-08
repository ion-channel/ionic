package scans

import (
	"encoding/json"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestScanResults(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Scan Results", func() {
		g.It("should unmarshal a scan results with about yml data", func() {
			var r Results
			err := json.Unmarshal([]byte(SampleValidScanResultsAboutYML), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("about_yml"))

			a, ok := r.Data.(AboutYMLResults)
			Expect(ok).To(Equal(true))
			Expect(a.Content).To(Equal("some content"))
		})

		g.It("should unmarshal a scan results with coverage data", func() {
			var r Results
			err := json.Unmarshal([]byte(SampleValidScanResultsCoverage), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("external_coverage"))

			c, ok := r.Data.(CoverageResults)
			Expect(ok).To(Equal(true))
			Expect(c.Value).To(Equal(42.0))
		})

		g.It("should unmarshal a scan results with dependency data", func() {
			var r Results
			err := json.Unmarshal([]byte(SampleValidScanResultsDependency), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("dependency"))

			d, ok := r.Data.(DependencyResults)
			Expect(ok).To(Equal(true))
			Expect(len(d.Dependencies)).To(Equal(7))
			Expect(d.Meta.FirstDegreeCount).To(Equal(3))
			Expect(d.Meta.NoVersionCount).To(Equal(0))
			Expect(d.Meta.TotalUniqueCount).To(Equal(7))
			Expect(d.Meta.UpdateAvailableCount).To(Equal(2))
		})

		g.It("should unmarshal a scan results with ecosystem data", func() {
			var r Results
			err := json.Unmarshal([]byte(SampleValidScanResultsEcosystems), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("ecosystems"))

			e, ok := r.Data.(EcosystemResults)
			Expect(ok).To(Equal(true))
			Expect(len(e.Ecosystems)).To(Equal(3))
		})

		g.It("should unmarshal a scan results with license data", func() {
			var r Results
			err := json.Unmarshal([]byte(SampleValidScanResultsLicense), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("license"))

			l, ok := r.Data.(LicenseResults)
			Expect(ok).To(Equal(true))
			Expect(l.License.Name).To(Equal("Not found"))
		})

		g.It("should return an error for an invalid results type", func() {
			var r Results
			err := json.Unmarshal([]byte(SampleInvalidResults), &r)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("invalid results type"))
		})
	})
}

const (
	SampleValidScanResultsAboutYML   = `{"type":"about_yml", "data":{"message": "foo message", "valid": true, "content": "some content"}}`
	SampleValidScanResultsCoverage   = `{"type":"external_coverage", "data":{"value":42.0}}`
	SampleValidScanResultsDependency = `{"type":"dependency","data":{"dependencies":[{"latest_version":"2.0","org":"net.sourceforge.javacsv","name":"javacsv","type":"maven","package":"jar","version":"2.0","scope":"compile"},{"latest_version":"4.12","org":"junit","name":"junit","type":"maven","package":"jar","version":"4.11","scope":"test"},{"latest_version":"1.4-atlassian-1","org":"org.hamcrest","name":"hamcrest-core","type":"maven","package":"jar","version":"1.3","scope":"test"},{"latest_version":"4.5.2","org":"org.apache.httpcomponents","name":"httpclient","type":"maven","package":"jar","version":"4.3.4","scope":"compile"},{"latest_version":"4.4.5","org":"org.apache.httpcomponents","name":"httpcore","type":"maven","package":"jar","version":"4.3.2","scope":"compile"},{"latest_version":"99.0-does-not-exist","org":"commons-logging","name":"commons-logging","type":"maven","package":"jar","version":"1.1.3","scope":"compile"},{"latest_version":"20041127.091804","org":"commons-codec","name":"commons-codec","type":"maven","package":"jar","version":"1.6","scope":"compile"}],"meta":{"first_degree_count":3,"no_version_count":0,"total_unique_count":7,"update_available_count":2}}}`
	SampleValidScanResultsEcosystems = `{"type":"ecosystems","data":{"ecosystems":[{"ecosystem":"Java","lines":2430},{"ecosystem":"Makefile","lines":210},{"ecosystem":"Ruby","lines":666}]}}`
	SampleValidScanResultsLicense    = `{"type":"license","data":{"license":{"name":"Not found","type":[]}}}`
	SampleInvalidResults             = `{"type":"fooresult", "data":"I pitty the foo"}`
)
