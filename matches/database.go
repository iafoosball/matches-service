package matches

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
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
	// This is the server address. It should probably not be here, but where else?
	addr   string
	er     error
	exists bool
)

const (
	dbName           = "iaf-matches"
	goalsColName     = "goals"
	matchesColName   = "matches"
	goalsToMatchName = "goalsToMatch"
)

//Addr to set address of the server
func Addr(address string) {
	addr = "http://" + address
}

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
				col(goalsColName)
				col(matchesColName)
				graph(goalsToMatchName)
			} else if e != nil {
				log.Println(e)
			} else {
				db, err = client.Database(nil, dbName)
				col(goalsColName)
				col(matchesColName)
				graph(goalsToMatchName)
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

// Open a collection. If the collection does not exist, initialize it first.
func col(name string) driver.Collection {
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
	var col driver.Collection
	var exists bool
	if exists, er = db.CollectionExists(nil, name); !exists && er == nil {
		if col, er = db.CreateCollection(nil, name, &driver.CreateCollectionOptions{
			Type: driver.CollectionType(colType),
		}); er != nil {
			return col
		}
	} else if exists && er == nil {

		if col, er = db.Collection(nil, name); er == nil {
			return col
		}
	}
	if er != nil {
		log.Println(er)
	}
	return nil
}

func graph(name string) driver.Graph {
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

// initGraph is explicitly designed to create a graph with matches as vertices and goals as edges.
// As we only have one graph for now, I don't think it is necessary to make this func more general,
// but it easy to do so in the future if necessary.
func initGraph(name string) driver.Graph {
	var graph driver.Graph
	if exists, er = db.GraphExists(nil, name); !exists && er == nil {
		if graph, er = db.CreateGraph(nil, name, &driver.CreateGraphOptions{
			EdgeDefinitions: []driver.EdgeDefinition{{
				Collection: goalsColName,
				To:         []string{matchesColName},
				From:       []string{matchesColName},
			}},
		}); er != nil {
			return graph
		}
	} else if er == nil {
		if graph, er = db.Graph(nil, name); er == nil {
			return graph
		}
	}
	if er != nil {
		log.Println(er)
	}
	return nil
}

// Get an integer from a query.
func queryInt(query string) int {
	if cursor, err := db.Query(nil, query, make(map[string]interface{})); err != nil {
		log.Println(err)
	} else {
		defer cursor.Close()
		var i int
		for cursor.HasMore() {
			if _, err = cursor.ReadDocument(nil, &i); err != nil {
				log.Println(err)
			}
			return i
		}
	}
	log.Panic("query was not successful")
	return 0
}
