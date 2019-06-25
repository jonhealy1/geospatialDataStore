package main

// Catalog struct
type Catalog struct {
	RId          int    `json:"rid,omitempty"`
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
	Href  string `json:"href,omitempy"`
	Rel   string `json:"rel,omitempy"`
	Title string `json:"title,omitempty"`
}

//type Links []Link
type Catalogs []Catalog

//var catalogBasic map[string]string
