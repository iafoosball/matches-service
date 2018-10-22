package matches

import (
	"bytes"
	"encoding/json"
	"github.com/iafoosball/matches-service/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestCreateGoal(*testing.T) {
	createMatches(10)
	jsonGoal, _ := json.Marshal(models.Goal{
		ID:   "goals/1",
		Key:  "1",
		From: "matches/test8",
		To:   "matches/test8",
	})
	if resp, err := http.Post(testAddr+"goals/", "application/json", bytes.NewReader(jsonGoal)); err != nil || http.StatusOK != resp.StatusCode {
		log.Println(resp)
		log.Fatal(err)
	}
}

func TestPagedGoals(t *testing.T) {
	totalMatches := 10
	totalGoals := 10
	size := 5
	start := 1
	createMatches(totalMatches)
	createGoals(totalGoals)
	var pagedGoals models.PagedGoals
	query := testAddr + "goals/?filter=match_id=='test-1'&ASC=false&size=" + strconv.Itoa(size) + "&start=" + strconv.Itoa(start)
	for query != "" {
		log.Println(query)
		if resp, err = http.Get(query); err != nil || http.StatusOK != resp.StatusCode {
			log.Fatal(err)
		}
		if body, err = ioutil.ReadAll(resp.Body); err != nil {
			log.Fatal(err)
		}
		err = json.NewDecoder(strings.NewReader(string(body))).Decode(&pagedGoals)
		for _, m := range pagedGoals.Links {
			if m.Rel == "Next" && m.Href != "" {
				query = m.Href
				m.Href = ""
				break
			} else {
				query = ""
			}
		}
		start += size
	}
	if err != nil {
		log.Fatal(err)
	}
	content := pagedGoals.Content
	if len(content) != size {
		log.Println(len(content))
		log.Fatal("not correct number of items")
	}

}

func createGoals(amount int) {
	for match := 1; match < 3; match++ {
		for c := 0; c < amount; c++ {
			_, err = col(goalsColName).CreateDocument(nil, models.Goal{
				//Key:     "test-" + strconv.Itoa(c),
				//ID:      "goals/test-" + strconv.Itoa(c*match),
				From:    "matches/test-" + strconv.Itoa(match),
				To:      "matches/test-" + strconv.Itoa(match),
				MatchID: "test-" + strconv.Itoa(match),
				Side:    "blue",
			},
			)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
