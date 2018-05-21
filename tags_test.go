package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestTags(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Tags", func() {
		var server *bogus.Bogus
		var h, p string
		var client *IonClient

		g.BeforeEach(func() {
			server = bogus.New()
			h, p = server.HostPort()
			client, _ = New(fmt.Sprintf("http://%v:%v", h, p))
		})

		g.It("should get a tag", func() {
			server.AddPath("/v1/tag/getTag").
				SetMethods("GET").
				SetPayload([]byte(SampleValidTag)).
				SetStatus(http.StatusOK)

			tag, err := client.GetTag("27A06D70-A8AA-4532-ABBF-C83ADF49F855", "A3EB1676-C91A-4FCF-AE1D-9F887C0D4B66", "atoken")
			Expect(err).To(BeNil())
			Expect(tag.ID).To(Equal("27A06D70-A8AA-4532-ABBF-C83ADF49F855"))
			Expect(tag.Name).To(Equal("Jenkins"))
		})

		g.It("should get tags", func() {
			server.AddPath("/v1/tag/getTags").
				SetMethods("GET").
				SetPayload([]byte(SampleGetTagsResp)).
				SetStatus(http.StatusOK)

			tags, err := client.GetTags("646fa3e5-e274-4884-aef2-1d47f029c289", "atoken")
			Expect(err).To(BeNil())
			Expect(tags[0].Name).To(Equal("Jdk"))
			Expect(tags[1].Name).To(Equal("Java"))
			Expect(tags).To(HaveLen(2))
		})

		g.It("should get raw tags", func() {
			server.AddPath("/v1/tag/getTags").
				SetMethods("GET").
				SetPayload([]byte(SampleGetTagsResp)).
				SetStatus(http.StatusOK)

			tags, err := client.GetRawTags("646fa3e5-e274-4884-aef2-1d47f029c289", "atoken")

			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			hitRecord := hitRecords[0]
			Expect(hitRecord.Header.Get("Authorization")).To(Equal("Bearer atoken"))

			Expect(err).To(BeNil())
			Expect(tags).NotTo(BeNil())
			Expect(string(tags)).To(ContainSubstring("e1e7f72c-e32a-42e5-adef-955c9fd0f5f3"))
			Expect(string(tags)).To(ContainSubstring("646fa3e5-e274-4884-aef2-1d47f029c289"))
			Expect(string(tags)).To(ContainSubstring("Java"))
		})

		g.It("should create a tag", func() {
			server.AddPath("/v1/tag/createTag").
				SetMethods("POST").
				SetPayload([]byte(SampleValidTag)).
				SetStatus(http.StatusCreated)

			tag, err := client.CreateTag("A3EB1676-C91A-4FCF-AE1D-9F887C0D4B66", "Jenkins", "wild description", "atoken")

			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			hitRecord := hitRecords[0]
			Expect(hitRecord.Header.Get("Authorization")).To(Equal("Bearer atoken"))
			Expect(err).To(BeNil())
			Expect(tag.ID).To(Equal("27A06D70-A8AA-4532-ABBF-C83ADF49F855"))
			Expect(tag.TeamID).To(Equal("A3EB1676-C91A-4FCF-AE1D-9F887C0D4B66"))
			Expect(tag.Name).To(Equal("Jenkins"))
			Expect(tag.Description).To(Equal("CI"))
		})
	})
}

const (
	SampleValidTag    = `{"data":{"id":"27A06D70-A8AA-4532-ABBF-C83ADF49F855","team_id":"A3EB1676-C91A-4FCF-AE1D-9F887C0D4B66","name":"Jenkins","description":"CI","created_at":"2017-07-10T18:39:53.914Z","updated_at":"2017-07-10T18:39:53.914Z"}}`
	SampleGetTagsResp = `{"data":[{"id":"e1e7f72c-e32a-42e5-adef-955c9fd0f5f3","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","name":"Jdk","description":"","created_at":"2018-03-22T12:58:41.544924Z","updated_at":"2018-03-22T12:58:41.544924Z"},{"id":"6dada8a5-a593-4d39-909b-30402215043c","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","name":"Java","description":"","created_at":"2018-03-22T12:58:41.544924Z","updated_at":"2018-03-22T12:58:41.544924Z"}],"meta":{"copyright":"Copyright 2018 Selection Pressure LLC www.selectpress.net","authors":["Ion Channel Dev Team"],"version":"v1"}}`
)
