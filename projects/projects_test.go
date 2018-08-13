package projects

import (
	"encoding/json"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestAnalysis(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Projects", func() {
		g.It("should return no error if a project is valid", func() {
			var p Project
			err := json.Unmarshal([]byte(sampleValidProject), &p)
			Expect(err).To(BeNil())

			fs, err := p.Validate()
			Expect(err).To(BeNil())
			Expect(len(fs)).To(Equal(0))
		})

		g.It("should return no errors for a blank field", func() {
			var p Project
			err := json.Unmarshal([]byte(sampleValidBlankProject), &p)
			Expect(err).To(BeNil())
			Expect(p.ID).NotTo(BeNil())
			Expect(*p.ID).To(Equal(""))

			fs, err := p.Validate()
			Expect(err).To(BeNil())
			Expect(len(fs)).To(Equal(0))
		})

		g.It("should return missing fields as a list and error", func() {
			var p Project
			err := json.Unmarshal([]byte(sampleInvalidProject), &p)
			Expect(err).To(BeNil())
			Expect(p.Name).To(BeNil())
			Expect(p.Type).To(BeNil())
			Expect(p.Branch).To(BeNil())

			fs, err := p.Validate()
			Expect(err).To(Equal(ErrInvalidProject))
			Expect(len(fs)).To(Equal(3))
			Expect(fs[0]).To(Equal("name"))
			Expect(fs[1]).To(Equal("type"))
			Expect(fs[2]).To(Equal("branch"))
		})

		g.It("should say a project is invalid if an email is invalid", func() {
			var p Project
			err := json.Unmarshal([]byte(sampleValidProject), &p)
			Expect(err).To(BeNil())

			p.POCEmail = "dev@ionchannel.io"
			fs, err := p.Validate()
			Expect(err).To(BeNil())
			Expect(fs).To(BeNil())

			p.POCEmail = "dev@howmanyscootersareinthewillamette.science"
			fs, err = p.Validate()
			Expect(err).To(BeNil())
			Expect(fs).To(BeNil())

			p.POCEmail = "notavalidemail"
			fs, err = p.Validate()
			Expect(err).To(Equal(ErrInvalidProjectEmail))
			Expect(fs[0]).To(Equal("poc_email"))
		})
	})
}

const (
	sampleValidProject      = `{"id":"someid","team_id":"someteamid","ruleset_id":"someruleset","name":"coolproject","type":"git","source":"github","branch":"master","description":"the coolest project around","active":true,"chat_channel":"#thechan","created_at":"2018-08-07T13:42:47.258415155-07:00","updated_at":"2018-08-07T13:42:47.258415176-07:00","deploy_key":"thekey","should_monitor":false,"poc_name":"youknowit","poc_email":"you@know.it","username":"knowit","password":"supersecret","key_fingerprint":"supersecret","aliases":null,"tags":null}`
	sampleInvalidProject    = `{"id":"someid","team_id":"someteamid","ruleset_id":"someruleset","source":"github","description":"the coolest project around","active":true,"chat_channel":"#thechan","created_at":"2018-08-07T13:46:06.529187652-07:00","updated_at":"2018-08-07T13:46:06.529187674-07:00","deploy_key":"thekey","should_monitor":false,"poc_name":"youknowit","poc_email":"you@know.it","username":"knowit","password":"supersecret","key_fingerprint":"supersecret","aliases":null,"tags":null}`
	sampleValidBlankProject = `{"id":"","team_id":"someteamid","ruleset_id":"someruleset","name":"coolproject","type":"git","source":"github","branch":"master","description":"the coolest project around","active":true,"chat_channel":"#thechan","created_at":"2018-08-07T13:42:47.258415155-07:00","updated_at":"2018-08-07T13:42:47.258415176-07:00","deploy_key":"thekey","should_monitor":false,"poc_name":"youknowit","poc_email":"you@know.it","username":"knowit","password":"supersecret","key_fingerprint":"supersecret","aliases":null,"tags":null}`
)
