package matches

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/iafoosball/matches-service/models"
	"log"
	"net/http"
	"strconv"
	"testing"
)

// TODO: Hey Aelks what do you think how our test structure should be?
// TODO: Should we have a file for unit tests (which are sort of small integration tests and then this file for integration and benchmarking and possible other tests? Lets have a meeting :D :)

// used by all test classes in package matches
var (
	testHost         string
	testPort         string
	dbHost           string
	dbPort           string
	databaseUser     string
	databasePassword string
	testURL          string
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	flag.StringVar(&testHost, "testHost", "0.0.0.0", "the test host")
	flag.StringVar(&testPort, "testPort", "8000", "the port of the matches service where the test should connect")
	flag.StringVar(&dbHost, "dbhost", "0.0.0.0", "the test host")
	flag.StringVar(&dbPort, "dbport", "8001", "the port of the matches service where the test should connect")
	flag.StringVar(&databaseUser, "dbuser", "root", "the test host")
	flag.StringVar(&databasePassword, "dbpassword", "matches-password", "the port of the matches service where the test should connect")

	flag.Parse()
	testURL = "http://" + testHost + ":" + testPort + "/"
	if i, err := strconv.Atoi(dbPort); err != nil {
		log.Println(err)
	} else {
		InitDatabase(dbHost, i, databaseUser, databasePassword)
	}
}

func setup() {

}

func TestGraph(*testing.T) {
	setup()

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
