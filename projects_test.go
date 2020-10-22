package ionic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	"github.com/ion-channel/ionic/projects"
	. "github.com/onsi/gomega"
)

func TestProjects(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Projects", func() {
		var server *bogus.Bogus
		var h, p string
		var client *IonClient

		g.BeforeEach(func() {
			server = bogus.New()
			h, p = server.HostPort()
			client, _ = New(fmt.Sprintf("http://%v:%v", h, p))
		})

		g.It("should create a project", func() {
			project := &projects.Project{}
			server.AddPath("/v1/project/createProject").
				SetMethods("POST").
				SetPayload([]byte(SampleValidProject)).
				SetStatus(http.StatusCreated)
			server.AddPath("/v1/ruleset/getRuleset").
				SetMethods("HEAD").
				SetStatus(http.StatusOK)

			project, err := client.CreateProject(project, "bef86653-1926-4990-8ef8-5f26cd59d6fc", "")
			Expect(err).To(BeNil())
			Expect(*project.ID).To(Equal("334c183d-4d37-4515-84c4-0d0ed0fb8db0"))
			Expect(*project.Name).To(Equal("Statler"))
		})

		g.It("should get a project", func() {
			server.AddPath("/v1/project/getProject").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProject)).
				SetStatus(http.StatusOK)

			project, err := client.GetProject("334c183d-4d37-4515-84c4-0d0ed0fb8db0", "bef86653-1926-4990-8ef8-5f26cd59d6fc", "")
			Expect(err).To(BeNil())
			Expect(*project.ID).To(Equal("334c183d-4d37-4515-84c4-0d0ed0fb8db0"))
			Expect(*project.Name).To(Equal("Statler"))
		})

		g.It("should get a raw project", func() {
			server.AddPath("/v1/project/getProject").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProject)).
				SetStatus(http.StatusOK)

			raw, err := client.GetRawProject("334c183d-4d37-4515-84c4-0d0ed0fb8db0", "bef86653-1926-4990-8ef8-5f26cd59d6fc", "")
			Expect(err).To(BeNil())
			Expect(raw).To(Equal(json.RawMessage(SampleValidRawProject)))
		})

		g.It("should get all projects", func() {
			server.AddPath("/v1/project/getProjects").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProjects)).
				SetStatus(http.StatusOK)

			projects, err := client.GetProjects("bef86653-1926-4990-8ef8-5f26cd59d6fc", "", nil, nil)
			Expect(err).To(BeNil())
			Expect(len(projects)).To(Equal(2))
			Expect(*projects[0].ID).To(Equal("334c183d-4d37-4515-84c4-0d0ed0fb8db0"))
			Expect(*projects[0].Name).To(Equal("Statler"))
		})

		g.It("should get a project by the url", func() {
			server.AddPath("/v1/project/getProjectByUrl").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProject)).
				SetStatus(http.StatusOK)

			project, err := client.GetProjectByURL("git@github.com:ion-channel/statler.git", "bef86653-1926-4990-8ef8-5f26cd59d6fc", "")
			Expect(err).To(BeNil())
			Expect(*project.ID).To(Equal("334c183d-4d37-4515-84c4-0d0ed0fb8db0"))
			Expect(*project.Name).To(Equal("Statler"))

			hr := server.HitRecords()
			Expect(len(hr)).To(Equal(1))
			Expect(hr[0].Query.Get("url")).To(Equal("git@github.com:ion-channel/statler.git"))
			Expect(hr[0].Query.Get("team_id")).To(Equal("bef86653-1926-4990-8ef8-5f26cd59d6fc"))
		})

		g.It("should create projects from a csv", func() {
			server.AddPath("/v1/project/createProjectsCSV").
				SetMethods("POST").
				SetPayload([]byte(SampleValidCSVProjects)).
				SetStatus(http.StatusCreated)

			resp, err := client.CreateProjectsFromCSV("./README.md", "someteamid", "")
			Expect(err).To(BeNil())
			Expect(len(resp.Projects)).To(Equal(9))
			Expect(len(resp.Errors)).To(Equal(0))

			hr := server.HitRecords()
			Expect(hr[0].Verb).To(Equal("POST"))
			Expect(hr[0].Query.Get("team_id")).To(Equal("someteamid"))
		})

		g.It("should return errors from a csv", func() {
			server.AddPath("/v1/project/createProjectsCSV").
				SetMethods("POST").
				SetPayload([]byte(SampleErrorCSVProjects)).
				SetStatus(http.StatusCreated)

			resp, err := client.CreateProjectsFromCSV("./README.md", "someteamid", "")
			Expect(err).To(BeNil())
			Expect(len(resp.Projects)).To(Equal(0))
			Expect(len(resp.Errors)).To(Equal(9))

			hr := server.HitRecords()
			Expect(hr[0].Verb).To(Equal("POST"))
			Expect(hr[0].Query.Get("team_id")).To(Equal("someteamid"))
		})

		g.It("should return ruleset ids from team", func() {
			server.AddPath("/v1/project/getUsedRulesetIds").
				SetMethods("GET").
				SetPayload([]byte(SampleRulesetIds)).
				SetStatus(http.StatusOK)

			resp, err := client.GetUsedRulesetIds("someId", "sometoken")
			Expect(err).To(BeNil())
			Expect(len(resp)).To(Equal(1))
			Expect(resp[0].RulesetID).To(Equal("04d95a49-2df5-4204-bed9-d0dcd0438ef7"))
		})

		g.It("should return names and versions for projects", func() {
			server.AddPath("/v1/project/getNamesVersions").
				SetMethods("POST").
				SetPayload([]byte(SampleNamesVersions)).
				SetStatus(http.StatusOK)

			resp, err := client.GetNamesVersions("someId", []string{"sometoken"}, "token")
			Expect(err).To(BeNil())
			Expect(len(resp)).To(Equal(4))
			Expect(resp[0].ID).To(Equal("d73559a2-7cc8-430f-85fc-7dd7dff71695"))
			Expect(resp[0].Name).To(Equal("bar"))
			Expect(resp[0].Version).To(Equal(""))
		})
	})
}

const (
	SampleValidProject     = `{"data":{"active":true,"aliases":[],"branch":"master","chat_channel":"foo","created_at":"2016-08-29T17:38:40.401Z","deploy_key":null,"description":"Statler Travis CI testing","id":"334c183d-4d37-4515-84c4-0d0ed0fb8db0","key_fingerprint":"","name":"Statler","password":null,"poc_email":"","poc_email_hash":"","poc_name":"","poc_name_hash":"","ruleset_id":"f7583ed9-c939-4b51-a865-394cc8ddcffa","should_monitor":false,"source":"git@github.com:ion-channel/statler.git","tags":[],"team_id":"bef86653-1926-4990-8ef8-5f26cd59d6fc","type":"git","updated_at":"2017-05-22T18:00:54.982Z","username":null}}`
	SampleValidRawProject  = `{"active":true,"aliases":[],"branch":"master","chat_channel":"foo","created_at":"2016-08-29T17:38:40.401Z","deploy_key":null,"description":"Statler Travis CI testing","id":"334c183d-4d37-4515-84c4-0d0ed0fb8db0","key_fingerprint":"","name":"Statler","password":null,"poc_email":"","poc_email_hash":"","poc_name":"","poc_name_hash":"","ruleset_id":"f7583ed9-c939-4b51-a865-394cc8ddcffa","should_monitor":false,"source":"git@github.com:ion-channel/statler.git","tags":[],"team_id":"bef86653-1926-4990-8ef8-5f26cd59d6fc","type":"git","updated_at":"2017-05-22T18:00:54.982Z","username":null}`
	SampleValidProjects    = `{"data":[{"active":true,"aliases":[],"branch":"master","chat_channel":"foo","created_at":"2016-08-29T17:38:40.401Z","deploy_key":null,"description":"Statler Travis CI testing","id":"334c183d-4d37-4515-84c4-0d0ed0fb8db0","key_fingerprint":"","name":"Statler","password":null,"poc_email":"","poc_email_hash":"","poc_name":"","poc_name_hash":"","ruleset_id":"f7583ed9-c939-4b51-a865-394cc8ddcffa","should_monitor":false,"source":"git@github.com:ion-channel/statler.git","tags":[],"team_id":"bef86653-1926-4990-8ef8-5f26cd59d6fc","type":"git","updated_at":"2017-05-22T18:00:54.982Z","username":null},{"active":false,"aliases":[],"branch":"master","chat_channel":"bar","created_at":"2017-08-29T17:38:40.401Z","deploy_key":null,"description":"Statler Travis CI testing","id":"334c183d-4d37-4515-84c4-0d0ed0fb8db0","key_fingerprint":"","name":"baz","password":null,"poc_email":"","poc_email_hash":"","poc_name":"","poc_name_hash":"","ruleset_id":"f7583ed9-c939-4b51-a865-394cc8ddcffa","should_monitor":false,"source":"git@github.com:ion-channel/baz.git","tags":[],"team_id":"bef86653-1926-4990-8ef8-5f26cd59d6fc","type":"git","updated_at":"2018-05-22T18:00:54.982Z","username":null}]}`
	SampleValidCSVProjects = `{"data":{"projects":[{"id":"d02da880-ea1d-4b06-8370-95c677d4e796","team_id":"fead431b-28e7-4a1e-bd84-ffd35069ebcc","ruleset_id":"51b0e779-0a0b-4ce5-a828-62c1eaa9a454","name":"node-glob","type":"git","source":"https://github.com/isaacs/node-glob","branch":"master","description":null,"active":true,"chat_channel":null,"created_at":"2019-02-06T19:54:05.749Z","updated_at":"2019-02-06T19:54:05.749Z","deploy_key":null,"should_monitor":true,"poc_name":"Megan Benton","poc_email":"megan.benton@ionchannel.io","username":null,"password":null,"key_fingerprint":null,"poc_name_hash":"b30bcd9d0b4d9d71b92017d18e4e3d72bfbfdce444990b5718f11f6970e74c24","poc_email_hash":"9e0c4a5b04f261424653c7d1e5354fa315af0654106a6fc45a5780c6b22056b7"},{"id":"b3a0e5ac-27d9-4c18-8db7-fc4714e66db7","team_id":"fead431b-28e7-4a1e-bd84-ffd35069ebcc","ruleset_id":"51b0e779-0a0b-4ce5-a828-62c1eaa9a454","name":"is-retry-allowed","type":"git","source":"https://github.com/floatdrop/is-retry-allowed","branch":"master","description":null,"active":true,"chat_channel":null,"created_at":"2019-02-06T19:54:05.749Z","updated_at":"2019-02-06T19:54:05.749Z","deploy_key":null,"should_monitor":true,"poc_name":"Megan Benton","poc_email":"megan.benton@ionchannel.io","username":null,"password":null,"key_fingerprint":null,"poc_name_hash":"02b9188e3c0cc43aeb165c7c377628b5968d22017969366bc9d73b3faa13c3f2","poc_email_hash":"86cd142afd157ad8d7536526d40c0669e6f6b97067f1564be5f62ee40252a793"},{"id":"6f95b836-0f73-49f5-ba08-800fe917ccd0","team_id":"fead431b-28e7-4a1e-bd84-ffd35069ebcc","ruleset_id":"51b0e779-0a0b-4ce5-a828-62c1eaa9a454","name":"node-jslint","type":"git","source":"https://github.com/reid/node-jslint","branch":"master","description":null,"active":true,"chat_channel":null,"created_at":"2019-02-06T19:54:05.749Z","updated_at":"2019-02-06T19:54:05.749Z","deploy_key":null,"should_monitor":true,"poc_name":"Megan Benton","poc_email":"megan.benton@ionchannel.io","username":null,"password":null,"key_fingerprint":null,"poc_name_hash":"e02adff8edbe5e64086ff5b1208a30c39d4c6564b62bc5ecafedd5d72fc95610","poc_email_hash":"c8cd36d9fce78f25f8b376890b3dfa336b932d1e5e396448d56ea6cb3f42807a"},{"id":"4cf7a587-8b02-4239-b6b7-139d935bf0fe","team_id":"fead431b-28e7-4a1e-bd84-ffd35069ebcc","ruleset_id":"51b0e779-0a0b-4ce5-a828-62c1eaa9a454","name":"ieee754","type":"git","source":"https://github.com/feross/ieee754","branch":"master","description":null,"active":true,"chat_channel":null,"created_at":"2019-02-06T19:54:05.749Z","updated_at":"2019-02-06T19:54:05.749Z","deploy_key":null,"should_monitor":true,"poc_name":"Megan Benton","poc_email":"megan.benton@ionchannel.io","username":null,"password":null,"key_fingerprint":null,"poc_name_hash":"96fa874e8a63e2a677a26f5b53d508d4de810bb5be17aa9301b94e40854f8634","poc_email_hash":"a8d8debd1f60a5a5a488622623902360d96ef00548d7bc170a15ee337d130304"},{"id":"da65083c-d22d-454a-bfb7-d5825e80ed2b","team_id":"fead431b-28e7-4a1e-bd84-ffd35069ebcc","ruleset_id":"51b0e779-0a0b-4ce5-a828-62c1eaa9a454","name":"lowercase-keys","type":"git","source":"https://github.com/sindresorhus/lowercase-keys","branch":"master","description":null,"active":true,"chat_channel":null,"created_at":"2019-02-06T19:54:05.749Z","updated_at":"2019-02-06T19:54:05.749Z","deploy_key":null,"should_monitor":true,"poc_name":"Megan Benton","poc_email":"megan.benton@ionchannel.io","username":null,"password":null,"key_fingerprint":null,"poc_name_hash":"7c60934cf493d93258ed4a0b5527a56a4d8a2461270f593c001fe8a2dc248557","poc_email_hash":"1fa16e162df3960649ec763e1998611c9e3960cb3104e414c098b4c34029623c"},{"id":"0d9d389b-01fb-4e20-8c26-1c75bd87235e","team_id":"fead431b-28e7-4a1e-bd84-ffd35069ebcc","ruleset_id":"51b0e779-0a0b-4ce5-a828-62c1eaa9a454","name":"isarray","type":"git","source":"https://github.com/juliangruber/isarray","branch":"master","description":null,"active":true,"chat_channel":null,"created_at":"2019-02-06T19:54:05.749Z","updated_at":"2019-02-06T19:54:05.749Z","deploy_key":null,"should_monitor":true,"poc_name":"Megan Benton","poc_email":"megan.benton@ionchannel.io","username":null,"password":null,"key_fingerprint":null,"poc_name_hash":"c6b54f7b3efe1aaa578ae3be5912202d8025c37db5f4e70690f28ea1766f96d8","poc_email_hash":"ec75bd203cc515a6fb622ca1b01f1cf4d08b44f4d34c1d5bfa74425dd682cb26"},{"id":"1e12d6c0-f48e-4e7b-a81e-a44cf1e4096a","team_id":"fead431b-28e7-4a1e-bd84-ffd35069ebcc","ruleset_id":"51b0e779-0a0b-4ce5-a828-62c1eaa9a454","name":"is-redirect","type":"git","source":"https://github.com/sindresorhus/is-redirect","branch":"master","description":null,"active":true,"chat_channel":null,"created_at":"2019-02-06T19:54:05.749Z","updated_at":"2019-02-06T19:54:05.749Z","deploy_key":null,"should_monitor":true,"poc_name":"Megan Benton","poc_email":"megan.benton@ionchannel.io","username":null,"password":null,"key_fingerprint":null,"poc_name_hash":"33e4d3c953a79b0df7f78500e3ef7059fb1c26135003689ae81dcb4972999d4c","poc_email_hash":"94a97ac9411d35ca255bad0e3f91c5cd144544097d30c049568a4ecd84817d46"},{"id":"5123ca79-c505-406f-bc8d-1e151fccdd14","team_id":"fead431b-28e7-4a1e-bd84-ffd35069ebcc","ruleset_id":"51b0e779-0a0b-4ce5-a828-62c1eaa9a454","name":"is-stream","type":"git","source":"https://github.com/sindresorhus/is-stream","branch":"master","description":null,"active":true,"chat_channel":null,"created_at":"2019-02-06T19:54:05.749Z","updated_at":"2019-02-06T19:54:05.749Z","deploy_key":null,"should_monitor":true,"poc_name":"Megan Benton","poc_email":"megan.benton@ionchannel.io","username":null,"password":null,"key_fingerprint":null,"poc_name_hash":"5af4763fb0bf9fd88e05c9105706061f450e36be7900aa5ea33eab16e72c4a50","poc_email_hash":"c947285b72be9e8e8df328f75c647124d56d8b3f0b06be934005a2b12be26b08"},{"id":"819221d0-de0c-4845-9866-706284574ac8","team_id":"fead431b-28e7-4a1e-bd84-ffd35069ebcc","ruleset_id":"51b0e779-0a0b-4ce5-a828-62c1eaa9a454","name":"jmespath.js","type":"git","source":"https://github.com/jmespath/jmespath.js","branch":"master","description":"Javascript implementation of JMESPath, a query language for JSON http://jmespath.org","active":true,"chat_channel":null,"created_at":"2019-02-06T19:54:05.749Z","updated_at":"2019-02-06T19:54:05.749Z","deploy_key":null,"should_monitor":true,"poc_name":"Megan Benton","poc_email":"megan.benton@ionchannel.io","username":null,"password":null,"key_fingerprint":null,"poc_name_hash":"d05983513b96ba7b1a429fa94323dc714e58a7e4d204ef079b0a389bfacbc4cb","poc_email_hash":"a44d277e391d933196033020acd3b317df04461308b932f902861a1e2d376493"}],"errors":[]}}`
	SampleErrorCSVProjects = `{"data":{"projects":[],"errors":[{"message":"Project already exists with url and branch (https://github.com/floatdrop/is-retry-allowed,master)"},{"message":"Project already exists with url and branch (https://github.com/reid/node-jslint,master)"},{"message":"Project already exists with url and branch (https://github.com/feross/ieee754,master)"},{"message":"Project already exists with url and branch (https://github.com/sindresorhus/lowercase-keys,master)"},{"message":"Project already exists with url and branch (https://github.com/juliangruber/isarray,master)"},{"message":"Project already exists with url and branch (https://github.com/isaacs/node-glob,master)"},{"message":"Project already exists with url and branch (https://github.com/sindresorhus/is-redirect,master)"},{"message":"Project already exists with url and branch (https://github.com/sindresorhus/is-stream,master)"},{"message":"Project already exists with url and branch (https://github.com/jmespath/jmespath.js,master)"}]}}`
	SampleRulesetIds       = `{"data":[{"ruleset_id":"04d95a49-2df5-4204-bed9-d0dcd0438ef7"}],"meta":{"total_count":1,"offset":0,"last_update":"2020-08-18T21:30:34.0292132Z"}}`
	SampleNamesVersions    = `{"data":[{"project_id":"d73559a2-7cc8-430f-85fc-7dd7dff71695","name":"bar","version":""},{"project_id":"d73559a2-7cc8-430f-85fc-7dd7dff71695","name":"bar","version":"10.0.1"},{"project_id":"d73559a2-7cc8-430f-85fc-7dd7dff71695","name":"bar","version":"5.1"},{"project_id":"d73559a2-7cc8-430f-85fc-7dd7dff71695","name":"bar","version":"9.a0.0"}],"meta":{"total_count":4,"offset":0,"last_update":"2020-10-22T18:04:59.9235814Z"}}`
)
