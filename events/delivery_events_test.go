package events

import (
	"encoding/json"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestDeliveryEvents(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Delivery Events", func() {
		g.It("should unmarshal a delivery event action", func() {
			var ue DeliveryEvent
			err := json.Unmarshal([]byte(SampleValidDeliveryEvent), &ue)

			Expect(err).To(BeNil())
			Expect(ue.Action).To(Equal(DeliveryEventAction("delivery_failed")))
		})

		g.It("should return an error for an invalid action", func() {
			var ue DeliveryEvent
			err := json.Unmarshal([]byte(SampleInvalidDeliveryEvent), &ue)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("invalid delivery event action"))
		})

		g.It("should return an error for an invalid action on marshalling", func() {
			ue := DeliveryEvent{
				Action: "not valid",
			}
			_, err := json.Marshal(ue)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("invalid delivery event action"))
		})

		g.It("should correctly marshal the action", func() {
			ue := DeliveryEvent{
				Action: "delivery_finished",
			}
			bytes, err := json.Marshal(ue)

			Expect(err).To(BeNil())
			Expect(string(bytes)).To(Equal(SampleMarshalledDeliveryEvent))
		})
	})
}

const (
	SampleMarshalledDeliveryEvent = `{"action":"delivery_finished","analysis":"","project_id":"","team_id":"","timestamp":"","url":"","filename":"","type":""}`
	SampleValidDeliveryEvent      = `{"delivery":"foodelivery", "action":"delivery_failed"}`
	SampleInvalidDeliveryEvent    = `{"delivery":"foodelivery", "action":"foo_action"}`
)
