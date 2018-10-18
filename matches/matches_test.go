package matches

import (
	"bytes"
	"encoding/json"
	"github.com/iafoosball/matches-service/matches/pagination"
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
	setup()
	jsonObject, _ := json.Marshal(models.Match{
		ID:         "matches/test" + strconv.Itoa(1),
		Key:        "test" + strconv.Itoa(1),
		RatedMatch: true,
	})
	if resp, err := http.Post(testUrl+"matches/", "application/json", bytes.NewReader(jsonObject)); err != nil || http.StatusOK != resp.StatusCode {
		log.Println(resp)
		log.Fatal(err)
	}
}

func TestPagedMatches(t *testing.T) {
	amount := 30
	createMatches(amount)
	defer removeMatches(amount)
	// Make longer query
	query := "?filter=&ASC=false&size=30"
	queryUrl := testUrl + "matches" + query
	log.Println(queryUrl)
	if resp, err = http.Get(queryUrl); err != nil || http.StatusOK != resp.StatusCode {
		log.Fatal(err)
	}
	//Always close body to avoid memory leak
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Fatal(err)
	}
	var pagedMatches pagination.PagedMatches
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&pagedMatches)
	if err != nil {
		log.Fatal(err)
	}
	content := pagedMatches.Content
	if len(content) != amount {
		log.Println(len(content))
		log.Fatal("not correct number of items")
	}
	log.Printf("%+v\n", pagedMatches)

}

func BenchmarkCreateMatch(b *testing.B) {
	for c := 0; c < b.N; c++ {
		jsonObject, _ := json.Marshal(models.Match{
			ID:         "matches/test" + strconv.Itoa(c),
			Key:        "test" + strconv.Itoa(c),
			RatedMatch: true,
		})
		if resp, err := http.Post(testUrl+"matches/", "application/json", bytes.NewReader(jsonObject)); err != nil || http.StatusOK != resp.StatusCode {
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
		_, _ = Collection(matchesColName).CreateDocument(nil, models.Match{
			Key:        "pagedMatchesTest-" + strconv.Itoa(c),
			ID:         "matches/pagedMatchesTest-" + strconv.Itoa(c),
			EndTime:    strconv.Itoa(c),
			RatedMatch: true},
		)
	}
}

func removeMatches(amount int) {
	for c := 0; c < amount; c++ {
		if _, err := Collection(matchesColName).RemoveDocument(nil, "pagedMatchesTest-"+strconv.Itoa(c)); err != nil {
			log.Println(err)
		}
	}
}
