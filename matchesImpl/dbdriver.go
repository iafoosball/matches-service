package usersImpl

import (
	"log"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"time"
	"fmt"
)

var db driver.Database

func DB() driver.Database {
	if db == nil {
		time.Sleep(2*time.Second)

		fmt.Println("trying to connect")
		conn, err := http.NewConnection(http.ConnectionConfig{
			// Endpoints: []string{"http://http://192.38.56.114:9002"},
			Endpoints: []string{"http://arangodb:8529"},
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println()
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("joe", "joe"),
	})
	db, _ = c.Database(nil, "iaf")
	//if err != nil {
	//	log.Fatal(err)
	//}
	}

	return db
}

func Col(collection string) driver.Collection {
	db := DB()
	col, err := db.Collection(nil, collection)
	if err != nil {
		log.Fatal(err)
	}
	return col
}

func EdgeCol(c string) driver.Collection {
	db := DB()
	col, err := db.Collection(nil, c)
	if err != nil {
		log.Fatal(err)
	}
	return col
}
