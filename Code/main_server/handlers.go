package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func CatalogIndexTwo(w http.ResponseWriter, r *http.Request) {
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

func CatalogIndex(w http.ResponseWriter, r *http.Request) {
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

func CatalogCreate(w http.ResponseWriter, r *http.Request) {
	var catalog Catalog
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &catalog); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateCatalog(catalog)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}
