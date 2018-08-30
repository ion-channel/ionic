package products

import (
	"testing"

	"encoding/json"
	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestProductsStructs(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	g.Describe("marshaling and unmarshaling structs in Products", func() {
		g.It("should unmarshal a MavenSearchResult", func() {
			var msr MavenSearchResult
			err := json.Unmarshal([]byte(sampleMulchResponse), &msr)
			Expect(err).NotTo(HaveOccurred())
			Expect(msr).NotTo(BeNil())
			Expect(msr.GroupID).To(Equal("testgroup001"))
			Expect(msr.ArtifactID).To(Equal("testartifact001"))
			Expect(msr.Metadata).NotTo(BeNil())
			Expect(msr.Metadata.GroupID).To(Equal("testgroup001"))
			Expect(msr.Metadata.ArtifactID).To(Equal("testartifact001"))
			Expect(msr.Metadata.Version).To(Equal("1.0"))
			Expect(msr.Metadata.Versions).To(HaveLen(20))
			Expect(msr.Metadata.Versions[12]).To(Equal("1.2-rc1"))
		})
		g.It("should marshal a MavenSearchResult", func() {
			msr := MavenSearchResult{
				GroupID:    "testgroup001",
				ArtifactID: "testartifact001",
				Metadata: MavenMetadata{
					GroupID:    "testgroup001",
					ArtifactID: "testartifact001",
					Version:    "1.0",
					Versions: []string{
						"ver01",
						"ver02",
						"ver03",
					},
				},
			}
			marshalled, err := json.Marshal(msr)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(marshalled)).To(ContainSubstring(`"group_id":"testgroup001"`))
			Expect(string(marshalled)).To(ContainSubstring(`"artifact_id":"testartifact001"`))
			Expect(string(marshalled)).NotTo(ContainSubstring(`metadata":{`))
			Expect(string(marshalled)).To(ContainSubstring("\\u003cartifactId\\u003etestartifact001\\u003c/artifactId\\u003e\\n"))
		})
	})
}

var (
	sampleMulchResponse = `{"id":1,"group_id":"testgroup001","artifact_id":"testartifact001","metadata":"\u003cmetadata\u003e\n  \u003cgroupId\u003etestgroup001\u003c/groupId\u003e\n  \u003cartifactId\u003etestartifact001\u003c/artifactId\u003e\n  \u003cversion\u003e1.0\u003c/version\u003e\n  \u003cversioning\u003e\n    \u003cversions\u003e\n      \u003cversion\u003e1.0\u003c/version\u003e\n      \u003cversion\u003e1.0-m4\u003c/version\u003e\n      \u003cversion\u003e1.0-rc1\u003c/version\u003e\n      \u003cversion\u003e1.1\u003c/version\u003e\n      \u003cversion\u003e1.1-rc1\u003c/version\u003e\n      \u003cversion\u003e1.1-rc2\u003c/version\u003e\n      \u003cversion\u003e1.1.1\u003c/version\u003e\n      \u003cversion\u003e1.1.2\u003c/version\u003e\n      \u003cversion\u003e1.1.3\u003c/version\u003e\n      \u003cversion\u003e1.1.4\u003c/version\u003e\n      \u003cversion\u003e1.1.5\u003c/version\u003e\n      \u003cversion\u003e1.2\u003c/version\u003e\n      \u003cversion\u003e1.2-rc1\u003c/version\u003e\n      \u003cversion\u003e1.2-rc2\u003c/version\u003e\n      \u003cversion\u003e1.2.1\u003c/version\u003e\n      \u003cversion\u003e1.2.2\u003c/version\u003e\n      \u003cversion\u003e1.2.3\u003c/version\u003e\n      \u003cversion\u003e1.2.4\u003c/version\u003e\n      \u003cversion\u003e1.2.5\u003c/version\u003e\n      \u003cversion\u003e1.2.6\u003c/version\u003e\n    \u003c/versions\u003e\n  \u003c/versioning\u003e\n\u003c/metadata\u003e","created_at":"2018-08-29T19:15:33.045Z","updated_at":"2018-08-29T19:15:33.045Z"}`
)
