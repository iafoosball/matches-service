package matchesApi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/matches-service/matches/matchesImpl"
	"github.com/iafoosball/matches-service/restapi/operations"
)

func CreateGoal() func(params operations.PostGoalsParams) middleware.Responder {
	return func(params operations.PostGoalsParams) middleware.Responder {
		c, _ := matchesImpl.CreateGoal(*params.Body)
		return c
	}
}
