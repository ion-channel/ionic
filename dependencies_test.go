package ionic

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/franela/goblin"
	"github.com/gomicro/bogus"
	"github.com/ion-channel/ionic/dependencies"
	. "github.com/onsi/gomega"
)

func TestDependencies(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Dependencies", func() {
		var server *bogus.Bogus
		var h, p string
		var client *IonClient

		g.BeforeEach(func() {
			server = bogus.New()
			h, p = server.HostPort()
			client, _ = New(fmt.Sprintf("http://%v:%v", h, p))
		})

		g.It("should get the latest version of a dependency", func() {
			server.AddPath("/v1/dependency/getLatestVersionForDependency").
				SetMethods("GET").
				SetPayload([]byte(sampleLatestVersionResponse)).
				SetStatus(http.StatusOK)
			dep, err := client.GetLatestVersionForDependency("bundler", RubyEcosystem, "atoken")

			Expect(err).To(BeNil())
			Expect(dep.Version).To(Equal("1.16.3"))

			hrs := server.HitRecords()
			Expect(len(hrs)).To(Equal(1))
		})

		g.It("should get the latest version of a dependency", func() {
			server.AddPath("/v1/dependency/getVersionsForDependency").
				SetMethods("GET").
				SetPayload([]byte(sampleLatestVersionsResponse)).
				SetStatus(http.StatusOK)
			deps, err := client.GetVersionsForDependency("bundler", RubyEcosystem, "atoken")

			Expect(err).To(BeNil())
			Expect(deps[0].Version).To(Equal("1.16.3"))

			hrs := server.HitRecords()
			Expect(len(hrs)).To(Equal(1))
		})

		g.It("should resolve dependencies from a definition file", func() {
			server.AddPath("/v1/dependency/resolveDependenciesInFile").
				SetMethods("POST").
				SetPayload([]byte(sampleResolutionResponse)).
				SetStatus(http.StatusOK)

			tf, err := ioutil.TempFile("", "test")
			if err != nil {
				log.Fatal(err)
			}

			defer os.Remove(tf.Name()) // clean up

			tf.Write([]byte(samplePomSnippet))
			tf.Close()

			o := dependencies.DependencyResolutionRequest{
				File:      tf.Name(),
				Flatten:   true,
				Ecosystem: "maven",
			}

			deps, err := client.ResolveDependenciesInFile(o, "atoken")

			Expect(err).To(BeNil())
			Expect(deps.Dependencies[0].Version).To(Equal("1.16.3"))

			hrs := server.HitRecords()
			Expect(len(hrs)).To(Equal(1))
			Expect(string(hrs[0].Body)).To(ContainSubstring(samplePomSnippet))
		})

		g.It("should support search for dependencies", func() {
			server.AddPath("/v1/dependency/search").
				SetMethods("GET").
				SetPayload([]byte(sampleSearchResponse)).
				SetStatus(http.StatusOK)

			tf, err := ioutil.TempFile("", "test")
			if err != nil {
				log.Fatal(err)
			}

			defer os.Remove(tf.Name()) // clean up

			tf.Write([]byte(samplePomSnippet))
			tf.Close()

			deps, err := client.SearchDependencies("some org", "atoken")

			Expect(err).To(BeNil())
			Expect(deps[0].Version).To(Equal("1.16.3"))
			Expect(deps[0].Name).To(Equal("invaluable"))
			Expect(deps[0].Org).To(Equal("nefarious"))

			hrs := server.HitRecords()
			Expect(len(hrs)).To(Equal(1))
		})

		g.It("should get dependency versions", func() {
			server.AddPath("/v1/dependency/getVersions").
				SetMethods("GET").
				SetPayload([]byte(sampleDependencyVersions)).
				SetStatus(http.StatusOK)

			deps, err := client.GetDependencyVersions("hypercord", "npm", "", "atoken")

			Expect(err).To(BeNil())
			Expect(deps[0].Version).To(Equal("1.4.1"))
			Expect(deps[0].Name).To(Equal("hypercored"))
			Expect(deps[0].CreatedAt).NotTo(BeNil())
			Expect(len(deps)).To(Equal(4))

			hrs := server.HitRecords()
			Expect(len(hrs)).To(Equal(1))
		})
	})
}

const (
	sampleLatestVersionResponse  = `{"data":{"version":"1.16.3"}}`
	sampleLatestVersionsResponse = `{"data":["1.16.3","1.16.2"]}`
	sampleResolutionResponse     = `{"data":{"dependencies":[{"version":"1.16.3"}]}}`
	sampleSearchResponse         = `{"data":[{"name":"invaluable", "version":"1.16.3", "org": "nefarious"}]}`
	sampleDependencyVersions     = `{"data":[{"name":"hypercored","version":"1.4.1","latest_version":"","org":"","type":"","package":"","scope":"","requirement":"","dependencies":null,"confidence":0,"created_at":"2017-09-27T09:54:05.643Z","updated_at":"2021-02-04T23:32:37.078164Z"},{"name":"hypercored","version":"1.4.0","latest_version":"","org":"","type":"","package":"","scope":"","requirement":"","dependencies":null,"confidence":0,"created_at":"2017-09-27T07:34:04.857Z","updated_at":"2021-02-04T23:32:37.078164Z"},{"name":"hypercored","version":"1.3.0","latest_version":"","org":"","type":"","package":"","scope":"","requirement":"","dependencies":null,"confidence":0,"created_at":"2017-07-11T22:31:42.531Z","updated_at":"2021-02-04T23:32:37.078164Z"},{"name":"hypercored","version":"1.2.2","latest_version":"","org":"","type":"","package":"","scope":"","requirement":"","dependencies":null,"confidence":0,"created_at":"2017-06-13T14:14:28.372Z","updated_at":"2021-02-04T23:32:37.078164Z"}]}`

	samplePomSnippet = `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
<dependencies>
        <dependency>
            <groupId>org.apache.cxf</groupId>
            <artifactId>cxf-rt-frontend-jaxrs</artifactId>
            <version>2.7.3</version>
        </dependency>
    </dependencies>
</project>`
)
