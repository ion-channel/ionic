package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestScanner(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Scanner", func() {
		server := bogus.New()
    server.Start()
		h, p := server.HostPort()
		client, _ := New("", fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get an analysis", func() {
			server.AddPath("/v1/scanner/analyzeProject").
				SetMethods("POST").
				SetPayload([]byte(SampleValidAnalysisStatus)).
				SetStatus(http.StatusOK)

			analysisStatus, err := client.AnalyzeProject("ateamid", "aprojectid")
			Expect(err).To(BeNil())
			Expect(analysisStatus.ID).To(Equal("3cc5e53b-5298-47cd-ba6c-313bf47ae4a8"))
		})
	})
}

const (
	SampleValidAnalysisStatus = `{"branch": "master","build_number": null,"created_at": "2017-09-25T21:43:11.069Z","id": "3cc5e53b-5298-47cd-ba6c-313bf47ae4a8","message": "Request for analysis 3cc5e53b-5298-47cd-ba6c-313bf47ae4a8 on Amazon Web Services SDK has been accepted.","project_id": "93e2f31e-b579-4490-864d-7c630ac49720","status": "accepted","team_id": "9e0d2de6-7612-4349-a520-8fc0af2a23cf","updated_at": "2017-09-25T21:43:11.069Z"}`
)
