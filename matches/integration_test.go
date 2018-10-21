package matches

import (
	"flag"
	"log"
	"strconv"
)

// used by all test classes in package matches. Do we need this class.
var (
	testHost         string
	testPort         string
	dbHost           string
	dbPort           string
	databaseUser     string
	databasePassword string
	testURL          string
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	flag.StringVar(&testHost, "testhost", "0.0.0.0", "the test host")
	flag.StringVar(&testPort, "testport", "8000", "the port of the matches service where the test should connect")
	flag.StringVar(&dbHost, "dbhost", "0.0.0.0", "the test host")
	flag.StringVar(&dbPort, "dbport", "8001", "the port of the matches service where the test should connect")
	flag.StringVar(&databaseUser, "dbuser", "root", "the test host")
	flag.StringVar(&databasePassword, "dbpassword", "matches-password", "the port of the matches service where the test should connect")

	flag.Parse()
	testURL = "http://" + testHost + ":" + testPort + "/"
	if i, err := strconv.Atoi(dbPort); err != nil {
		log.Println(err)
	} else {
		InitDatabase(dbHost, i, databaseUser, databasePassword)
	}

	col(matchesColName).Truncate(nil)
	col(goalsColName).Truncate(nil)
}
