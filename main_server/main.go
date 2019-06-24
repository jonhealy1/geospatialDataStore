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
	"github.com/gorilla/mux"
	rejson "github.com/nitishm/go-rejson"
)

// to run redis and rejson
// docker run -p 6379:6379 --name redis-redisjson redislabs/rejson:latest
//
// gorilla/mux routing tutorial:
// https://thenewstack.io/make-a-restful-json-api-go/

// Catalog struct
type Catalog struct {
	Stac_version string `json:"stac_version,omitempy"`
	Id           string `json:"id,omitempy"`
	Title        string `json:"title,omitempy"`
	Description  string `json:"description,omitempy"`
	Links        Links  `json:"links,omitempy"`
}

// type Catalog struct {
// 	CatalogB map[string]string `json:"catalogb,omitempy"`
// 	Links    Links             `json:"links,omitempy"`
// }

type Links struct {
	Href string `json:"href,omitempy"`
	Rel  string `json:"rel,omitempy"`
}

//type Links []Link
type Catalogs []Catalog

//var catalogBasic map[string]string

func Example_JSONSet(rh *rejson.Handler) Catalog {

	//catalog := make(map[string]string)
	// catalogBasic = map[string]string{
	// 	"stac_version": "0.6.1",
	// 	"id":           "sample",
	// 	"title":        "Sample catalog",
	// 	"description":  "This is a very basic sample catalog.",
	// }
	// //catalog["stac_version"] = "0.6.1"
	// fmt.Println(catalogBasic)

	// catalog := Catalog{
	// 	CatalogB: catalogBasic,
	// 	Links: Links{
	// 		Href: "http://www.example.com/sample-catalog/catalog.json",
	// 		Rel:  "root",
	// 	},
	// }

	catalog := Catalog{
		Stac_version: "0.6.1",
		Id:           "sample",
		Title:        "Sample catalog",
		Description:  "This is a very basic sample catalog.",
		Links: Links{
			Href: "http://www.example.com/sample-catalog/catalog.json",
			Rel:  "item",
		},
	}

	res, err := rh.JSONSet("catalog", ".", catalog)
	if err != nil {
		log.Fatalf("Failed to JSONSet")
		return catalog
	}

	if res.(string) == "OK" {
		fmt.Printf("Success: %s\n", res)
	} else {
		fmt.Println("Failed to Set: ")
	}

	catalogJSON, err := redis.Bytes(rh.JSONGet("catalog", "."))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return catalog
	}

	//readCatalog := Catalog{}
	readCatalog := catalog
	err = json.Unmarshal(catalogJSON, &readCatalog)
	if err != nil {
		log.Fatalf("Failed to JSON Unmarshal")
		return readCatalog
	}

	fmt.Printf("Catalog read from redis : %#v\n", readCatalog)
	return readCatalog
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
	catalog := Example_JSONSet(rh)
	fmt.Println("Catalog Href: ", catalog.Links.Href)

	/* GORILLA/MUX */

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/catalogs", CatalogIndex)
	router.HandleFunc("/catalogs/{catalogId}", CatalogShow)

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
