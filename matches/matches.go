package matches

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/matches-service/matches/paged"
	"github.com/iafoosball/matches-service/models"
	"github.com/iafoosball/matches-service/restapi/operations"
	"log"
)

var (
	err error
)

// CreateMatch creates a match from the input. This is normally done through the Livematch and a full Match is expected.
// Has test
func CreateMatch() func(params operations.PostMatchesParams) middleware.Responder {
	return func(params operations.PostMatchesParams) middleware.Responder {
		meta, err := col(matchesColName).CreateDocument(nil, &params.Body)
		handleErr(err)
		return operations.NewPostMatchesOK().WithPayload(meta)
	}
}

// PagedMatches returns a page with matches as content
// has test (maybe rework)
func PagedMatches() func(params operations.GetMatchesParams) middleware.Responder {
	return func(params operations.GetMatchesParams) middleware.Responder {
		matches := queryMatches(paged.MatchesQuery(params))
		query := "For doc In matches " + paged.BuildFilter(*params.Filter) + " COLLECT WITH COUNT INTO length RETURN length"
		total := int64(queryInt(query))
		url := addr + "/matches/?filter=" + *params.Filter + "&sort=" + *params.Sort + "&order=" + *params.Order
		a := paged.Matches(matches, url, *params.Start, *params.Size, total)
		return operations.NewGetMatchesOK().WithPayload(a)
	}
}

// DeleteMatch needs implement + test
func DeleteMatch(matchID string) *operations.PostMatchesOK {
	matchesCol.RemoveDocument(nil, matchID)
	return operations.NewPostMatchesOK()
}

func playingWithStuffGetAllGoalsFromMatch() {
	cursor, _ := db.Query(nil, "FOR e IN goals FILTER e._from == 'matches/test1' RETURN e", nil)
	for c := 0; cursor.HasMore(); c++ {
		var goal models.Goal
		infot, _ := cursor.ReadDocument(nil, &goal)
		log.Println(goal)
		log.Println(infot)
	}
	log.Println()
}

// Gets a query as string and a slice and returns that slice with references to the found matches.
func queryMatches(query string) []*models.Match {
	matches := []*models.Match{}
	if cursor, err := db.Query(nil, query, make(map[string]interface{})); err != nil {
		log.Println(err)
	} else {
		defer cursor.Close()
		for cursor.HasMore() {
			match := &models.Match{}
			_, err := cursor.ReadDocument(nil, match)
			handleErr(err)
			matches = append(matches, match)
		}
	}
	return matches
}

// Generic method to handle errors
func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
