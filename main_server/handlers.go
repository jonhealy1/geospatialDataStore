package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func CatalogIndex(w http.ResponseWriter, r *http.Request) {
	catalogs := Catalogs{
		Catalog{Id: "Nanaimo Lidar",
			Title: "1986-1989",
			Links: Links{
				Href: "https://thenewstack.io/make-a-restful-json-api-go/",
			},
		},
		Catalog{Id: "Space Station"},
		Catalog{
			Stac_version: "0.6.1",
			Id:           "sample",
			Title:        "Sample catalog",
			Description:  "This is a very basic sample catalog.",
			Links: Links{
				Href: "http://www.example.com/sample-catalog/catalog.json",
				Rel:  "item",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(catalogs); err != nil {
		panic(err)
	}
}

func CatalogShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	catalogId := vars["catalogId"]
	fmt.Fprintln(w, "Catalog show:", catalogId)
}
