package paged

import (
	"github.com/iafoosball/matches-service/models"
	"strconv"
)

// Matches returns a pagination objects.
func Matches(m []*models.Match, url string, start int64, size int64, totalElements int64) *models.PagedMatches {
	pM := models.PagedMatches{}
	pM.Page = page(start, size, totalElements)
	pM.Content = m
	pM.Links = links(url, start, size, totalElements)
	return &pM
}

func page(start int64, size int64, totalElements int64) *models.Page {
	p := &models.Page{}
	p.Start = start
	p.Size = size
	p.TotalItems = totalElements
	p.CurrentPage = start / size
	p.TotalPages = totalElements
	return p
}

func links(addr string, start int64, size int64, totalItems int64) models.Links {
	return models.Links{
		&models.LinksItems0{
			Rel:  "first",
			Href: buildLink(addr, 1, size),
		},
		&models.LinksItems0{
			Rel:  "previous",
			Href: previous(addr, start, size),
		},
		&models.LinksItems0{
			Rel:  "self",
			Href: buildLink(addr, start, size),
		},
		&models.LinksItems0{
			Rel:  "next",
			Href: next(addr, start, size, totalItems),
		},
		&models.LinksItems0{
			Rel:  "last",
			Href: last(addr, start, size, totalItems),
		},
	}
}

// api can either be
func previous(url string, start int64, size int64) string {
	if start > size {
		return buildLink(url, start-size, size)
	}
	return ""
}

func next(addr string, start int64, size int64, total int64) string {
	if start+size < total {
		return buildLink(addr, start+size, size)
	}
	return ""
}
func last(addr string, start int64, size int64, total int64) string {
	if total-size > 1 {
		return buildLink(addr, total-size, size)
	}
	return buildLink(addr, 1, size)
}

func buildLink(addr string, start int64, size int64) string {
	link := addr + "&start=" + strconv.FormatInt(start, 10) + "&size=" + strconv.FormatInt(size, 10)
	return link
}
