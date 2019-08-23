package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CatalogCreate",
		"POST",
		"/catalogs",
		CatalogCreate,
	},
	Route{
		"CatalogIndex",
		"GET",
		"/catalogs",
		CatalogIndex,
	},
	Route{
		"CatalogShow",
		"GET",
		"/catalogs/{catalogId}",
		CatalogShow,
	},
}
