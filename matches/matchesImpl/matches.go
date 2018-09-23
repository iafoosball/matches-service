package matchesImpl

import (
	"github.com/arangodb/go-driver"
	"github.com/iafoosball/matches-service/models"
	"github.com/iafoosball/matches-service/restapi/operations"
)

func CreateMatch(match models.Match) (*operations.PostMatchesOK, driver.DocumentMeta) {
	meta, _ := matchesCol.CreateDocument(nil, &match)
	return operations.NewPostMatchesOK(), meta
}

func DeleteMatch(matchId string) *operations.PostMatchesOK {
	matchesCol.RemoveDocument(nil, matchId)
	return operations.NewPostMatchesOK()
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
