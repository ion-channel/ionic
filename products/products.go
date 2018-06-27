package products

import (
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
