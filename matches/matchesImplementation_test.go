package matches

import (
	"bytes"
	"encoding/json"
	"github.com/iafoosball/matches-service/models"
	"log"
	"net/http"
	"testing"
)

func TestCreateMatch(*testing.T) {
	jsonObject, _ := json.Marshal(models.Match{
		RatedMatch:     true,
		PositionAttack: true,
	})
	if resp, err := http.Post("http://0.0.0.0:9000/matches/", "application/jsonObject", bytes.NewReader(jsonObject)); err != nil || http.StatusOK != resp.StatusCode {
		log.Fatal(err)
	}
}

func BenchmarkCreateMatch(b *testing.B) {
	for c := 0; c < b.N; c++ {
		jsonObject, _ := json.Marshal(models.Match{
			RatedMatch:     true,
			PositionAttack: true,
		})
		if resp, err := http.Post("http://0.0.0.0:9000/matches/", "application/jsonObject", bytes.NewReader(jsonObject)); err != nil || http.StatusOK != resp.StatusCode {
			log.Fatal(err)
		}
	}
}

func TestDeleteMatch(*testing.T) {
	//_, s := CreateMatch(models.Match{})

}
