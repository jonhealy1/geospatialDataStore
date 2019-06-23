package main

import (
	"encoding/json"
	"flag"
	"fmt"

	//rejson "go-rejson"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
	rejson "github.com/nitishm/go-rejson"
)

// to run redis and rejson
// docker run -p 6379:6379 --name redis-redisjson redislabs/rejson:latest
//

// Catalog struct
type Catalog struct {
	Stac_version string `json:"stac_version,omitempy"`
	Id           string `json:"id,omitempy"`
	Title        string `json:"title,omitempy"`
	Description  string `json:"description,omitempy"`
	Links        Links  `json:"links,omitempy"`
}

type Links struct {
	Href string `json:"href,omitempy"`
	Rel  string `json:"rel,omitempy"`
}

func Example_JSONSet(rh *rejson.Handler) {

	catalog := Catalog{
		Stac_version: "0.6.1",
		Id:           "sample",
		Title:        "Sample catalog",
		Description:  "This is a very basic sample catalog.",
		Links: Links{
			"item.json",
			"item",
		},
	}
	res, err := rh.JSONSet("catalog", ".", catalog)
	if err != nil {
		log.Fatalf("Failed to JSONSet")
		return
	}

	if res.(string) == "OK" {
		fmt.Printf("Success: %s\n", res)
	} else {
		fmt.Println("Failed to Set: ")
	}

	catalogJSON, err := redis.Bytes(rh.JSONGet("catalog", "."))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return
	}

	readCatalog := Catalog{}
	err = json.Unmarshal(catalogJSON, &readCatalog)
	if err != nil {
		log.Fatalf("Failed to JSON Unmarshal")
		return
	}

	fmt.Printf("Catalog read from redis : %#v\n", readCatalog)
}

func main() {
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
	Example_JSONSet(rh)

	/* NET/HTTP */

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Hello World from path: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":"+PORT, nil)
}
