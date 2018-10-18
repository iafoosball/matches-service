package matches

import (
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/matches-service/matches/utils"
	"github.com/iafoosball/matches-service/models"
	"github.com/iafoosball/matches-service/restapi/operations"
	"log"
	"strconv"
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

// Still needs implementation of filtering, sorting, tests.
func PagedMatches() func(params operations.GetMatchesParams) middleware.Responder {
	return func(params operations.GetMatchesParams) middleware.Responder {
		var order = "ASC"
		if !*params.ASC {
			order = "DESC"
		}
		// Build arangodb query
		query := "FOR doc IN matches FILTER doc.rated_match == true SORT doc._id " + order + " Limit " + strconv.FormatInt(*params.Start-1, 10) + ", " + strconv.FormatInt(*params.Size, 10) + "RETURN doc"
		matches := queryList(query, &[]models.Match{})
		//log.Printf("%+v\n", matches)
		page, _ := json.Marshal(utils.ConstructPage(matches, *params.Start, *params.Size, 100, ""))
		//log.Printf("%+v\n", page)

		//var pagedMatches utils.PagedMatch
		//json.Unmarshal(page, &pagedMatches)
		//log.Printf("%+v\n", pagedMatches)

		// So on the test side content is an empty array, which should definitely not happen. Thats why there is all the logging stuff. Please leave until pagination is fixed.
		return operations.NewGetMatchesOK().WithPayload(string(page))
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
