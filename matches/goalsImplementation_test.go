package matches

import (
	"fmt"
	"github.com/iafoosball/matches-service/models"
	"log"
	"testing"
)

func TestCreateGoal(t *testing.T) {
	_, metaMatches := CreateMatch(models.Match{})
	s := metaMatches.ID.String()
	fmt.Println(metaMatches.ID)
	_, metaGoals := CreateGoal(models.Goal{
		Datetime: "123123",
		From:     &s,
		To:       s,
	})
	fmt.Print("The goal metaGoals data")
	fmt.Println(metaGoals)
	if created, _ := goalsCol.DocumentExists(nil, metaGoals.Key); created == false {
		log.Fatal("Document was not created")
	}
	fmt.Println(goalsCol.RemoveDocument(nil, metaGoals.Key))
	fmt.Println(matchesCol.RemoveDocument(nil, metaMatches.Key))
}

func BenchmarkCreateGoal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, metaMatches := CreateMatch(models.Match{})
		s := metaMatches.ID.String()
		//fmt.Println(metaMatches.ID)
		_, metaGoals := CreateGoal(models.Goal{
			Datetime: "123123",
			From:     &s,
			To:       s,
		})
		//fmt.Print("The goal metaGoals data")
		//fmt.Println(metaGoals)
		if created, _ := goalsCol.DocumentExists(nil, metaGoals.Key); created == false {
			log.Fatal("Document was not created")
		}
		//fmt.Println(goalsCol.RemoveDocument(nil, metaGoals.Key))
		//fmt.Println(matchesCol.RemoveDocument(nil, metaMatches.Key))
	}
	fmt.Println(b)

}
