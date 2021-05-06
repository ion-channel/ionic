package ionic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ion-channel/ionic/scanner"

	"github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestScanner(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Scanner", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should create an analysis for a project", func() {
			server.AddPath("/v1/scanner/analyzeProject").
				SetMethods("POST").
				SetPayload([]byte(SampleValidAnalysisStatus)).
				SetStatus(http.StatusOK)

			analysisStatus, err := client.AnalyzeProject("aprojectid", "ateamid", "abranch", "atoken")
			Expect(err).To(BeNil())
			Expect(analysisStatus.ID).To(Equal("analysis-id"))
			Expect(analysisStatus.Status).To(Equal("accepted"))
		})

		g.It("should create an analysis for a project whithout a branch", func() {
			server.AddPath("/v1/scanner/analyzeProject").
				SetMethods("POST").
				SetPayload([]byte(SampleValidAnalysisStatus)).
				SetStatus(http.StatusOK)

			analysisStatus, err := client.AnalyzeProject("aprojectid", "ateamid", "", "atoken")
			Expect(err).To(BeNil())
			Expect(analysisStatus.ID).To(Equal("analysis-id"))
			Expect(analysisStatus.Status).To(Equal("accepted"))
		})

		g.It("should get an analysis status for analysis", func() {
			server.AddPath("/v1/scanner/getAnalysisStatus").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysisScanStatus)).
				SetStatus(http.StatusOK)

			server.AddPath("/v1/ruleset/getAppliedRulesetForProject").
				SetMethods("GET").
				SetPayload([]byte(SampleAppliedRulesetForScanStatus)).
				SetStatus(http.StatusOK)

			analysisStatus, err := client.GetAnalysisStatus("analysis-id", "ateamid", "aprojectid", "atoken")
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

			analysisStatus, err := client.AddScanResult("analysis-id", "ateamid", "aprojectid", "coverage", "accepted", "atoken", scan)
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

			_, err = client.AddScanResult("analysis-id", "ateamid", "aprojectid", "coverage", "accepted", "atoken", scan)
			Expect(err).To(BeNil())
			Expect(server.HitRecords()[len(server.HitRecords())-1].Body).To(Equal([]byte(SampleRawExternalBody)))
		})

		g.It("should return a list of project scan states from requested projects", func() {
			server.AddPath("/v1/scanner/getProjectsStates").
				SetMethods("POST").
				SetPayload([]byte(SampleProjectsStates)).
				SetStatus(http.StatusOK)

			r, err := client.GetProjectsStates([]string{"aprojectid", "bprojectid"}, "", "sometoken")
			Expect(err).To(BeNil())
			Expect(len(r)).To(Equal(2))
			Expect(r[0].ID).To(Equal("2734ed5f-7af9-41ee-bd88-dbe7311cc81d"))
			Expect(r[0].Status).To(Equal("finished"))
			Expect(r[1].ID).To(Equal("5b975c65-8d9c-45f6-9341-aa8e8f75e261"))
			Expect(r[1].Status).To(Equal("errored"))
		})

		g.It("should analyze projects for a team", func() {
			server.AddPath("/v1/scanner/analyzeProject").
				SetMethods("POST").
				SetPayload([]byte(SampleValidAnalysisIDs)).
				SetStatus(http.StatusOK)

			ids, err := client.AnalyzeProjects("ateamid", "atoken", nil)
			Expect(err).To(BeNil())
			Expect(len(ids)).To(Equal(2))
		})

		g.It("should handle a team being analyzed recently", func() {
			server.AddPath("/v1/scanner/analyzeProject").
				SetMethods("POST").
				SetPayload([]byte(`{"message":"team has been analyzed recently","code":200}`)).
				SetStatus(http.StatusOK)

			ids, err := client.AnalyzeProjects("ateamid", "atoken", nil)
			Expect(err).To(BeNil())
			Expect(len(ids)).To(Equal(0))
		})
	})
}

const (
	SampleRawExternalBody             = `{"team_id":"ateamid","project_id":"aprojectid","analysis_id":"analysis-id","status":"","results":{"source":{"name":"","url":""},"notes":"","raw":{"big":["data"]}},"scan_type":"accepted"}`
	SampleRawExternalScanData         = `{"raw": {"big":["data"]}}`
	SampleValidAnalysisStatus         = `{"data":{"branch": "master","created_at": "2017-09-25T21:43:11.069Z","id": "analysis-id","message": "Request for analysis analysis-id on Amazon Web Services SDK has been accepted.","project_id": "93e2f31e-b579-4490-864d-7c630ac49720","status": "accepted","team_id": "team-id","updated_at": "2017-09-25T21:43:11.069Z"}}`
	SampleValidAnalysisIDs            = `{"data": ["0624c061-b39f-4138-a775-d3af53bc4a81", "d66e7c04-d2e9-490c-b2a5-7c66b7bb8445"]}`
	SampleValidAnalysisScanStatus     = `{"data":{"branch":"0.32","created_at":"2017-09-25T22:17:35.669Z","id":"analysis-id","message":"Completed compliance analysis","project_id":"project-id","scan_status":[{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:18:05.597Z","id":"2c7ed4ed-b0f2-ea2b-5057-9831f5f698bb","message":"Finished difference scan for Apache-Qpid, a difference was detected.","name":"difference","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:05.612Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:18:48.061Z","id":"21c64724-bca6-6c04-3196-30577bf5ad10","message":"Finished clamav scan for Apache-Qpid, found 0 infected files.","name":"virus","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:48.066Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:17:59.393Z","id":"9ab384a3-e127-5811-7f88-35074a9b96ac","message":"Finished dependency scan for Apache-Qpid, found 0 with no version and 0 with updates available.","name":"dependency","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:17:59.405Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:17:59.673Z","id":"9922e6c8-0149-4c69-4e40-da6a790c8df3","message":"Finished license scan for Apache-Qpid, found python license.","name":"license","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:17:59.685Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:18:00.605Z","id":"a84a574e-2f3e-aa0e-555c-05e7c3b37820","message":"Finished about_yml scan for Apache-Qpid, no valid .about.yml found.","name":"about_yml","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:00.617Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:18:00.997Z","id":"b4db0492-6832-8333-cae0-b331f199d855","message":"Finished vulnerability scan for Apache-Qpid, found 0 vulnerabilities.","name":"vulnerability","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:01.019Z"},{"analysis_status_id":"analysis-id","created_at":"2017-09-25T22:18:26.524Z","id":"a7a599cb-5fa4-e4c2-5b47-cddc094f4aac","message":"Finished ecosystems scan for Apache-Qpid, found {\"Shell\"=>347100, \"Python\"=>2565596, \"CMake\"=>202651, \"C++\"=>7472076, \"Ruby\"=>305919, \"PowerShell\"=>67938, \"C#\"=>449894, \"C\"=>70609, \"Perl\"=>86022, \"Perl 6\"=>2703, \"Gherkin\"=>15299, \"CSS\"=>80719, \"HTML\"=>379968, \"Roff\"=>10501, \"Logos\"=>2329, \"Emacs Lisp\"=>7379, \"Makefile\"=>14373, \"XQuery\"=>234, \"XSLT\"=>103182, \"M4\"=>18570, \"Java\"=>16313719, \"JavaScript\"=>1184295, \"Batchfile\"=>12789, \"Smarty\"=>19168, \"Gnuplot\"=>2182, \"Objective-C\"=>1544} ecosystems in project.","name":"ecosystems","project_id":"project-id","read":"f","status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:26.528Z"}],"status":"finished","team_id":"team-id","updated_at":"2017-09-25T22:18:48.619Z"}}`
	SampleProjectsStates              = `{"data":[{"id":"2734ed5f-7af9-41ee-bd88-dbe7311cc81d","status":"finished"},{"id":"5b975c65-8d9c-45f6-9341-aa8e8f75e261","status":"errored"}],"meta":{"total_count":0,"offset":0,"last_update":"0001-01-01T00:00:00Z"}}`
	SampleAppliedRulesetForScanStatus = `{
		"data": {
			"project_id": "32D701E1-E173-43EF-9CC8-E4CB27417FD8",
			"team_id": "800E898B-CCD8-4394-A559-F17D08030413",
			"analysis_id": "B061D58B-FDFD-46BF-A766-2D38DE3B1D7B",
			"rule_evaluation_summary": {
				"summary": "finished",
				"passed": true,
				"status": "finished",
				"ruleresults": [
					{
						"id": "f9eec625-88d9-fca1-02db-d5062957ced5",
						"analysis_id": "B061D58B-FDFD-46BF-A766-2D38DE3B1D7B",
						"team_id": "800E898B-CCD8-4394-A559-F17D08030413",
						"project_id": "32D701E1-E173-43EF-9CC8-E4CB27417FD8",
						"description": "some description",
						"name": "License",
						"summary": "Finished license scan for a-ionmock, failed to detect license.",
						"created_at": "2017-09-27T12:55:34.480Z",
						"updated_at": "2017-09-27T12:55:34.480Z",
						"results": {
							"license": {
								"license": {
									"name": "Not found",
									"type": []
								}
							}
						},
						"duration": 1.10026499987725,
						"passed": false,
						"risk": "n/a",
						"type": "Not Evaluated"
					},
					{
						"id": "e0eb6936-9074-6f03-e861-ae65290fa3c3",
						"analysis_id": "B061D58B-FDFD-46BF-A766-2D38DE3B1D7B",
						"team_id": "800E898B-CCD8-4394-A559-F17D08030413",
						"project_id": "32D701E1-E173-43EF-9CC8-E4CB27417FD8",
						"description": "some description",
						"name": "Ecosystems",
						"summary": "Finished ecosystems scan for a-ionmock, found {\"Java\"=>2582} ecosystems in project.",
						"created_at": "2017-09-27T12:55:34.503Z",
						"updated_at": "2017-09-27T12:55:34.503Z",
						"results": {
							"ecosystems": {
								"Java": 2582
							}
						},
						"duration": 16.95143500001,
						"passed": false,
						"risk": "n/a",
						"type": "Not Evaluated"
					},
					{
						"id": "d1035d70-6516-aa94-faa4-bf77b06bfa82",
						"analysis_id": "B061D58B-FDFD-46BF-A766-2D38DE3B1D7B",
						"team_id": "800E898B-CCD8-4394-A559-F17D08030413",
						"project_id": "32D701E1-E173-43EF-9CC8-E4CB27417FD8",
						"description": "some description",
						"name": "Difference",
						"summary": "Finished difference scan for a-ionmock, a difference was detected.",
						"created_at": "2017-09-27T12:55:34.783Z",
						"updated_at": "2017-09-27T12:55:34.783Z",
						"results": {
							"difference": {
								"difference": true,
								"checksum": "d63371f4cea3a8b80fc7838764e448955cc8ff32bdb41d06ba6055b98883380b"
							}
						},
						"duration": 542.002373999821,
						"passed": false,
						"risk": "n/a",
						"type": "Not Evaluated"
					},
					{
						"id": "02ce4f55-6038-05f0-0303-e5e43b36beed",
						"analysis_id": "B061D58B-FDFD-46BF-A766-2D38DE3B1D7B",
						"team_id": "800E898B-CCD8-4394-A559-F17D08030413",
						"project_id": "32D701E1-E173-43EF-9CC8-E4CB27417FD8",
						"description": "some description",
						"name": "About_yml",
						"summary": "Finished about_yml scan for a-ionmock, valid .about.yml found.",
						"created_at": "2017-09-27T12:55:35.354Z",
						"updated_at": "2017-09-27T12:55:35.354Z",
						"results": {
							"about_yml": {
								"message": "",
								"valid": true,
								"content": "---\n# .about.yml project metadata\n#\n# Copy this template into your project repository's root directory as\n# .about.yml and fill in the fields as described below.\n\n# This is a short name of your project that can be used as a URL slug.\n# (required)\nname: ionmockjavaapp\n\n# This is the display name of your project. (required)\nfull_name: ionmockjavaapp\n\n# What is the problem your project solves? What is the solution? Use the\n# format shown below. The #dashboard team will gladly help you put this\n# together for your project. (required)\ndescription: Provides a test harness for java (maven) projects\n\n# What is the measurable impact of your project? Use the format shown below.\n# The #dashboard team will gladly help you put this together for your project.\n# (required)\nimpact: high\n\n# What kind of team owns the repository? (required)\n# values: guild, working-group, project\nowner_type: project\n\n# What is your project's current status? (required)\n# values: discovery, alpha, beta, live\nstage: live\n\n# Should this repo have automated tests? If so, set to true. (required)\n# values: true, false\ntestable: true\n\nlicenses:\n  doozer:\n    name: GPLV2\n    url: https://github.com/ion-channel/java-lew/blob/master/license.txt\n\nteam:\n- github: kitplummer\n  role: lead\n"
							}
						},
						"duration": 1168.83430600001,
						"passed": false,
						"risk": "n/a",
						"type": "Not Evaluated"
					},
					{
						"id": "504ea20a-366e-ef90-0723-6febb6f350a1",
						"analysis_id": "B061D58B-FDFD-46BF-A766-2D38DE3B1D7B",
						"team_id": "800E898B-CCD8-4394-A559-F17D08030413",
						"project_id": "32D701E1-E173-43EF-9CC8-E4CB27417FD8",
						"description": "some description",
						"name": "Dependency",
						"summary": "Finished dependency scan for a-ionmock, found 0 with no version and 2 with updates available.",
						"created_at": "2017-09-27T12:55:48.484Z",
						"updated_at": "2017-09-27T12:55:48.484Z",
						"results": {
							"dependency": {
								"dependencies": [
									{
										"latest_version": "2.0",
										"org": "net.sourceforge.javacsv",
										"name": "javacsv",
										"type": "maven",
										"package": "jar",
										"version": "2.0",
										"scope": "compile"
									},
									{
										"latest_version": "4.12",
										"org": "junit",
										"name": "junit",
										"type": "maven",
										"package": "jar",
										"version": "4.11",
										"scope": "test"
									},
									{
										"latest_version": "1.4-atlassian-1",
										"org": "org.hamcrest",
										"name": "hamcrest-core",
										"type": "maven",
										"package": "jar",
										"version": "1.3",
										"scope": "test"
									},
									{
										"latest_version": "4.5.2",
										"org": "org.apache.httpcomponents",
										"name": "httpclient",
										"type": "maven",
										"package": "jar",
										"version": "4.3.4",
										"scope": "compile"
									},
									{
										"latest_version": "4.4.5",
										"org": "org.apache.httpcomponents",
										"name": "httpcore",
										"type": "maven",
										"package": "jar",
										"version": "4.3.2",
										"scope": "compile"
									},
									{
										"latest_version": "99.0-does-not-exist",
										"org": "commons-logging",
										"name": "commons-logging",
										"type": "maven",
										"package": "jar",
										"version": "1.1.3",
										"scope": "compile"
									},
									{
										"latest_version": "20041127.091804",
										"org": "commons-codec",
										"name": "commons-codec",
										"type": "maven",
										"package": "jar",
										"version": "1.6",
										"scope": "compile"
									}
								],
								"meta": {
									"first_degree_count": 3,
									"no_version_count": 0,
									"total_unique_count": 7,
									"update_available_count": 2
								}
							}
						},
						"duration": 14287.5910550001,
						"passed": false,
						"risk": "n/a",
						"type": "Not Evaluated"
					},
					{
						"id": "3abfd693-5f61-e4ac-ec72-763d42dfb4fb",
						"analysis_id": "B061D58B-FDFD-46BF-A766-2D38DE3B1D7B",
						"team_id": "800E898B-CCD8-4394-A559-F17D08030413",
						"project_id": "32D701E1-E173-43EF-9CC8-E4CB27417FD8",
						"description": "some description",
						"name": "Vulnerability",
						"summary": "Finished vulnerability scan for a-ionmock, found 0 vulnerabilities.",
						"created_at": "2017-09-27T12:55:48.719Z",
						"updated_at": "2017-09-27T12:55:48.719Z",
						"results": {
							"vulnerabilities": {
								"vulnerabilities": [],
								"meta": {
									"vulnerability_count": 0
								}
							}
						},
						"duration": 89.9386969999796,
						"passed": false,
						"risk": "n/a",
						"type": "Not Evaluated"
					},
					{
						"id": "ce49954f-02d3-9675-380a-eb974ab8b68d",
						"analysis_id": "B061D58B-FDFD-46BF-A766-2D38DE3B1D7B",
						"team_id": "800E898B-CCD8-4394-A559-F17D08030413",
						"project_id": "32D701E1-E173-43EF-9CC8-E4CB27417FD8",
						"description": "some description",
						"name": "Virus",
						"summary": "Finished clamav scan for a-ionmock, found 0 infected files.",
						"created_at": "2017-09-27T12:55:50.525Z",
						"updated_at": "2017-09-27T12:55:50.525Z",
						"results": {
							"clam_av_details": {
								"clamav_version": "ClamAV 0.99.2",
								"clamav_db_version": "Wed Sep 27 04:44:38 2017\n"
							},
							"clamav": {
								"known_viruses": 6303819,
								"engine_version": "0.99.2",
								"scanned_directories": 29,
								"scanned_files": 30,
								"infected_files": 0,
								"data_scanned": "0.04 MB",
								"data_read": "0.02 MB (ratio 1.83:1)",
								"time": "16.144 sec (0 m 16 s)",
								"file_notes": {}
							}
						},
						"duration": 16180.6532439996,
						"passed": true,
						"risk": "n/a",
						"type": "Not Evaluated"
					}
				]
			},
			"rule_eval_created_at": "2017-09-27T12:55:50+00:00",
			"created_at": "2017-09-27T12:55:50.814Z",
			"updated_at": "2017-09-27T12:55:50.814Z",
			"passed": true
		}
	}`
)
