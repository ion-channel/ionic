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
		var server *bogus.Bogus
		var h, p string
		var client *IonClient
		g.BeforeEach(func(){
			server = bogus.New()
			server.Start()
			h, p = server.HostPort()
			client, _ = New(fmt.Sprintf("http://%v:%v", h, p))
		})
		g.AfterEach(func(){
			server.Close()
		})

		g.It("should get a product", func() {
			server.AddPath("/v1/vulnerability/getProducts").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProduct)).
				SetStatus(http.StatusOK)

			product, err := client.GetProducts("cpe:/a:oracle:jdk:1.6.0:update_71", "someapikey")
			Expect(err).To(BeNil())
			Expect(product[0].Sources[0].Name).To(Equal("NVD"))
			Expect(product[0].ID).To(Equal(84647))
			Expect(product[0].Name).To(Equal("jdk"))
		})

		g.It("should get a raw product", func() {
			server.AddPath("/v1/vulnerability/getProducts").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProduct)).
				SetStatus(http.StatusOK)

			raw, err := client.GetRawProducts("cpe:/a:oracle:jdk:1.6.0:update_71", "fookey")
			Expect(err).To(BeNil())
			Expect(raw).To(Equal(json.RawMessage(SampleValidRawProduct)))
		})

		g.It("should search for a product", func() {
			server.AddPath("/v1/product/search").
				SetMethods("GET").
				SetPayload([]byte(sampleBunsenSearchResponse)).
				SetStatus(http.StatusOK)
			products, err := client.GetProductSearch("less", "mahVersion", "someapikey")
			Expect(err).To(BeNil())
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			hitRecord := hitRecords[0]
			Expect(hitRecord.Query.Get("user_query")).To(Equal("less"))
			Expect(hitRecord.Query.Get("version")).To(Equal("mahVersion"))
			Expect(products).To(HaveLen(5))
			Expect(products[0].ID).To(Equal(39862))
		})
		g.It("should omit version from product search when it is not given", func(){
			server.AddPath("/v1/product/search").
				SetMethods("GET").
				SetPayload([]byte(sampleBunsenSearchResponse)).
				SetStatus(http.StatusOK)
			products, err := client.GetProductSearch("less", "", "someapikey")
			Expect(err).To(BeNil())
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			hitRecord := hitRecords[0]
			Expect(hitRecord.Query.Get("user_query")).To(Equal("less"))
			Expect(hitRecord.Query.Get("version")).To(Equal(""))
			Expect(products).To(HaveLen(5))
			Expect(products[0].ID).To(Equal(39862))
		})
	})
}

const (
	SampleValidProduct         = `{"data":[{"id":84647,"name":"jdk","org":"oracle","version":"1.6.0","up":"update_71","edition":"","aliases":null,"created_at":"2017-02-13T20:02:42.600Z","updated_at":"2017-02-13T20:02:42.600Z","title":"Oracle JDK 1.6.0 Update 71","references":[{"April 2014 CPU":"http://www.oracle.com/technetwork/topics/security/cpuapr2014-1972952.html"}],"part":"/a","language":"","external_id":"cpe:/a:oracle:jdk:1.6.0:update_71","source":[{"id":1,"name":"NVD","description":"National Vulnerability Database","created_at":"2017-02-09T20:18:35.385Z","updated_at":"2017-02-13T20:12:05.342Z","attribution":"Copyright © 1999–2017, The MITRE Corporation. CVE and the CVE logo are registered trademarks and CVE-Compatible is a trademark of The MITRE Corporation.","license":"Submissions: For all materials you submit to the Common Vulnerabilities and Exposures (CVE®)","copyright_url":"http://cve.mitre.org/about/termsofuse.html"}]}]}`
	SampleValidRawProduct      = `[{"id":84647,"name":"jdk","org":"oracle","version":"1.6.0","up":"update_71","edition":"","aliases":null,"created_at":"2017-02-13T20:02:42.600Z","updated_at":"2017-02-13T20:02:42.600Z","title":"Oracle JDK 1.6.0 Update 71","references":[{"April 2014 CPU":"http://www.oracle.com/technetwork/topics/security/cpuapr2014-1972952.html"}],"part":"/a","language":"","external_id":"cpe:/a:oracle:jdk:1.6.0:update_71","source":[{"id":1,"name":"NVD","description":"National Vulnerability Database","created_at":"2017-02-09T20:18:35.385Z","updated_at":"2017-02-13T20:12:05.342Z","attribution":"Copyright © 1999–2017, The MITRE Corporation. CVE and the CVE logo are registered trademarks and CVE-Compatible is a trademark of The MITRE Corporation.","license":"Submissions: For all materials you submit to the Common Vulnerabilities and Exposures (CVE®)","copyright_url":"http://cve.mitre.org/about/termsofuse.html"}]}]`
	sampleBunsenSearchResponse = `{"data":[{"id":39862,"name":"less","org":"gnu","version":"-","up":"","edition":"","aliases":null,"created_at":"2017-02-13T20:02:36.794Z","updated_at":"2017-02-13T20:02:36.794Z","title":"GNU less","references":[],"part":"/a","language":"","external_id":"cpe:/a:gnu:less:-"},{"id":39863,"name":"less","org":"gnu","version":"358","up":"","edition":"","aliases":null,"created_at":"2017-02-13T20:02:36.794Z","updated_at":"2017-02-13T20:02:36.794Z","title":"GNU less 358","references":[],"part":"/a","language":"","external_id":"cpe:/a:gnu:less:358"},{"id":39864,"name":"less","org":"gnu","version":"381","up":"","edition":"","aliases":null,"created_at":"2017-02-13T20:02:36.794Z","updated_at":"2017-02-13T20:02:36.794Z","title":"GNU less 381","references":[],"part":"/a","language":"","external_id":"cpe:/a:gnu:less:381"},{"id":39865,"name":"less","org":"gnu","version":"382","up":"","edition":"","aliases":null,"created_at":"2017-02-13T20:02:36.794Z","updated_at":"2017-02-13T20:02:36.794Z","title":"GNU less 382","references":[],"part":"/a","language":"","external_id":"cpe:/a:gnu:less:382"},{"id":39866,"name":"less","org":"gnu","version":"471","up":"","edition":"","aliases":null,"created_at":"2017-02-13T20:02:36.794Z","updated_at":"2017-02-13T20:02:36.794Z","title":"GNU less 471","references":[{"Vendor Website":"http://www.gnu.org/software/less/"}],"part":"/a","language":"","external_id":"cpe:/a:gnu:less:471"}],"meta":{"copyright":"Copyright 2018 Selection Pressure LLC www.selectpress.net","authors":["Ion Channel Dev Team"],"version":"v1","last_update":"2018-05-03T16:27:42.409Z","total_count":5,"limit":10,"offset":0},"links":{"self":"https://api.ionchannel.io/v1/product/search?user_query=less"}}`
)
