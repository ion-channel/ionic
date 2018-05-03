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
