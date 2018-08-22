package matches

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/users-service/usersImpl"
	"github.com/iafoosball/users-service/restapi/operations"
	"github.com/iafoosball/users-service/models"
)

var dbUsers = usersImpl.DB()
var colUsers = usersImpl.Col("users")

func GetUserByID() func(params operations.GetUsersUserIDParams) middleware.Responder {
	return func(params operations.GetUsersUserIDParams) middleware.Responder {
		//Log the user
		var u = models.User{}
		_, _ = colUsers.ReadDocument(nil, params.UserID, &u)
		return operations.NewGetUsersUserIDOK().WithPayload(&u)
	}
}

func CreateUser() func(params operations.PostUsersParams) middleware.Responder {
	return func(params operations.PostUsersParams) middleware.Responder {
		u := params.Body
		meta, _ := colUsers.CreateDocument(nil, u)
		u.UserID = meta.Key
		colUsers.UpdateDocument(nil, meta.Key, u)
		return operations.NewGetUsersUserIDOK()
	}
}
