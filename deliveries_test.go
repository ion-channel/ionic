package ionic

import (
	"fmt"
	"net/http"
	"testing"

	goblin "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	"github.com/ion-channel/ionic/deliveries"
	. "github.com/onsi/gomega"
)

const (
	testToken = "token"
)

func TestGetDeliveryDestinations(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Get Delivery Destinations", func() {
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
			Expect(deliveries[0].ID).To(Equal("B3DFA2C7-6DE6-4629-9F19-B493BBE6F2DC"))
			Expect(deliveries[1].Name).To(Equal("location2Name"))
		})

		g.It("should return an error for an invalid action", func() {
			server.AddPath("/v1/teams/getDeliveryDestinations").
				SetMethods("GET").
				SetPayload([]byte(SampleInvalidDeliveryDestinations)).
				SetStatus(http.StatusOK)

			deliveries, err := client.GetDeliveryDestinations("7660D469-45DA-4AA3-A421-4F65E9C0CEE9", "token")
			Expect(err).NotTo(BeNil())
			Expect(deliveries).To(BeNil())
		})
	})
}

func TestDeleteDeliveryDestination(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Delete Delivery Destinations", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get no error when sucessful", func() {
			server.AddPath("/v1/teams/deleteDeliveryDestination").
				SetMethods("DELETE").
				SetStatus(http.StatusNoContent)

			err := client.DeleteDeliveryDestination("7660D469-45DA-4AA3-A421-4F65E9C0CEE9", "token")
			if err != nil {
				fmt.Printf("\nHERE:\n%v\n", err.Error())
			}
			Expect(err).To(BeNil())
		})

		g.It("should get err when API returns bad request", func() {
			server.AddPath("/v1/teams/deleteDeliveryDestination").
				SetMethods("DELETE").
				SetStatus(http.StatusBadRequest)

			err := client.DeleteDeliveryDestination("7660D469-45DA-4AA3-A421-4F65E9C0CEE9", "token")
			Expect(err).NotTo(BeNil())
		})

		g.It("should get err when API returns not found", func() {
			server.AddPath("/v1/teams/deleteDeliveryDestination").
				SetMethods("DELETE").
				SetStatus(http.StatusNotFound)

			err := client.DeleteDeliveryDestination("7660D469-45DA-4AA3-A421-4F65E9C0CEE9", "token")
			Expect(err).NotTo(BeNil())
		})

		g.It("should get err when API returns internal server error", func() {
			server.AddPath("/v1/teams/deleteDeliveryDestination").
				SetMethods("DELETE").
				SetStatus(http.StatusInternalServerError)

			err := client.DeleteDeliveryDestination("7660D469-45DA-4AA3-A421-4F65E9C0CEE9", "token")
			Expect(err).NotTo(BeNil())
		})
	})
}

func TestCreateDeliveryDestinations(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Delivery Destinations", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should add a delivery destination", func() {
			server.AddPath("/v1/teams/createDeliveryDestination").
				SetMethods("POST").
				SetPayload([]byte(SampleValidDeliveryDestination)).
				SetStatus(http.StatusOK)

			d := &deliveries.CreateDestination{
				Destination: deliveries.Destination{
					TeamID:   "7660D469-45DA-4AA3-A421-4F65E9C0CEE9",
					Location: "location1",
					Region:   "us-east-1",
					Name:     "location1Name",
					DestType: "s3",
				},
				AccessKey: "",
				SecretKey: "",
			}

			dest, err := client.CreateDeliveryDestinations(d, "token")
			Expect(err).To(BeNil())
			Expect(dest.ID).To(Equal("334c183d-4d37-4515-84c4-0d0ed0fb8db0"))
			Expect(dest.Name).To(Equal("location1Name"))
			Expect(dest.Region).To(Equal("us-east-1"))
		})
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
		]
	  }
	`

	SampleValidDeliveryDestination = `
		{  
			"data":{  
			"id":"334c183d-4d37-4515-84c4-0d0ed0fb8db0",
			"team_id":"7660D469-45DA-4AA3-A421-4F65E9C0CEE9",
			"location":"location1",
			"region":"us-east-1",
			"name":"location1Name",
			"type":"s3"
			}
		}
	`

	SampleInvalidDeliveryDestinations = `{"analysis":"fooanalysis", "action":"foo_action"}`
)
