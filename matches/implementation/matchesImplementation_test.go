package implementation

import (
	"github.com/iafoosball/matches-service/models"
	"testing"
)

func TestCreateMatch(*testing.T) {
	CreateMatch(models.Match{
		RatedMatch: true,
	})
}

func TestDeleteMatch(*testing.T) {
	//_, s := CreateMatch(models.Match{})

}
