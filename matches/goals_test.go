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
	setup()
	jsonGoal, _ := json.Marshal(models.Goal{
		ID:   "goals/1",
		Key:  "1",
		From: "matches/test1",
		To:   "matches/test1",
	})
	if resp, err := http.Post(testURL+"goals/", "application/json", bytes.NewReader(jsonGoal)); err != nil || http.StatusOK != resp.StatusCode {
		log.Println(resp)
		log.Fatal(err)
	}
}
