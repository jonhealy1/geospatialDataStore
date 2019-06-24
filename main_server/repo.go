package main

import "fmt"

var currentId int

var catalogs Catalogs

// Give us some seed data
func init() {
	RepoCreateCatalog(Catalog{Id: "Nanaimo Lidar"})
	RepoCreateCatalog(Catalog{Id: "sample"})
}

func RepoFindCatalog(id int) Catalog {
	for _, c := range catalogs {
		if c.RId == id {
			return c
		}
	}
	// return empty Todo if not found
	return Catalog{}
}

func RepoCreateCatalog(c Catalog) Catalog {
	currentId += 1
	c.RId = currentId
	catalogs = append(catalogs, c)
	return c
}

func RepoDestroyCatalog(id int) error {
	for i, c := range catalogs {
		if c.RId == id {
			catalogs = append(catalogs[:i], catalogs[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Catalog with id of %d to delete", id)
}
