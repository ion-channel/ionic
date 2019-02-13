package digests

import (
	"encoding/json"
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestVulnerabilitiesDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Vulnerabilities", func() {
		g.It("should produce digests", func() {
			s := &scanner.ScanStatus{}
			e := scans.NewEval()

			var r scans.VulnerabilityResults
			b := []byte(`{"vulnerabilities":[{"id":910244496,"name":"eslint-scope","org":"eslint","version":"3.7.2","up":"","edition":"","aliases":null,"created_at":"2018-08-15T22:40:27.281Z","updated_at":"2018-08-15T22:41:03.309Z","title":"Eslint Eslint-scope 3.7.2","references":[],"part":"/a","language":"","external_id":"cpe:/a:eslint:eslint-scope:3.7.2","cpe23":null,"target_hw":null,"target_sw":null,"sw_edition":null,"other":null,"vulnerabilities":[{"id":267967698,"external_id":"NPM-673","title":"Malicious package","summary":"Version 3.7.2 of eslint-scope was published without authorization and was found to contain malicious code. This code would read the users .npmrc file and send any found authentication tokens to 2 remote servers.","score":"10.0","score_version":"","score_system":"CVSS","score_details":{"cvssv2":null,"cvssv3":{"vectorString":"","accessVector":"","accessComplexity":"","privilegesRequired":"","userInteraction":"","scope":"","confidentialityImpact":"","integrityImpact":"","availabilityImpact":"","baseScore":10.0,"baseSeverity":""}},"vector":"","access_complexity":"","vulnerability_authentication":null,"confidentiality_impact":"","integrity_impact":"","availability_impact":"","vulnerability_source":null,"assessment_check":null,"scanner":null,"recommendation":"","references":[{"type":"UNKNOWN","source":"NPM","url":"https://www.npmjs.com/advisories/673","text":"https://www.npmjs.com/advisories/673"},{"id":910244496}],"modified_at":"0001-01-01T00:00:00.000Z","published_at":"0001-01-01T00:00:00.000Z","created_at":"2018-08-15T22:56:54.294Z","updated_at":"2018-08-15T22:56:54.294Z","source":[{"id":3,"name":"NPM","description":"NPM Security advisories","created_at":"2018-08-15T22:03:09.285Z","updated_at":"2018-08-15T22:03:09.285Z","attribution":"NPM Term and Licenses https://www.npmjs.com/policies/terms","license":"https://www.npmjs.com/policies/open-source-terms","copyright_url":"https://www.npmjs.com/policies/dmca"}]}]}, {"id":910244496,"name":"eslint-scope","org":"eslint","version":"3.7.2","up":"","edition":"","aliases":null,"created_at":"2018-08-15T22:40:27.281Z","updated_at":"2018-08-15T22:41:03.309Z","title":"Eslint Eslint-scope 3.7.2","references":[],"part":"/a","language":"","external_id":"cpe:/a:eslint:eslint-scope:3.7.2","cpe23":null,"target_hw":null,"target_sw":null,"sw_edition":null,"other":null,"vulnerabilities":[{"id":267967698,"external_id":"NPM-673","title":"Malicious package","summary":"Version 3.7.2 of eslint-scope was published without authorization and was found to contain malicious code. This code would read the users .npmrc file and send any found authentication tokens to 2 remote servers.","score":"10.0","score_version":"","score_system":"CVSS","score_details":{"cvssv2":null,"cvssv3":{"vectorString":"","accessVector":"","accessComplexity":"","privilegesRequired":"","userInteraction":"","scope":"","confidentialityImpact":"","integrityImpact":"","availabilityImpact":"","baseScore":10.0,"baseSeverity":""}},"vector":"","access_complexity":"","vulnerability_authentication":null,"confidentiality_impact":"","integrity_impact":"","availability_impact":"","vulnerability_source":null,"assessment_check":null,"scanner":null,"recommendation":"","references":[{"type":"UNKNOWN","source":"NPM","url":"https://www.npmjs.com/advisories/673","text":"https://www.npmjs.com/advisories/673"},{"id":910244496}],"modified_at":"0001-01-01T00:00:00.000Z","published_at":"0001-01-01T00:00:00.000Z","created_at":"2018-08-15T22:56:54.294Z","updated_at":"2018-08-15T22:56:54.294Z","source":[{"id":3,"name":"NPM","description":"NPM Security advisories","created_at":"2018-08-15T22:03:09.285Z","updated_at":"2018-08-15T22:03:09.285Z","attribution":"NPM Term and Licenses https://www.npmjs.com/policies/terms","license":"https://www.npmjs.com/policies/open-source-terms","copyright_url":"https://www.npmjs.com/policies/dmca"}]}]}],"meta":{"vulnerability_count":2}}`)
			err := json.Unmarshal(b, &r)
			Expect(err).To(BeNil())

			e.TranslatedResults = &scans.TranslatedResults{
				Type: "vulnerability",
				Data: r,
			}

			ds, err := vulnerabilityDigests(e, s)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))

			Expect(ds[0].Title).To(Equal("total vulnerabilities"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":2}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())

			Expect(ds[1].Title).To(Equal("unique vulnerability"))
			Expect(string(ds[1].Data)).To(Equal(`{"count":1}`))
			Expect(ds[1].Pending).To(BeFalse())
			Expect(ds[1].Errored).To(BeFalse())
		})
	})
}
