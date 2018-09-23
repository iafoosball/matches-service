package matchesApi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/matches-service/matches/matchesImpl"
	"github.com/iafoosball/matches-service/restapi/operations"
)

func CreateMatch() func(params operations.PostMatchesParams) middleware.Responder {
	return func(params operations.PostMatchesParams) middleware.Responder {
		c, _ := matchesImpl.CreateMatch(*params.Body)
		return c
	}
}
