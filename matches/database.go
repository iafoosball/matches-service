package matches

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/iafoosball/matches-service/models"
	"log"
	"strconv"
	"time"
)

var (
	db                driver.Database
	host              string
	port              int
	dbUser            string
	dbPassword        string
	goalsCol          driver.Collection
	matchesCol        driver.Collection
	goalsToMatchGraph driver.Graph
)

const (
	dbName           = "iaf-matches"
	goalsColName     = "goals"
	matchesColName   = "matches"
	goalsToMatchName = "goalsToMatch"
)

// InitDatabase tries establishes a connection to the db inside a for loop with 10 repetitions.
// If the db is not available it will sleep the time of the counter (c) in seconds.
// When a connection is established it will also open the Collections assiciated with this service.
func InitDatabase(dbHost string, dbPort int, dbUser string, dbPassword string) {
	port = dbPort
	host = dbHost
	dbUser = dbUser
	dbPassword = dbPassword
	log.Println("host: " + dbHost + " port: " + strconv.Itoa(dbPort))
	c := 0
	for db == nil && c <= 10 {
		dbDriver(dbUser, dbPassword)
		c++
		if db == nil {
			log.Println("Database not reachable! Sleep seconds: " + strconv.Itoa(c))
			time.Sleep(time.Duration(c) * 1000 * time.Millisecond)
		}
	}
}

// Authenticate with the arangodb and get the db
func dbDriver(user string, password string) {
	if conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://" + host + ":" + strconv.Itoa(port)},
	}); err == nil {
		if client, err := driver.NewClient(driver.ClientConfig{
			Connection:     conn,
			Authentication: driver.BasicAuthentication(user, password),
		}); err == nil {
			if dbExists, e := client.DatabaseExists(nil, dbName); !dbExists && e == nil {
				if db, err = client.CreateDatabase(nil, dbName, &driver.CreateDatabaseOptions{
					[]driver.CreateDatabaseUserOptions{
						{
							UserName: user,
						},
					},
				}); err != nil {
					log.Fatal(err)
					return
				}
				log.Println("Connected to db: " + dbName)
				Collection(goalsColName)
				Collection(matchesColName)
				Graph(goalsToMatchName)
			} else if e != nil {
				log.Println(e)
			} else {
				db, err = client.Database(nil, dbName)
				Collection(goalsColName)
				Collection(matchesColName)
				Graph(goalsToMatchName)
			}
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}
}

// Open a certain collection. If the collection does not exist, initialize it first.
func Collection(name string) driver.Collection {
	if db == nil {
		InitDatabase(host, port, dbUser, dbPassword)
	} else {
		if name == matchesColName && matchesCol == nil {
			matchesCol = initCollection(name, 2)
			return matchesCol
		} else if name == matchesColName {
			return matchesCol
		}
		if name == goalsColName && goalsCol == nil {
			goalsCol = initCollection(name, 3)
			return goalsCol
		} else if name == goalsColName {
			return goalsCol
		}
	}
	return nil
}

// Initializes a collections
func initCollection(name string, colType int) driver.Collection {
	if exists, err := db.CollectionExists(nil, name); !exists {
		if col, e := db.CreateCollection(nil, name, &driver.CreateCollectionOptions{
			Type: driver.CollectionType(colType),
		}); e != nil {
			return col
		} else if err != nil {
			log.Println(err)
		}
	} else if err != nil {
		log.Println(err)
	} else {

		if col, err := db.Collection(nil, name); err == nil {
			return col
		} else {
			log.Println(err)
		}
	}
	return nil
}

func Graph(name string) driver.Graph {
	if db == nil {
		InitDatabase(host, port, dbUser, dbPassword)
	} else {
		if name == goalsToMatchName && goalsToMatchGraph == nil {
			goalsToMatchGraph = initGraph(name)
			return goalsToMatchGraph
		} else if name == goalsToMatchName {
			return goalsToMatchGraph
		}
	}
	return nil
}

// initGraph is explicitly designed to create a Graph with matches as vertices and goals as edges.
// As we only have one graph for now, I don't think it is necessary to make this func more general,
// but it easy to do so in the future if necessary.
func initGraph(name string) driver.Graph {
	if exists, err := db.GraphExists(nil, name); !exists {
		if col, e := db.CreateGraph(nil, name, &driver.CreateGraphOptions{
			EdgeDefinitions: []driver.EdgeDefinition{{
				Collection: goalsColName,
				To:         []string{matchesColName},
				From:       []string{matchesColName},
			}},
		}); e != nil {
			return col
		} else if err != nil {
			log.Println(err)
		}
	} else if err != nil {
		log.Println(err)
	} else {

		if col, err := db.Graph(nil, name); err == nil {
			return col
		} else {
			log.Println(err)
		}
	}
	return nil
}

// I think this is really bad practice #sometimesGenericsAreNice
// Basically interface is only working with matches. I think it can be rewritten taking multiple slices. lets see
func queryList(query string, matches []*models.Match) []*models.Match {
	if cursor, err := db.Query(nil, query, make(map[string]interface{})); err != nil {
		log.Println(err)
	} else {
		defer cursor.Close()
		for cursor.HasMore() {
			match := &models.Match{}
			if _, err = cursor.ReadDocument(nil, match); err != nil {
				log.Println(err)
			}
			matches = append(matches, match)
		}
	}
	return matches
}
