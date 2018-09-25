package matchesImpl

import (
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
	"strconv"
	"time"
)

var (
	database             driver.Database = dbDriver()
	initDatabase         bool            = false
	port                 int             = 9002
	url                  string          = "http://localhost:"
	databaseUser         string          = "iaf-matches"
	databaseUserPassword string          = "iaf-matches-2018@secret"
	databaseName         string          = "iaf-matches"

	goalsColName   string            = "goals"
	goalsCol       driver.Collection = Col(goalsColName)
	matchesColName string            = "matches"
	matchesCol     driver.Collection = Col(matchesColName)

	matchesCollectionExists bool = false
	goalsCollectionExists   bool = false
)

func dbDriver() driver.Database {
	counter := 0
	var db driver.Database
	for db == nil && counter < 10 {
		counter++
		// create repeated call library until connection is established with increasing sleep timer
		// can be put in docker-compose with health-check
		fmt.Println("Connecting to " + url + strconv.Itoa(port))
		if conn, err := http.NewConnection(http.ConnectionConfig{
			Endpoints: []string{url + strconv.Itoa(port)},
		}); err == nil {
			if c, e := driver.NewClient(driver.ClientConfig{
				Connection:     conn,
				Authentication: driver.BasicAuthentication("root", "iafoosball@matches for the win"),
			}); e == nil {
				if !initDatabase {
					db = ensureDatabaseName(databaseName, c, db)
					initDatabase = true
				}
			} else {
				log.Fatal(e)
			}

		} else {
			log.Fatal(err)
		}
		if db == nil {
			time.Sleep(2 * time.Second)
		}
	}
	return db
}

func ensureDatabaseName(name string, c driver.Client, db driver.Database) driver.Database {
	fmt.Println("Create new database with user iaf-matches. If already there skip")
	if db == nil {
		if db, err := c.CreateDatabase(nil, databaseName, &driver.CreateDatabaseOptions{
			[]driver.CreateDatabaseUserOptions{
				{
					UserName: databaseUser,
					Password: databaseUserPassword,
				},
			},
		},
		); err == nil {
			fmt.Print("create database")
			if _, err := db.CreateCollection(nil, goalsColName, &driver.CreateCollectionOptions{
				Type: driver.CollectionTypeEdge,
			}); err != nil {
				fmt.Print("sddfff")
				fmt.Println(err)
			}
			db.CreateCollection(nil, matchesColName, &driver.CreateCollectionOptions{
				Type: driver.CollectionTypeDocument,
			})
			fmt.Print("create database")
		} else {
			log.Print(err)
		}
		db, _ = c.Database(nil, "iaf-matches")

		//database.CreateGraph(nil, graphMatches, &driver.CreateGraphOptions{OrphanVertexCollections: {
		//	[1]string{collectionsMatches},
		//}
		//})
	}
	return db
}

func Col(collection string) driver.Collection {
	log.Println("Open collection: " + collection)
	if database != nil {
		col, err := database.Collection(nil, collection)
		if err != nil {
			log.Fatal(err)
		}
		return col
	} else {
		panic("No database!!!")
	}
}
