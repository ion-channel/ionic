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
			b := []byte(validHighCriticalVulns)
			err := json.Unmarshal(b, &r)
			Expect(err).To(BeNil())

			e.TranslatedResults = &scans.TranslatedResults{
				Type: "vulnerability",
				Data: r,
			}

			ds, err := vulnerabilityDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))

			Expect(ds[0].Title).To(Equal("total vulnerabilities"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":2}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())

			Expect(ds[1].Title).To(Equal("unique vulnerability"))
			Expect(string(ds[1].Data)).To(Equal(`{"count":2}`))
			Expect(ds[1].Pending).To(BeFalse())
			Expect(ds[1].Errored).To(BeFalse())
		})
	})
}

const (
	validHighCriticalVulns = `{"vulnerabilities":[{"id":92400,"external_id":"cpe:/a:rack_project:rack:1.1.0","source_id":0,"title":"Rack_project Rack 1.1.0","name":"rack","org":"rack_project","version":"1.1.0","up":"","edition":"","aliases":null,"created_at":"2017-02-13T20:02:43.83Z","updated_at":"2018-05-25T04:10:15.382Z","references":[],"part":"/a","language":"","vulnerabilities":[{"id":267868924,"external_id":"CVE-2013-0184","source":[{"id":1,"name":"NVD","description":"National Vulnerability Database","created_at":"2017-02-09T20:18:35.385Z","updated_at":"2017-02-13T20:12:05.342Z","attribution":"Copyright © 1999–2017, The MITRE Corporation. CVE and the CVE logo are registered trademarks and CVE-Compatible is a trademark of The MITRE Corporation.","license":"Submissions: For all materials you submit to the Common Vulnerabilities and Exposures (CVE®), you hereby grant to The MITRE Corporation (MITRE) and all CVE Numbering Authorities (CNAs) a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to reproduce, prepare derivative works of, publicly display, publicly perform, sublicense, and distribute such materials and derivative works. Unless required by applicable law or agreed to in writing, you provide such materials on an \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied, including, without limitation, any warranties or conditions of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE.\n\nCVE Usage: MITRE hereby grants you a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to reproduce, prepare derivative works of, publicly display, publicly perform, sublicense, and distribute Common Vulnerabilities and Exposures (CVE®). Any copy you make for such purposes is authorized provided that you reproduce MITRE's copyright designation and this license in any such copy.\n","copyright_url":"http://cve.mitre.org/about/termsofuse.html"}],"title":"CVE-2013-0184","summary":"Unspecified vulnerability in Rack::Auth::AbstractRequest in Rack 1.1.x before 1.1.5, 1.2.x before 1.2.7, 1.3.x before 1.3.9, and 1.4.x before 1.4.4 allows remote attackers to cause a denial of service via unknown vectors related to \"symbolized arbitrary strings.\"","score":"9.5","score_version":"3.0","score_system":"CVSS","score_details":{"cvssv2":{"vectorString":"AV:N/AC:M/Au:N/C:N/I:N/A:P","accessVector":"NETWORK","accessComplexity":"MEDIUM","authentication":"NONE","confidentialityImpact":"NONE","integrityImpact":"NONE","availabilityImpact":"PARTIAL","baseScore":4.3}},"vector":"","access_complexity":"","vulnerability_authentication":"","confidentiality_impact":"","integrity_impact":"","availability_impact":"","vulnerabilty_source":"","assessment_check":null,"scanner":null,"recommendation":"","dependencies":null,"references":[{"type":"UNKNOWN","source":"openSUSE-SU-2013:0462","url":"http://lists.opensuse.org/opensuse-updates/2013-03/msg00048.html","text":"http://lists.opensuse.org/opensuse-updates/2013-03/msg00048.html"},{"type":"UNKNOWN","source":"RHSA-2013:0544","url":"http://rhn.redhat.com/errata/RHSA-2013-0544.html","text":"http://rhn.redhat.com/errata/RHSA-2013-0544.html"},{"type":"UNKNOWN","source":"RHSA-2013:0548","url":"http://rhn.redhat.com/errata/RHSA-2013-0548.html","text":"http://rhn.redhat.com/errata/RHSA-2013-0548.html"},{"type":"UNKNOWN","source":"DSA-2783","url":"http://www.debian.org/security/2013/dsa-2783","text":"http://www.debian.org/security/2013/dsa-2783"},{"type":"UNKNOWN","source":"https://bugzilla.redhat.com/show_bug.cgi?id=895384","url":"https://bugzilla.redhat.com/show_bug.cgi?id=895384","text":"https://bugzilla.redhat.com/show_bug.cgi?id=895384"}],"modified_at":"2013-10-31T03:30:00Z","published_at":"2013-03-01T05:40:00Z","created_at":"2018-03-09T00:48:08.313Z","updated_at":"2019-01-29T08:54:59.385Z"},{"id":267937758,"external_id":"CVE-2011-5036","source":[{"id":1,"name":"NVD","description":"National Vulnerability Database","created_at":"2017-02-09T20:18:35.385Z","updated_at":"2017-02-13T20:12:05.342Z","attribution":"Copyright © 1999–2017, The MITRE Corporation. CVE and the CVE logo are registered trademarks and CVE-Compatible is a trademark of The MITRE Corporation.","license":"Submissions: For all materials you submit to the Common Vulnerabilities and Exposures (CVE®), you hereby grant to The MITRE Corporation (MITRE) and all CVE Numbering Authorities (CNAs) a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to reproduce, prepare derivative works of, publicly display, publicly perform, sublicense, and distribute such materials and derivative works. Unless required by applicable law or agreed to in writing, you provide such materials on an \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied, including, without limitation, any warranties or conditions of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE.\n\nCVE Usage: MITRE hereby grants you a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to reproduce, prepare derivative works of, publicly display, publicly perform, sublicense, and distribute Common Vulnerabilities and Exposures (CVE®). Any copy you make for such purposes is authorized provided that you reproduce MITRE's copyright designation and this license in any such copy.\n","copyright_url":"http://cve.mitre.org/about/termsofuse.html"}],"title":"CVE-2011-5036","summary":"Rack before 1.1.3, 1.2.x before 1.2.5, and 1.3.x before 1.3.6 computes hash values for form parameters without restricting the ability to trigger hash collisions predictably, which allows remote attackers to cause a denial of service (CPU consumption) by sending many crafted parameters.","score":"7.5","score_version":"3.0","score_system":"CVSS","score_details":{"cvssv2":{"vectorString":"AV:N/AC:L/Au:N/C:N/I:N/A:P","accessVector":"NETWORK","accessComplexity":"LOW","authentication":"NONE","confidentialityImpact":"NONE","integrityImpact":"NONE","availabilityImpact":"PARTIAL","baseScore":5}},"vector":"","access_complexity":"","vulnerability_authentication":"","confidentiality_impact":"","integrity_impact":"","availability_impact":"","vulnerabilty_source":"","assessment_check":null,"scanner":null,"recommendation":"","dependencies":null,"references":[{"type":"UNKNOWN","source":"20111228 n.runs-SA-2011.004 - web programming languages and platforms - DoS through hash table","url":"http://archives.neohapsis.com/archives/bugtraq/2011-12/0181.html","text":"http://archives.neohapsis.com/archives/bugtraq/2011-12/0181.html"},{"type":"UNKNOWN","source":"DSA-2783","url":"http://www.debian.org/security/2013/dsa-2783","text":"http://www.debian.org/security/2013/dsa-2783"},{"type":"UNKNOWN","source":"VU#903934","url":"http://www.kb.cert.org/vuls/id/903934","text":"http://www.kb.cert.org/vuls/id/903934"},{"type":"UNKNOWN","source":"http://www.nruns.com/_downloads/advisory28122011.pdf","url":"http://www.nruns.com/_downloads/advisory28122011.pdf","text":"http://www.nruns.com/_downloads/advisory28122011.pdf"},{"type":"UNKNOWN","source":"http://www.ocert.org/advisories/ocert-2011-003.html","url":"http://www.ocert.org/advisories/ocert-2011-003.html","text":"http://www.ocert.org/advisories/ocert-2011-003.html"},{"type":"UNKNOWN","source":"https://gist.github.com/52bbc6b9cc19ce330829","url":"https://gist.github.com/52bbc6b9cc19ce330829","text":"https://gist.github.com/52bbc6b9cc19ce330829"}],"modified_at":"2013-10-31T03:21:00Z","published_at":"2011-12-30T01:55:00Z","created_at":"2018-03-10T22:53:07.473Z","updated_at":"2018-11-24T10:05:23.28Z"}]}],"meta":{"vulnerability_count":2}}`
)
