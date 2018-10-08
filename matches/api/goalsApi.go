package api

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/matches-service/matches/implementation"
	"github.com/iafoosball/matches-service/restapi/operations"
)

func CreateGoal() func(params operations.PostGoalsParams) middleware.Responder {
	return func(params operations.PostGoalsParams) middleware.Responder {
		c, _ := implementation.CreateGoal(*params.Body)
		return c
	}
}
