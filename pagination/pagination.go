package pagination

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	// DefaultOffset is the standard offset chosen when nothing is provided
	DefaultOffset = 0
	// DefaultLimit is the standard limit chosen when nothing is provided
	DefaultLimit = 10
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

// ParseFromRequest parses pagination params from an http request's query params
// or from the headers. The URL query params are favored over the header values
// if both are provided. It will defer to defaults if the pagination params are
// not found.
func ParseFromRequest(req *http.Request) *Pagination {
	var oStr, lStr string

	if req.URL != nil {
		oStr = req.URL.Query().Get("offset")
		lStr = req.URL.Query().Get("limit")
	}

	if oStr == "" && lStr == "" && req.Header != nil {
		oStr = req.Header.Get("Offset")
		lStr = req.Header.Get("Limit")
	}

	if oStr == "" && lStr == "" {
		return New(DefaultOffset, DefaultLimit)
	}

	if oStr == "" {
		oStr = strconv.Itoa(DefaultOffset)
	}

	o, err := strconv.Atoi(oStr)
	if err != nil {
		o = DefaultOffset
	}

	if lStr == "" {
		oStr = strconv.Itoa(DefaultLimit)
	}

	l, err := strconv.Atoi(lStr)
	if err != nil {
		l = DefaultLimit
	}

	return New(o, l)
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
