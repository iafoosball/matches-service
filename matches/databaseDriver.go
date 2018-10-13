package matches

import (
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
	"strconv"
	"time"
)

var (
	database         driver.Database
	host             string
	port             int
	databaseUser     string
	databasePassword string
	databaseName     = "iaf-matches"

	goalsColName   = "goals"
	goalsCol       driver.Collection
	matchesColName = "matches"
	matchesCol     driver.Collection
)

func InitDatabase(dbHost string, dbPort int, dbUser string, dbPassword string) {
	port = dbPort
	host = dbHost
	databaseUser = dbUser
	databasePassword = dbPassword

	fmt.Println(host + strconv.Itoa(dbPort) + dbUser + dbPassword)
	c := 0
	for database == nil && c < 10 {
		initDatabaseDriver(dbUser, dbPassword)
		c++
		if database == nil {
			log.Println("Database not reachable! Sleep seconds: " + strconv.Itoa(c))
			time.Sleep(time.Duration(c) * 1000 * time.Millisecond)
		}
	}
}

func initDatabaseDriver(user string, password string) {
	if conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://" + host + ":" + strconv.Itoa(port)},
	}); err == nil {
		if client, err := driver.NewClient(driver.ClientConfig{
			Connection:     conn,
			Authentication: driver.BasicAuthentication(user, password),
		}); err == nil {
			if dbExists, e := client.DatabaseExists(nil, databaseName); !dbExists && e == nil {
				if database, err = client.CreateDatabase(nil, databaseName, &driver.CreateDatabaseOptions{
					[]driver.CreateDatabaseUserOptions{
						{
							UserName: user,
						},
					},
				}); err != nil {
					log.Fatal(err)
					return
				}
				log.Println("Connected to database: " + databaseName)
				Collection(goalsColName)
				Collection(matchesColName)
			} else if e != nil {
				log.Println(e)
			} else {
				database, err = client.Database(nil, databaseName)
				Collection(goalsColName)
				Collection(matchesColName)
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

func Collection(name string) driver.Collection {
	if database == nil {
		InitDatabase(host, port, databaseUser, databasePassword)
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

func initCollection(name string, colType int) driver.Collection {
	if exists, err := database.CollectionExists(nil, name); !exists {
		if col, e := database.CreateCollection(nil, name, &driver.CreateCollectionOptions{
			Type: driver.CollectionType(colType),
		}); e != nil {
			return col
		} else if err != nil {
			log.Println(err)
		}
	} else if err != nil {
		log.Println(err)
	} else {

		if col, err := database.Collection(nil, name); err == nil {
			return col
		} else {
			log.Println(err)
		}
	}
	return nil
}
