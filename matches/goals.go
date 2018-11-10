package matches

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/matches-service/matches/paged"
	"github.com/iafoosball/matches-service/models"
	"github.com/iafoosball/matches-service/restapi/operations"
	"log"
	"strconv"
	"strings"
)

// CreateGoal from input
// Has test
func CreateGoal() func(params operations.PostGoalsParams) middleware.Responder {
	return func(params operations.PostGoalsParams) middleware.Responder {
		_, err := col(goalsColName).CreateDocument(nil, &params.Body)
		handleErr(err)
		return operations.NewPostGoalsOK()
	}
}

func BatchCreateGoals() func(params operations.PostGoalsBatchCreateParams) middleware.Responder {
	return func(params operations.PostGoalsBatchCreateParams) middleware.Responder {
		for _, g := range params.Body {
			_, err := col(goalsColName).CreateDocument(nil, g)
			handleErr(err)
		}
		return operations.NewPostGoalsBatchCreateOK()
	}
}

// PagedGoals returns a page with matches as content
// Has test
func PagedGoals() func(params operations.GetGoalsParams) middleware.Responder {
	return func(params operations.GetGoalsParams) middleware.Responder {
		sort := *params.Sort
		sort = "Sort doc." + strings.Replace(sort, ",", ", doc.", -1)
		filter := *params.Filter
		if filter != "" {
			filter = "Filter doc." + strings.Replace(filter, ",", ", doc.", -1)
		}
		// is this possible in one query, getting number of total items and the selected items?
		query := "FOR doc IN goals " + filter + " " + sort + " " + *params.Order + " Limit " + strconv.FormatInt(*params.Start-1, 10) + ", " + strconv.FormatInt(*params.Size, 10) + " RETURN doc"
		goals := queryGoals(query)
		query = "For doc In goals " + filter + " COLLECT WITH COUNT INTO length RETURN length"
		total := int64(queryInt(query))

		for _, goal := range goals {
			log.Printf("%+v\n", goal)
		}

		url := addr + "/goals/?filter=" + *params.Filter + "&sort=" + *params.Sort + "&order=" + *params.Order
		a := paged.Goals(goals, url, *params.Start, *params.Size, total)
		return operations.NewGetGoalsOK().WithPayload(a)
	}
}

// Gets a query as string and a slice and returns that slice with references to the found matches.
func queryGoals(query string) []*models.Goal {
	goals := []*models.Goal{}
	if cursor, err := db.Query(nil, query, make(map[string]interface{})); err != nil {
		log.Println(err)
	} else {
		defer cursor.Close()
		for cursor.HasMore() {
			g := &models.Goal{}
			if _, err = cursor.ReadDocument(nil, g); err != nil {
				log.Println(err)
			}
			goals = append(goals, g)
		}
	}
	return goals
}
