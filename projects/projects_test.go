package projects

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

const (
	testToken = "token"
)

func TestProject(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Project Validation", func() {
		var client *http.Client
		var server *bogus.Bogus
		var host, port string

		g.Before(func() {
			server = bogus.New()
			host, port = server.HostPort()

			server.AddPath("/goodurl").
				SetMethods("HEAD")

			server.AddPath("/badurl").
				SetMethods("HEAD").
				SetStatus(http.StatusNotFound)

			client = &http.Client{
				Timeout: time.Second * 1,
			}
		})

		g.It("should return no error if a project is valid", func() {
			server.AddPath("/v1/ruleset/getRuleset").
				SetMethods("HEAD").
				SetStatus(http.StatusOK)

			var p Project
			err := json.Unmarshal([]byte(fmt.Sprintf(sampleValidProject, host, port)), &p)
			b, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
			Expect(err).To(BeNil())

			p.DeployKey = sampleValidKey

			fs, err := p.Validate(client, b, testToken)
			Expect(err).To(BeNil())
			Expect(len(fs)).To(Equal(0))
		})

		g.It("should return no errors for a blank field", func() {
			server.AddPath("/v1/ruleset/getRuleset").
				SetMethods("HEAD").
				SetStatus(http.StatusOK)

			var p Project
			err := json.Unmarshal([]byte(fmt.Sprintf(sampleValidBlankProject, host, port)), &p)
			b, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
			Expect(err).To(BeNil())
			Expect(p.ID).NotTo(BeNil())
			Expect(*p.ID).To(Equal(""))

			fs, err := p.Validate(client, b, testToken)
			Expect(err).To(BeNil())
			Expect(len(fs)).To(Equal(0))
		})

		g.It("should return missing fields as a list and error", func() {
			server.AddPath("/v1/ruleset/getRuleset").
				SetMethods("HEAD").
				SetStatus(http.StatusOK)

			var p Project
			err := json.Unmarshal([]byte(fmt.Sprintf(sampleInvalidProject, host, port)), &p)
			b, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
			Expect(err).To(BeNil())
			Expect(p.Name).To(BeNil())
			Expect(p.Type).To(BeNil())
			Expect(p.Branch).To(BeNil())

			fs, err := p.Validate(client, b, testToken)
			Expect(err).To(Equal(ErrInvalidProject))
			Expect(len(fs)).To(Equal(2))
			Expect(fs["name"]).To(Equal("missing name"))
			Expect(fs["type"]).To(Equal("missing type"))
		})

		g.It("should say a project is invalid if a deploy key is invalid", func() {
			server.AddPath("/v1/ruleset/getRuleset").
				SetMethods("HEAD").
				SetStatus(http.StatusOK)

			var p Project
			err := json.Unmarshal([]byte(fmt.Sprintf(sampleValidProject, host, port)), &p)
			b, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
			Expect(err).To(BeNil())

			p.DeployKey = sampleInvalidKey
			fs, err := p.Validate(client, b, testToken)
			Expect(err).To(Equal(ErrInvalidProject))
			Expect(fs["deploy_key"]).To(Equal("must be a valid ssh key"))

			p.DeployKey = "not valid"
			fs, err = p.Validate(client, b, testToken)
			Expect(err).To(Equal(ErrInvalidProject))
			Expect(fs["deploy_key"]).To(Equal("must be a valid ssh key"))
		})

		g.It("should return string in JSON", func() {
			id := "someid"
			teamID := "someteamiD"
			rulesetID := "someruleset"
			name := "coolproject"
			projectType := "artifact"
			source := "http://%v:%v/goodurl"
			branch := "master"
			desc := "the coolest project around"
			createdAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)
			updatedAt := time.Date(2018, 07, 07, 13, 42, 47, 651387237, time.UTC)

			p := Project{
				ID:             &id,
				TeamID:         &teamID,
				RulesetID:      &rulesetID,
				Name:           &name,
				Type:           &projectType,
				Source:         &source,
				Branch:         &branch,
				Description:    &desc,
				Active:         true,
				ChatChannel:    "#thechan",
				CreatedAt:      createdAt,
				UpdatedAt:      updatedAt,
				DeployKey:      "thekey",
				Monitor:        false,
				POCName:        "youknowit",
				POCEmail:       "you@know.it",
				Username:       "knowit",
				Password:       "supersecret",
				KeyFingerprint: "supersecret",
				Private:        true,
				Aliases:        nil,
				Tags:           nil,
			}
			Expect(fmt.Sprintf("%v", p)).To(Equal(`{"id":"someid","team_id":"someteamiD","ruleset_id":"someruleset","name":"coolproject","type":"artifact","source":"http://%v:%v/goodurl","branch":"master","description":"the coolest project around","active":true,"chat_channel":"#thechan","created_at":"2018-07-07T13:42:47.651387237Z","updated_at":"2018-07-07T13:42:47.651387237Z","deploy_key":"thekey","should_monitor":false,"poc_name":"youknowit","poc_email":"you@know.it","username":"knowit","password":"supersecret","key_fingerprint":"supersecret","private":true,"aliases":null,"tags":null}`))

		})

		g.Describe("Type", func() {
			g.BeforeEach(func() {
				server.AddPath("/v1/ruleset/getRuleset").
					SetMethods("HEAD").
					SetStatus(http.StatusOK)
			})

			g.It("should say a project is valid if the type is valid", func() {
				var p Project
				err := json.Unmarshal([]byte(fmt.Sprintf(sampleValidBlankProject, host, port)), &p)
				b, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
				Expect(err).To(BeNil())

				t := "git"
				p.Type = &t
				fs, err := p.Validate(client, b, testToken)
				Expect(err).To(BeNil())
				Expect(len(fs)).To(Equal(0))

				t = "svn"
				p.Type = &t
				fs, err = p.Validate(client, b, testToken)
				Expect(err).To(BeNil())
				Expect(len(fs)).To(Equal(0))

				t = "artifact"
				p.Type = &t
				fs, err = p.Validate(client, b, testToken)
				Expect(err).To(BeNil())
				Expect(len(fs)).To(Equal(0))

				t = "GiT"
				p.Type = &t
				fs, err = p.Validate(client, b, testToken)
				Expect(err).To(BeNil())
				Expect(len(fs)).To(Equal(0))
			})

			g.It("should say a project is invalid if the type is invalid", func() {
				var p Project
				err := json.Unmarshal([]byte(fmt.Sprintf(sampleValidBlankProject, host, port)), &p)
				b, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
				Expect(err).To(BeNil())

				t := "gahhhbage"
				p.Type = &t
				fs, err := p.Validate(client, b, testToken)
				Expect(err).NotTo(BeNil())
				Expect(len(fs)).To(Equal(1))
			})
		})

		g.Describe("Email", func() {
			g.BeforeEach(func() {
				server.AddPath("/v1/ruleset/getRuleset").
					SetMethods("HEAD").
					SetStatus(http.StatusOK)
			})

			g.It("should say a project is valid if an email is valid", func() {
				var p Project
				err := json.Unmarshal([]byte(fmt.Sprintf(sampleValidBlankProject, host, port)), &p)
				b, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
				Expect(err).To(BeNil())

				p.POCEmail = "dev@ionchannel.io"
				fs, err := p.Validate(client, b, testToken)
				Expect(err).To(BeNil())
				Expect(len(fs)).To(Equal(0))

				p.POCEmail = "dev@howmanyscootersareinthewillamette.science"
				fs, err = p.Validate(client, b, testToken)
				Expect(err).To(BeNil())
				Expect(len(fs)).To(Equal(0))

				p.POCEmail = "me+idontbelieveyouwontspamme@gmail.com"
				fs, err = p.Validate(client, b, testToken)
				Expect(err).To(BeNil())
				Expect(len(fs)).To(Equal(0))

				p.POCEmail = "Acapemail@gmail.com"
				fs, err = p.Validate(client, b, testToken)
				Expect(err).To(BeNil())
				Expect(len(fs)).To(Equal(0))

				p.POCEmail = "  ivegotspaces@gmail.com  "
				fs, err = p.Validate(client, b, testToken)
				Expect(err).To(BeNil())
				Expect(len(fs)).To(Equal(0))
			})

			g.It("should say a project is invalid if an email is invalid", func() {
				var p Project
				err := json.Unmarshal([]byte(fmt.Sprintf(sampleValidBlankProject, host, port)), &p)
				b, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
				Expect(err).To(BeNil())

				p.POCEmail = "notavalidemail"
				fs, err := p.Validate(client, b, testToken)
				Expect(err).To(Equal(ErrInvalidProject))
				Expect(fs["poc_email"]).To(Equal("invalid email supplied"))
			})
		})

		g.Describe("Source", func() {
			g.BeforeEach(func() {
				server.AddPath("/v1/ruleset/getRuleset").
					SetMethods("HEAD").
					SetStatus(http.StatusOK)
			})

			g.It("should permit valid urls", func() {
				var p Project
				err := json.Unmarshal([]byte(fmt.Sprintf(sampleValidBlankProject, host, port)), &p)
				b, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
				Expect(err).To(BeNil())

				us := []string{
					"file:///path/to/repo.git/",
					"file://~/path/to/repo.git/",
					"git://host.xz/path/to/repo.git/",
					"git://host.xz/~user/path/to/repo.git/",
					"git@github.com:foo/bar.git",
					"git@host.xz:/path/to/repo.git/",
					"git@host.xz:path/to/repo.git",
					"git@host.xz:~user/path/to/repo.git/",
					"git@host.xz@~user/path/to/repo.git/",
					"http://host.xz/path/to/repo.git/",
					"http://www.google.com",
					"https://host.xz/path/to/repo.git/",
					"https://www.google.com?y=b",
					"rsync://host.xz/path/to/repo.git/",
					"ssh://host.xz/path/to/repo.git/",
					"ssh://host.xz/path/to/repo.git/",
					"ssh://host.xz/~/path/to/repo.git",
					"ssh://host.xz/~user/path/to/repo.git/",
					"ssh://host.xz:port/path/to/repo.git/",
					"ssh://user@host.xz/path/to/repo.git/",
					"ssh://user@host.xz/path/to/repo.git/",
					"ssh://user@host.xz/~/path/to/repo.git",
					"ssh://user@host.xz/~user/path/to/repo.git/",
					"ssh://user@host.xz:port/path/to/repo.git/",
					"svn+ssh://foo@svn.bar.com/project",
					"svn://svn.code.sf.net/p/regshot/code/trunk",
				}

				for _, val := range us {
					s := val
					t := "git"

					p.Source = &s
					p.Type = &t

					fs, err := p.Validate(client, b, testToken)
					Expect(err).To(BeNil(), fmt.Sprintf("Expected\n%v\nto be nil for repo\n%v\n", err, *p.Source))
					Expect(len(fs)).To(Equal(0))
				}
			})

			g.It("should detect bad urls", func() {
				var p Project
				err := json.Unmarshal([]byte(fmt.Sprintf(sampleValidBlankProject, host, port)), &p)
				b, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
				Expect(err).To(BeNil())

				us := []string{
					"svn://svn.code.sf.net/p/regshot/code/trunk blah",
					"www.google.com",
					"somebody@google.com",
					"mailto:somebody@google.com",
					"www.url-with-querystring.com/?url=has-querystring",
				}

				for _, val := range us {
					s := val
					t := "git"

					p.Source = &s
					p.Type = &t

					fs, err := p.Validate(client, b, testToken)
					Expect(err).NotTo(BeNil())
					Expect(len(fs)).To(Equal(1))
				}
			})
		})

		g.Describe("Ruleset", func() {
			g.It("should return an error if ruleset is not valid", func() {
				server.AddPath("/v1/ruleset/getRuleset").
					SetMethods("HEAD").
					SetStatus(http.StatusNotFound)

				var p Project
				err := json.Unmarshal([]byte(fmt.Sprintf(sampleValidBlankProject, host, port)), &p)
				b, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
				Expect(err).To(BeNil())

				fs, err := p.Validate(client, b, testToken)
				Expect(err).To(Equal(ErrInvalidProject))
				Expect(fs["ruleset_id"]).To(Equal("ruleset id does not match to a valid ruleset"))
			})
		})
	})

	g.Describe("Project Filters", func() {
		g.It("should convert the filter to params", func() {
			a := false
			t := "git"

			pf := Filter{
				Type:   &t,
				Active: &a,
			}

			Expect(pf.Param()).To(Equal("type:git,active:false"))
		})

		g.It("should not include blank filters in the params", func() {
			t := "git"

			pf := Filter{
				Type: &t,
			}

			Expect(pf.Param()).To(Equal("type:git"))
		})

		g.It("should parse a filter from a param", func() {
			a := false
			t := "git"

			pf := Filter{
				Type:   &t,
				Active: &a,
			}

			newPf := ParseParam(pf.Param())
			Expect(newPf).NotTo(BeNil())

			Expect(newPf.Type).NotTo(BeNil())
			Expect(*newPf.Type).To(Equal(t))

			Expect(newPf.Active).NotTo(BeNil())
			Expect(*newPf.Active).To(Equal(a))

			Expect(newPf.TeamID).To(BeNil())
			Expect(newPf.Source).To(BeNil())
		})

		g.It("should convert a filter to a where clause with an identifier", func() {
			a := false
			t := "git"
			ti := "someteamid"

			pf := Filter{
				Type:   &t,
				Active: &a,
				TeamID: &ti,
			}

			query, vals := pf.SQL("p")
			Expect(query).To(Equal("p.team_id=$1 AND p.type=$2 AND p.active=$3"))
			Expect(len(vals)).To(Equal(3))

			teamID, ok := vals[0].(string)
			Expect(ok).To(BeTrue())
			Expect(teamID).To(Equal(ti))

			typeStr, ok := vals[1].(string)
			Expect(ok).To(BeTrue())
			Expect(typeStr).To(Equal(t))

			active, ok := vals[2].(bool)
			Expect(ok).To(BeTrue())
			Expect(active).To(Equal(a))
		})

		g.It("should convert a filter to a where clause without an identifier", func() {
			a := false
			t := "git"
			ti := "someteamid"

			pf := Filter{
				Type:   &t,
				Active: &a,
				TeamID: &ti,
			}

			query, vals := pf.SQL("")
			Expect(query).To(Equal("team_id=$1 AND type=$2 AND active=$3"))
			Expect(len(vals)).To(Equal(3))
		})
	})
}

const (
	sampleValidProject      = `{"id":"someid","team_id":"someteamid","ruleset_id":"someruleset","name":"coolproject","type":"artifact","source":"http://%v:%v/goodurl","branch":"master","description":"the coolest project around","active":true,"chat_channel":"#thechan","created_at":"2018-08-07T13:42:47.258415155-07:00","updated_at":"2018-08-07T13:42:47.258415176-07:00","deploy_key":"thekey","should_monitor":false,"poc_name":"youknowit","poc_email":"you@know.it","username":"knowit","password":"supersecret","key_fingerprint":"supersecret","aliases":null,"tags":null}`
	sampleInvalidProject    = `{"id":"someid","team_id":"someteamid","ruleset_id":"someruleset","source":"http://%v:%v/badurl","description":"the coolest project around","active":true,"chat_channel":"#thechan","created_at":"2018-08-07T13:46:06.529187652-07:00","updated_at":"2018-08-07T13:46:06.529187674-07:00","deploy_key":"","should_monitor":false,"poc_name":"youknowit","poc_email":"you@know.it","username":"knowit","password":"supersecret","key_fingerprint":"supersecret","aliases":null,"tags":null}`
	sampleValidBlankProject = `{"id":"","team_id":"someteamid","ruleset_id":"someruleset","name":"coolproject","type":"artifact","source":"http://%v:%v/goodurl","branch":"master","description":"the coolest project around","active":true,"chat_channel":"#thechan","created_at":"2018-08-07T13:42:47.258415155-07:00","updated_at":"2018-08-07T13:42:47.258415176-07:00","deploy_key":"","should_monitor":false,"poc_name":"youknowit","poc_email":"you@know.it","username":"knowit","password":"supersecret","key_fingerprint":"supersecret","aliases":null,"tags":null}`

	sampleValidKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`
	sampleInvalidKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`
)
