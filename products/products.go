package products

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// Product represents a software product within the system for identification
// across multiple sources
type Product struct {
	ID         int           `json:"id" xml:"id"`
	Name       string        `json:"name" xml:"name"`
	Org        string        `json:"org" xml:"org"`
	Version    string        `json:"version" xml:"version"`
	Up         string        `json:"up" xml:"up"`
	Edition    string        `json:"edition" xml:"edition"`
	Aliases    interface{}   `json:"aliases" xml:"aliases"`
	CreatedAt  time.Time     `json:"created_at" xml:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" xml:"updated_at"`
	Title      string        `json:"title" xml:"title"`
	References []interface{} `json:"references" xml:"references"`
	Part       string        `json:"part" xml:"part"`
	Language   string        `json:"language" xml:"language"`
	ExternalID string        `json:"external_id" xml:"external_id"`
	Sources    []Source      `json:"source" xml:"source"`
}

// Source represents information about where the product data came from
type Source struct {
	ID           int       `json:"id" xml:"id"`
	Name         string    `json:"name" xml:"name"`
	Description  string    `json:"description" xml:"description"`
	CreatedAt    time.Time `json:"created_at" xml:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" xml:"updated_at"`
	Attribution  string    `json:"attribution" xml:"attribution"`
	License      string    `json:"license" xml:"license"`
	CopyrightURL string    `json:"copyright_url" xml:"copyright_url"`
}

// ProductSearchResult represents information about a product as well as
// other info, like Git repository, committer counts, etc
type ProductSearchResult struct {
	Product   Product              `json:"product" xml:"product"`
	Github    Github               `json:"github,omitempty" xml:"github,omitempty"`
	MeanScore float64              `json:"mean_score" xml:"mean_score"`
	Scores    []ProductSearchScore `json:"scores" xml:"scores"`
}

// ProductSearchScore represents the TF;IDF score for a given search result
// and a given search term
type ProductSearchScore struct {
	Term  string  `json:"term" xml:"term"`
	Score float64 `json:"score" xml:"score"`
}

// Github represents information from Github about a given repository
type Github struct {
	URI            string `json:"uri" xml:"uri"`
	CommitterCount uint   `json:"committer_count" xml:"committer_count"`
}

// ProductSearchQuery collects all the various searching options that
// the productSearchEndpoint supports for use in a POST request
type ProductSearchQuery struct {
	SearchType        string   `json:"search_type" xml:"search_type"`
	SearchStrategy    string   `json:"search_strategy" xml:"search_strategy"`
	ProductIdentifier string   `json:"product_identifier" xml:"product_identifier"`
	Version           string   `json:"version" xml:"version"`
	Vendor            string   `json:"vendor" xml:"vendor"`
	Terms             []string `json:"terms" xml:"terms"`
}

// MavenSearchQuery collections all the various searching options that
// the mavenSearchEndpoint supports for use in a POST request
type MavenSearchQuery struct {
	GroupID        string `json:"group_id" xml:"group_id"`
	ArtifactID     string `json:"artifact_id" xml:"artifact_id"`
	SearchType     string `json:"search_type" xml:"search_type"`
	SearchStrategy string `json:"search_strategy", xml:"search_strategy"`
}

// MavenSearchResult represents information about a maven repo
type MavenSearchResult struct {
	GroupID    string        `json:"group_id" xml:"group_id"`
	ArtifactID string        `json:"artifact_id" xml:"artifact_id"`
	Metadata   MavenMetadata `json:"metadata" xml:"metadata"`
}
type MavenMetadata struct {
	XMLName    xml.Name `xml:"metadata"`
	GroupID    string   `xml:"groupId"`
	ArtifactID string   `xml:"artifactId"`
	Version    string   `xml:"version"`
	Versions   []string `xml:"versioning>versions>version"`
}

func (m *MavenMetadata) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Replace(s, "\n", "\\n", -1)
	s = strings.Replace(s, "\\u003c", "<", -1)
	s = strings.Replace(s, "\\u003e", ">", -1)
	var parsed MavenMetadata
	err := xml.Unmarshal([]byte(s), &parsed)
	if err != nil {
		return err
	}
	m.Versions = parsed.Versions
	m.Version = parsed.Version
	m.ArtifactID = parsed.ArtifactID
	m.GroupID = parsed.GroupID
	return nil
}

func (m MavenMetadata) MarshalJSON() ([]byte, error) {
	marshalled, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}
	s := string(marshalled)
	//s = strings.Replace(s, "\\n", "\n", -1)
	s = strings.Replace(s, "\n", "\\n", -1)
	s = strings.Replace(s, "<", "\\u003c", -1)
	s = strings.Replace(s, ">", "\\u003e", -1)
	return []byte(fmt.Sprintf("\"%v\"", s)), nil
}

// IsValid checks some of the constraints on the ProductSearchQuery to
// help the programmer determine if productSearchEndpoint will accept it
func (p *ProductSearchQuery) IsValid() bool {
	if len(p.SearchStrategy) > 0 {
		if p.SearchType == "concatenated" || p.SearchType == "deconcatenated" {
			return true
		}
	}
	return false
}
