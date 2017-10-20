package ionic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestProducts(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Products", func() {
		server := bogus.New()
		server.Start()
		h, p := server.HostPort()
		client, _ := New("", fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get a product", func() {
			server.AddPath("/v1/vulnerability/getProducts").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProduct)).
				SetStatus(http.StatusOK)

			product, err := client.GetProducts("cpe:/a:ruby-lang:ruby:1.8.7")
			Expect(err).To(BeNil())
			Expect(product[0].ID).To(Equal(99182))
			Expect(product[0].Name).To(Equal("ruby"))
		})

		g.It("should get a raw product", func() {
			server.AddPath("/v1/vulnerability/getProducts").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProduct)).
				SetStatus(http.StatusOK)

			raw, err := client.GetRawProducts("cpe:/a:ruby-lang:ruby:1.8.7")
			Expect(err).To(BeNil())
			Expect(raw).To(Equal(json.RawMessage(SampleValidRawProduct)))
		})
	})
}

const (
	SampleValidProduct    = `{"data":[{"id":99182,"name":"ruby","org":"ruby-lang","version":"1.8.7","up":"","edition":"","aliases":null,"created_at":"2017-06-29T03:43:04.145Z","updated_at":"2017-06-29T03:43:04.145Z","title":"ruby-lang Ruby 1.8.7","references":[],"part":"/a","language":"","source_id":1,"external_id":"cpe:/a:ruby-lang:ruby:1.8.7","source":{"id":1,"name":"NVD","description":"National Vulnerability Database","created_at":"2017-06-29T03:41:35.919Z","updated_at":"2017-06-29T03:41:35.919Z","attribution":null,"license":null,"copyright_url":null}}]}`
	SampleValidRawProduct = `["id":99182,"name":"ruby","org":"ruby-lang","version":"1.8.7","up":"","edition":"","aliases":null,"created_at":"2017-06-29T03:43:04.145Z","updated_at":"2017-06-29T03:43:04.145Z","title":"ruby-lang Ruby 1.8.7","references":[],"part":"/a","language":"","source_id":1,"external_id":"cpe:/a:ruby-lang:ruby:1.8.7","source":{"id":1,"name":"NVD","description":"National Vulnerability Database","created_at":"2017-06-29T03:41:35.919Z","updated_at":"2017-06-29T03:41:35.919Z","attribution":null,"license":null,"copyright_url":null}]`
)
