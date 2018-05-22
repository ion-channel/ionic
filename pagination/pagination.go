package pagination

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// AllItems is a convenience for requesting all items of a given entity
var AllItems = &Pagination{Offset: 0, Limit: -1}

// Pagination represents the necessary elements for a paginated request
type Pagination struct {
	Offset int
	Limit  int
}

// New takes an offset and limit.  It returns a newly created Pagination object
// and prevents the offset and limit from being set to illegal values.
func New(offset, limit int) *Pagination {
	p := &Pagination{
		Offset: offset,
		Limit:  limit,
	}

	if p.Offset < 0 {
		p.Offset = 0
	}

	if p.Limit < 1 {
		p.Limit = -1
	}

	return p
}

// ParseFromRequest parses pagination params from an http request and returns
// the pagination object and an error if the pagination is not found
func ParseFromRequest(req *http.Request) *Pagination {
	oStr := req.Header.Get("Offset")
	lStr := req.Header.Get("Limit")

	if oStr == "" {
		oStr = "0"
	}

	if lStr == "" {
		lStr = "10"
	}

	o, err := strconv.Atoi(oStr)
	if err != nil {
		o = 0
	}

	l, err := strconv.Atoi(lStr)
	if err != nil {
		l = 10
	}

	return &Pagination{Offset: o, Limit: l}
}

// AddParams appends the pagination params to the provided set of URL values
func (p *Pagination) AddParams(params *url.Values) {
	params.Set("offset", strconv.Itoa(p.Offset))
	params.Set("limit", strconv.Itoa(p.Limit))
}

// Down increments the offset down by the limit.  It will not increment the
// offset past 0.
func (p *Pagination) Down() {
	if p.Limit > 0 {
		p.Offset -= p.Limit
		if p.Offset < 0 {
			p.Offset = 0
		}
	}
}

// SQL returns a valid string representation of the pagination object
func (p *Pagination) SQL() string {
	strs := []string{}

	if p.Offset > 0 {
		strs = append(strs, fmt.Sprintf("OFFSET %d", p.Offset))
	}

	switch {
	case p.Limit > 0:
		strs = append(strs, fmt.Sprintf("LIMIT %d", p.Limit))
	case p.Limit <= 0:
	}

	return strings.Join(strs, " ")
}

// Up increments the offset up by the limit.
func (p *Pagination) Up() {
	if p.Limit > 0 {
		p.Offset += p.Limit
	}
}
