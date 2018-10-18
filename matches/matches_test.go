package matches

import (
	"bytes"
	"encoding/json"
	"github.com/iafoosball/matches-service/matches/utils"
	"github.com/iafoosball/matches-service/models"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"
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
	for c := 0; c < 100; c++ {
		_, _ = Collection(matchesColName).CreateDocument(nil, models.Match{
			Key:        "pagedMatchesTest-" + strconv.Itoa(c),
			ID:         "matches/pagedMatchesTest-" + strconv.Itoa(c),
			RatedMatch: true},
		)
	}
	defer func() {
		time.Sleep(1 * time.Second)
		for c := 0; c < 100; c++ {
			if _, err := Collection(matchesColName).RemoveDocument(nil, "pagedMatchesTest-"+strconv.Itoa(c)); err != nil {
				log.Println(err)
			}
		}
	}()
	// Make longer query
	query := "?filter=matches"
	queryUrl := testUrl + "matches" + query
	log.Println(queryUrl)
	if resp, err := http.Get(queryUrl); err != nil || http.StatusOK != resp.StatusCode {
		log.Fatal(err)
	} else {
		// Body should be a paged content
		pagedMatches := utils.PagedMatch{}
		json.NewDecoder(resp.Body).Decode(&pagedMatches)
		log.Printf("%+v\n", pagedMatches)

	}

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
