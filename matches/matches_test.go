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

var (
	resp *http.Response
	body []byte
)

func TestCreateMatch(*testing.T) {
	jsonObject, _ := json.Marshal(models.Match{
		ID:         "matches/test" + strconv.Itoa(1),
		Key:        "test" + strconv.Itoa(1),
		RatedMatch: true,
	})
	if resp, err := http.Post(testURL+"matches/", "application/json", bytes.NewReader(jsonObject)); err != nil || http.StatusOK != resp.StatusCode {
		log.Println(resp)
		log.Fatal(err)
	}
}

func TestPagedMatches(t *testing.T) {
	totalMatches := 1000
	size := 10
	start := 981
	createMatches(totalMatches)
	var pagedMatches models.PagedMatches
	query := testURL + "matches/?filter=&ASC=false&size=" + strconv.Itoa(size) + "&start=" + strconv.Itoa(start)
	for query != "" {
		log.Println(query)
		if resp, err = http.Get(query); err != nil || http.StatusOK != resp.StatusCode {
			log.Fatal(err)
		}
		if body, err = ioutil.ReadAll(resp.Body); err != nil {
			log.Fatal(err)
		}
		err = json.NewDecoder(strings.NewReader(string(body))).Decode(&pagedMatches)
		for _, m := range pagedMatches.Links {
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
	content := pagedMatches.Content
	if len(content) != size {
		log.Println(len(content))
		log.Fatal("not correct number of items")
	}

}

func BenchmarkCreateMatch(b *testing.B) {
	for c := 0; c < b.N; c++ {
		jsonObject, _ := json.Marshal(models.Match{
			ID:         "matches/test" + strconv.Itoa(c),
			Key:        "test" + strconv.Itoa(c),
			RatedMatch: true,
		})
		if resp, err := http.Post(testURL+"matches/", "application/json", bytes.NewReader(jsonObject)); err != nil || http.StatusOK != resp.StatusCode {
			log.Println(resp)
			log.Fatal(err)
		}
	}
}

func TestDeleteMatch(*testing.T) {
	//_, s := CreateMatch(models.Match{})

}

func createMatches(amount int) {
	for c := 0; c < amount; c++ {
		_, _ = col(matchesColName).CreateDocument(nil, models.Match{
			Key:        "test-" + strconv.Itoa(c),
			ID:         "matches/test-" + strconv.Itoa(c),
			EndTime:    strconv.Itoa(c),
			RatedMatch: true},
		)
	}
}
