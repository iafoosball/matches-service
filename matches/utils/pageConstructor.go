package utils

import "github.com/iafoosball/matches-service/models"

// I need to learn interfaces...
type PagedMatch struct {
	Links   Links          `json:"links"`
	Content []models.Match `json:"content"`
	Page    Page           `json:"page"`
}

type Links struct {
	First    string `json:"first"`
	Previous string `json:"previous"`
	Self     string `json:"self"`
	Next     string `json:"next"`
	Last     string `json:"last"`
}
type Page struct {
	Start         int64 `json:"start"`
	Size          int64 `json:"size"`
	PageNumber    int64 `json:"pageNumber"`
	TotalPages    int64 `json:"totalPages"`
	TotalElements int64 `json:"totalElements"`
}

func ConstructPage(match *[]models.Match, start int64, size int64, totalElements int64, url string) *PagedMatch {
	p := &PagedMatch{}
	p.Page.Start = start
	p.Page.Size = size
	p.Page.TotalElements = totalElements
	p.Page.PageNumber = start / size
	p.Page.TotalPages = totalElements
	p.Content = *match
	return p
}
