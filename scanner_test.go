package ionic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ion-channel/ionic/scanner"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestScanner(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Scanner", func() {
		server := bogus.New()
		server.Start()
		h, p := server.HostPort()
		client, _ := New("", fmt.Sprintf("http://%v:%v", h, p))

		g.It("should create an analysis for a project", func() {
			server.AddPath("/v1/scanner/analyzeProject").
				SetMethods("POST").
				SetPayload([]byte(SampleValidAnalysisStatus)).
				SetStatus(http.StatusOK)

			analysisStatus, err := client.AnalyzeProject("aprojectid", "ateamid", "abranch")
			Expect(err).To(BeNil())
			Expect(analysisStatus.ID).To(Equal("analysis-id"))
			Expect(analysisStatus.Status).To(Equal("accepted"))
		})

		g.It("should create an analysis for a project whithout a branch", func() {
			server.AddPath("/v1/scanner/analyzeProject").
				SetMethods("POST").
				SetPayload([]byte(SampleValidAnalysisStatus)).
				SetStatus(http.StatusOK)

			analysisStatus, err := client.AnalyzeProject("aprojectid", "ateamid", "")
			Expect(err).To(BeNil())
			Expect(analysisStatus.ID).To(Equal("analysis-id"))
			Expect(analysisStatus.Status).To(Equal("accepted"))
		})

		g.It("should get an analysis status for analysis", func() {
			server.AddPath("/v1/scanner/getAnalysisStatus").
				SetMethods("Get").
				SetPayload([]byte(SampleValidAnalysisScanStatus)).
				SetStatus(http.StatusOK)

			analysisStatus, err := client.GetAnalysisStatus("analysis-id", "ateamid", "aprojectid")
			Expect(err).To(BeNil())
			Expect(analysisStatus.ID).To(Equal("analysis-id"))
			Expect(analysisStatus.Message).To(Equal("Completed compliance analysis"))
			Expect(analysisStatus.Status).To(Equal("finished"))
			Expect(analysisStatus.ScanStatus[0].Status).To(Equal("finished"))
			Expect(analysisStatus.ScanStatus[0].Message).To(Equal("Finished difference scan for Apache-Qpid, a difference was detected."))
		})

		g.It("should add a scan to an analysis", func() {
			server.AddPath("/v1/scanner/addScanResult").
				SetMethods("POST").
				SetPayload([]byte(SampleValidAnalysisStatus)).
				SetStatus(http.StatusOK)

			coverage := scanner.ExternalCoverage{Value: 99.99}
			scan := scanner.ExternalScan{}
			scan.Coverage = &coverage

			analysisStatus, err := client.AddScanResult("analysis-id", "ateamid", "aprojectid", "coverage", "accepted", scan)
			Expect(err).To(BeNil())
			Expect(analysisStatus.ID).To(Equal("analysis-id"))
			Expect(analysisStatus.Status).To(Equal("accepted"))
		})

		g.It("should add a scan with raw data to an analysis", func() {
			server.AddPath("/v1/scanner/addScanResult").
				SetMethods("POST").
				SetPayload([]byte(SampleValidAnalysisStatus)).
				SetStatus(http.StatusOK)

			scan := scanner.ExternalScan{}
			err := json.Unmarshal([]byte(SampleRawExternalScanData), &scan)
			Expect(err).To(BeNil())

			_, err = client.AddScanResult("analysis-id", "ateamid", "aprojectid", "coverage", "accepted", scan)
			Expect(err).To(BeNil())
			Expect(server.HitRecords()[len(server.HitRecords())-1].Body).To(Equal([]byte(SampleRawExternalBody)))
		})
	})
}

const (
	SampleRawExternalBody         = `{"team_id":"ateamid","project_id":"aprojectid","analysis_id":"analysis-id","status":"","results":{"source":{"name":"","url":""},"notes":"","raw":{"big":["data"]}},"scan_type":"accepted"}`
	SampleRawExternalScanData     = `{"raw": {"big":["data"]}}`
	SampleValidAnalysisStatus     = `{"data":{"branch": "master","build_number": null,"created_at": "2017-09-25T21:43:11.069Z","id": "analysis-id","message": "Request for analysis analysis-id on Amazon Web Services SDK has been accepted.","project_id": "93e2f31e-b579-4490-864d-7c630ac49720","status": "accepted","team_id": "team-id","updated_at": "2017-09-25T21:43:11.069Z"}}`
	SampleValidAnalysisScanStatus = `{"data":{"branch":"0.32","build_number":null,"created_at":"2017-09-25T22:17:35.669Z","id":"analysis-id","message":"Completed compliance analysis","project_id":"project-id","scan_status":[{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:18:05.597Z","id":"2c7ed4ed-b0f2-ea2b-5057-9831f5f698bb","message":"Finished difference scan for Apache-Qpid, a difference was detected.","name":"difference","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:05.612Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:18:48.061Z","id":"21c64724-bca6-6c04-3196-30577bf5ad10","message":"Finished clamav scan for Apache-Qpid, found 0 infected files.","name":"virus","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:48.066Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:17:59.393Z","id":"9ab384a3-e127-5811-7f88-35074a9b96ac","message":"Finished dependency scan for Apache-Qpid, found 0 with no version and 0 with updates available.","name":"dependency","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:17:59.405Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:17:59.673Z","id":"9922e6c8-0149-4c69-4e40-da6a790c8df3","message":"Finished license scan for Apache-Qpid, found python license.","name":"license","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:17:59.685Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:18:00.605Z","id":"a84a574e-2f3e-aa0e-555c-05e7c3b37820","message":"Finished about_yml scan for Apache-Qpid, no valid .about.yml found.","name":"about_yml","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:00.617Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:18:00.997Z","id":"b4db0492-6832-8333-cae0-b331f199d855","message":"Finished vulnerability scan for Apache-Qpid, found 0 vulnerabilities.","name":"vulnerability","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:01.019Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:18:26.524Z","id":"a7a599cb-5fa4-e4c2-5b47-cddc094f4aac","message":"Finished ecosystems scan for Apache-Qpid, found {\"Shell\"=>347100, \"Python\"=>2565596, \"CMake\"=>202651, \"C++\"=>7472076, \"Ruby\"=>305919, \"PowerShell\"=>67938, \"C#\"=>449894, \"C\"=>70609, \"Perl\"=>86022, \"Perl 6\"=>2703, \"Gherkin\"=>15299, \"CSS\"=>80719, \"HTML\"=>379968, \"Roff\"=>10501, \"Logos\"=>2329, \"Emacs Lisp\"=>7379, \"Makefile\"=>14373, \"XQuery\"=>234, \"XSLT\"=>103182, \"M4\"=>18570, \"Java\"=>16313719, \"JavaScript\"=>1184295, \"Batchfile\"=>12789, \"Smarty\"=>19168, \"Gnuplot\"=>2182, \"Objective-C\"=>1544} ecosystems in project.","name":"ecosystems","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:26.528Z"}],"status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:48.619Z"}}`
)
