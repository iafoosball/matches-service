package paged

import (
	"github.com/iafoosball/matches-service/models"
	"github.com/iafoosball/matches-service/restapi/operations"
	"log"
	"strconv"
	"strings"
)

// Matches returns a pagination objects  with content of array of matches.
func Matches(m []*models.Match, url string, start int64, size int64, totalElements int64) *models.PagedMatches {
	pM := models.PagedMatches{}
	pM.Page = page(start, size, totalElements)
	pM.Content = m
	pM.Links = links(url, start, size, totalElements)
	return &pM
}

func MatchesQuery(params operations.GetMatchesParams) string {
	sort := *params.Sort
	sort = "Sort doc." + strings.Replace(sort, ",", ", doc.", -1)
	filter := BuildFilter(*params.Filter)
	// is this possible in one query, getting number of total items and the selected items?
	log.Println("FOR doc IN matches " + filter + " " + sort + " " + *params.Order + " Limit " + strconv.FormatInt(*params.Start-1, 10) + ", " + strconv.FormatInt(*params.Size, 10) + " RETURN doc")
	return "FOR doc IN matches " + filter + " " + sort + " " + *params.Order + " Limit " + strconv.FormatInt(*params.Start-1, 10) + ", " + strconv.FormatInt(*params.Size, 10) + " RETURN doc"

}
