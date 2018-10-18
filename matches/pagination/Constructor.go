package pagination

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/iafoosball/matches-service/models"
)

// This file is only for things directly concerning paged resources.

type PagedMatches struct {
	Links   Links          `json:"links"`
	Content []models.Match `json:"content"`
	Page    Page           `json:"page"`
}

type PagedGoals struct {
	Links   Links          `json:"links"`
	Content []*models.Goal `json:"content"`
	Page    Page           `json:"page"`
}

type Links struct {
	First    string `json:"first,omitempty"`
	Previous string `json:"previous,omitempty"`
	Self     string `json:"self,omitempty"`
	Next     string `json:"next,omitempty"`
	Last     string `json:"last,omitempty"`
}
type Page struct {
	Start         int64 `json:"start,omitempty"`
	Size          int64 `json:"size,omitempty"`
	PageNumber    int64 `json:"pageNumber,omitempty"`
	TotalPages    int64 `json:"totalPages,omitempty"`
	TotalElements int64 `json:"totalElements,omitempty"`
}

func (p *PagedMatches) ConstructPage(match []*models.Match, start int64, size int64, totalElements int64, url string) *PagedMatches {
	p.Page.Start = start
	p.Page.Size = size
	p.Page.TotalElements = totalElements
	p.Page.PageNumber = start / size
	p.Page.TotalPages = totalElements
	m := []models.Match{}
	for _, v := range match {
		m = append(m, *v)
	}
	p.Content = m
	return p
}

// [Start: Ensure that system correctly marshalls a struct]
// Validate validates the pagedMatch
func (m *PagedMatches) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PagedMatches) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagedMatches) UnmarshalBinary(b []byte) error {
	var res PagedMatches
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// [End: Ensure that system correctly marshalls a struct]
