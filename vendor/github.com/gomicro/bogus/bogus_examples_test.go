package bogus_test

import (
	"fmt"
	"github.com/gomicro/bogus"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func ExampleBogus() {
	// This would normally be provided by a normal testing function setup
	var t *testing.T

	server := bogus.New()
	server.AddPath("/foo/bar").
		SetMethods("GET").
		SetPayload([]byte("some return payload")).
		SetStatus(http.StatusOK)
	host, port := server.HostPort()

	resp, err := http.Get(fmt.Sprintf("https://%v:%v", host, port))
	if err != nil {
		t.Errorf("expected nil error, got %v", err.Error())
	}
	defer resp.Body.Close()

	if server.Hits() != 1 {
		t.Errorf("expected server to be hit once, got %v", server.Hits())
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("expected nil error, got: %v", err.Error())
	}

	if string(bodyBytes) != "some return payload" {
		t.Errorf("Expected a different payload, got %v", string(bodyBytes))
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %v", resp.StatusCode)
	}
}

func ExampleBogus_goblinGomega() {
	// This would normally be provided by a normal testing function setup
	var t *testing.T

	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Tests needing a test server", func() {
		var server *bogus.Bogus

		g.BeforeEach(func() {
			server = bogus.New()
		})

		g.It("should connect to a test server", func() {
			server.AddPath("/").
				SetMethods("GET")

			host, port := server.HostPort()

			_, err := http.Get(fmt.Sprintf("https://%v:%v", host, port))
			Expect(err).To(BeNil())
			Expect(server.Hits()).To(Equal(1))
		})
	})

	g.Describe("Tests needing a test server", func() {
		var server *bogus.Bogus

		g.BeforeEach(func() {
			server = bogus.New()
		})

		g.It("should connect to a test server", func() {
			server.AddPath("/foo/bar").
				SetMethods("GET").
				SetPayload([]byte("some return payload")).
				SetStatus(http.StatusOK)
			host, port := server.HostPort()

			resp, err := http.Get(fmt.Sprintf("https://%v:%v", host, port))
			Expect(err).To(BeNil())
			defer resp.Body.Close()

			Expect(server.Hits()).To(Equal(1))

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			Expect(err).To(BeNil())
			Expect(string(bodyBytes)).To(Equal("some return payload"))
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})
}
