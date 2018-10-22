package matches

import (
	"bytes"
	"encoding/json"
	"github.com/iafoosball/matches-service/models"
	"log"
	"net/http"
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
