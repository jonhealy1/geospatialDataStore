package main

import (

	//rejson "go-rejson"

	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	rejson "github.com/nitishm/go-rejson"

	"github.com/gomodule/redigo/redis"
)

// to run redis and rejson
// docker run -p 6379:6379 --name redis-redisjson redislabs/rejson:latest
//
// gorilla/mux routing tutorial:
// https://thenewstack.io/make-a-restful-json-api-go/

func main() {

	// JSON SET
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

	time.Sleep(2 * time.Second)

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

	redisurl := os.Getenv("REDIS_URL")
	var addr = flag.String("Server", redisurl, "Redis server address")

	rh := rejson.NewReJSONHandler()
	flag.Parse()

	fmt.Println("Hello!")

	// Redigo Client
	conn, err := redis.Dial("tcp", *addr)
	if err != nil {
		log.Fatalf("Failed to connect to redis-server @ %s", *addr)
	}
	defer func() {
		_, err = conn.Do("FLUSHALL")
		err = conn.Close()
		if err != nil {
			log.Fatalf("Failed to communicate to redis-server @ %v", err)
		}
	}()
	rh.SetRedigoClient(conn)
	fmt.Println("Executing Example_JSONSET for Redigo Client")
	catalog := Example_JSONSet(rh)
	fmt.Println("Catalog Href: ", catalog.Links.Href)

	/* GORILLA/MUX */

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", Index)
	// router.HandleFunc("/catalogs", CatalogIndex)
	// router.HandleFunc("/catalogs/{catalogId}", CatalogShow)
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+PORT, router))
	// log.Fatal(http.ListenAndServe(":8080", router))
}
