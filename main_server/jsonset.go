package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	rejson "github.com/nitishm/go-rejson"
)

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
