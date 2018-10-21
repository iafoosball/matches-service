package matches

import (
	"bytes"
	"encoding/json"
	"github.com/iafoosball/matches-service/models"
	"log"
	"net/http"
	"strconv"
	"testing"
)

// Just tests that columns
func TestInitDatabase(*testing.T) {
	user := "root"
	password := "matches-password"
	dbDriver(user, password)
	if col(matchesColName) == nil || col(goalsColName) == nil || graph(goalsToMatchName) == nil || db == nil {
		log.Fatal("Something is not initilized")
	}
}

func TestGraph(*testing.T) {
	jsonObject, _ := json.Marshal(models.Match{
		ID:         "matches/test" + strconv.Itoa(3),
		Key:        "test" + strconv.Itoa(3),
		RatedMatch: true,
	})
	if resp, err := http.Post(testURL+"matches/", "application/json", bytes.NewReader(jsonObject)); err != nil || http.StatusOK != resp.StatusCode {
		log.Println(resp)
		log.Fatal(err)
	}

	for c := 0; c <= 10; c++ {
		jsonGoal, _ := json.Marshal(models.Goal{
			ID:   "goals/" + strconv.Itoa(90+c),
			Key:  strconv.Itoa(90 + c),
			From: "matches/test3",
			To:   "matches/test3",
		})
		if resp, err := http.Post(testURL+"goals/", "application/json", bytes.NewReader(jsonGoal)); err != nil || http.StatusOK != resp.StatusCode {
			log.Println(resp)
			log.Fatal(err)
		}
	}
}
