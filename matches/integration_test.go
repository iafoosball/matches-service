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

// TODO: Hey Aelks what do you think how our test structure should be?
// TODO: Should we have a file for unit tests (which are sort of small integration tests and then this file for integration and benchmarking and possible other tests? Lets have a meeting :D :)

// used by all test classes in package matches
const (
	urlTesting = "http://0.0.0.0:8000/"
)

func setup() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func TestGraph(*testing.T) {
	setup()

	jsonObject, _ := json.Marshal(models.Match{
		ID:             "matches/test" + strconv.Itoa(3),
		Key:            "test" + strconv.Itoa(3),
		RatedMatch:     true,
		PositionAttack: true,
	})
	if resp, err := http.Post(urlTesting+"matches/", "application/json", bytes.NewReader(jsonObject)); err != nil || http.StatusOK != resp.StatusCode {
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
		if resp, err := http.Post(urlTesting+"goals/", "application/json", bytes.NewReader(jsonGoal)); err != nil || http.StatusOK != resp.StatusCode {
			log.Println(resp)
			log.Fatal(err)
		}
	}
}
