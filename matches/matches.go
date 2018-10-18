package matches

import (
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/matches-service/matches/pagination"
	"github.com/iafoosball/matches-service/models"
	"github.com/iafoosball/matches-service/restapi/operations"
	"log"
	"strconv"
	"strings"
)

var (
	err error
)

// Has test
func CreateMatch() func(params operations.PostMatchesParams) middleware.Responder {
	return func(params operations.PostMatchesParams) middleware.Responder {
		if _, err := Collection(matchesColName).CreateDocument(nil, &params.Body); err != nil {
			request, _ := json.Marshal(params.Body)
			log.Println(string(request))
			log.Println(err)
		}
		return operations.NewPostMatchesOK()
	}
}

// Still needs implementation of tests and links.
func PagedMatches() func(params operations.GetMatchesParams) middleware.Responder {
	return func(params operations.GetMatchesParams) middleware.Responder {
		// Build arangodb query
		sort := *params.Sort
		sort = "Sort doc." + strings.Replace(sort, ",", ", doc.", -1)
		log.Println(sort)
		filter := *params.Filter
		if filter != "" {
			filter = "Filter doc." + strings.Replace(filter, ",", ", doc.", -1)
		}
		query := "FOR doc IN matches " + filter + " " + sort + " " + *params.Order + " Limit " + strconv.FormatInt(*params.Start-1, 10) + ", " + strconv.FormatInt(*params.Size, 10) + " RETURN doc"
		matches := queryMatches(query, []*models.Match{})
		query = "For doc In matches " + filter + " COLLECT WITH COUNT INTO length RETURN length"
		total := int64(queryInt(query))

		page := pagination.PagedMatches{}
		page.ConstructPage(matches, addr, *params.Filter, *params.Sort, *params.Order, *params.Start, *params.Size, total)

		return operations.NewGetMatchesOK().WithPayload(page)
	}
}

// TODO: implement + test
func DeleteMatch(matchId string) *operations.PostMatchesOK {
	matchesCol.RemoveDocument(nil, matchId)
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

//func GetUserByID() func(params operations.GetUsersUserIDParams) middleware.Responder {
//	return func(params operations.GetUsersUserIDParams) middleware.Responder {
//		//Log the user
//		var u = models.User{}
//		_, _ = colMatches.ReadDocument(nil, params.UserID, &u)
//		return operations.NewGetUsersUserIDOK().WithPayload(&u)
//	}
//}
//
//func CreateUser() func(params operations.PostUsersParams) middleware.Responder {
//	return func(params operations.PostUsersParams) middleware.Responder {
//		u := params.Body
//		meta, _ := colMatches.CreateDocument(nil, u)
//		u.UserID = meta.Key
//		colMatches.UpdateDocument(nil, meta.Key, u)
//		return operations.NewGetUsersUserIDOK()
//	}
//}

// Gets a query as string and a slice and returns that slice with references to the found matches.
func queryMatches(query string, matches []*models.Match) []*models.Match {
	if cursor, err := db.Query(nil, query, make(map[string]interface{})); err != nil {
		log.Println(err)
	} else {
		defer cursor.Close()
		for cursor.HasMore() {
			match := &models.Match{}
			if _, err = cursor.ReadDocument(nil, match); err != nil {
				log.Println(err)
			}
			matches = append(matches, match)
		}
	}
	return matches
}
