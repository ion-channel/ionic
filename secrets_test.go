package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestSecrets(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	g.Describe("Secrets", func() {
		var server *bogus.Bogus
		var h, p string
		var client *IonClient
		g.BeforeEach(func() {
			server = bogus.New()
			h, p = server.HostPort()
			client, _ = New(fmt.Sprintf("http://%v:%v", h, p))
		})
		g.AfterEach(func() {
			server.Close()
		})
		g.It("should get secrets", func() {
			server.AddPath("/v1/metadata/getSecrets").
				SetMethods("POST").
				SetPayload([]byte(sampleValidGetSecrets)).
				SetStatus(http.StatusOK)
			secrets, err := client.GetSecrets("text with a secret", "sometoken")
			Expect(err).NotTo(HaveOccurred())
			Expect(secrets).To(HaveLen(1))
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			Expect(hitRecords[0].Header.Get("Authorization")).To(Equal("Bearer sometoken"))
			Expect(secrets[0].Rule).To(Equal("RSA private key"))
			// intentionally blank the second half of the string
			Expect(secrets[0].Match).To(Equal("-----BEGIN RSA ***************"))
		})
	})
}

const sampleValidGetSecrets = `{"data":[{"rule":"RSA private key","match":"-----BEGIN RSA PRIVATE KEY-----","confidence":1}],"meta":{"total_count":0,"offset":0,"last_update":"0001-01-01T00:00:00Z"}}`
