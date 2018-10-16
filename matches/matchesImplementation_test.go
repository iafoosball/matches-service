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

func TestCreateMatch(*testing.T) {
	setup()
	jsonObject, _ := json.Marshal(models.Match{
		ID:         "goals/test" + strconv.Itoa(1),
		Key:        "test" + strconv.Itoa(1),
		RatedMatch: true,
	})
	if resp, err := http.Post(testUrl+"matches/", "application/json", bytes.NewReader(jsonObject)); err != nil || http.StatusOK != resp.StatusCode {
		log.Println(resp)
		log.Fatal(err)
	}
}

func BenchmarkCreateMatch(b *testing.B) {
	for c := 0; c < b.N; c++ {
		jsonObject, _ := json.Marshal(models.Match{
			ID:         "goals/test" + strconv.Itoa(c),
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
