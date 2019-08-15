package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

const (
	testToken = "token"
)

func TestGetDeliveryDestinations(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Delivery Destinations", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should unmarshal list of destinations", func() {
			server.AddPath("/v1/teams/getDeliveryDestinations").
				SetMethods("GET").
				SetPayload([]byte(SampleValidDeliveryDestinations)).
				SetStatus(http.StatusOK)

			deliveries, err := client.GetDeliveryDestinations("7660D469-45DA-4AA3-A421-4F65E9C0CEE9", "token")
			Expect(err).To(BeNil())
			Expect(deliveries).NotTo(BeNil())
			//Expect(alias.ID).To(Equal("334c183d-4d37-4515-84c4-0d0ed0fb8db0"))
			//Expect(alias.Name).To(Equal("name"))
			//Expect(alias.Version).To(Equal("version"))

		})

		//g.It("should return an error for an invalid action", func() {
		//	var ue AnalysisEvent
		//	err := json.Unmarshal([]byte(SampleInvalidAnalysisEvent), &ue)

		//	Expect(err).NotTo(BeNil())
		//	Expect(err.Error()).To(ContainSubstring("invalid analysis event action"))
		//})
	})
}

const (
	SampleValidDeliveryDestinations = `
	{
		"data": [
		  {
			"id": "B3DFA2C7-6DE6-4629-9F19-B493BBE6F2DC",
			"team_id": "7660D469-45DA-4AA3-A421-4F65E9C0CEE9",
			"location": "location1",
			"region": "us-east-1",
			"name": "location1Name",
			"type": "s3"
		  },
		  {
			"id": "0728CC78-868D-4C27-89DE-29C6E6FD11F5",
			"team_id": "7660D469-45DA-4AA3-A421-4F65E9C0CEE9",
			"location": "location2",
			"region": "us-east-1",
			"name": "location2Name",
			"type": "s3",
			"deleted_at": "2019-08-10T20:00:00.840678Z"
		  }
		],
		"meta": {
		  "copyright": "Copyright 2018 Selection Pressure LLC www.selectpress.net",
		  "authors": [
			"Ion Channel Dev Team"
		  ],
		  "version": "v1",
		  "total_count": 1,
		  "offset": 0
		}
	  }
	`

	//SampleInvalidAnalysisEvent = `{"analysis":"fooanalysis", "action":"foo_action"}`
)
