package main

// Item struct
type Item struct {
	RId          int        `json:"rid,omitempty"`
	Stac_version string     `json:"stac_version,omitempty"`
	Type         string     `json:"type,omitempty"`
	Id           string     `json:"id,omitempty"`
	Bbox         []float64  `json:"bbox,omitempty"`
	Geometry     Geometry   `json:"geometry,omitempty"`
	Properties   Properties `json:"properties,omitempty"`
	Links        Links      `json:"links,omitempty"`
	Assets       Assets     `json:"assets,omitempty"`
}

type Geometry struct {
	Type        string    `json:"type,omitempty"`
	Coordinates []float64 `json:"coordinates,omitempty"`
}

type Properties struct {
	Datetime       string  `json:"datetime,omitempty"`
	Eo_sun_azimuth float64 `json:"eo:sun_azimuth,omitempty"`
	Eo_cloud_cover float64 `json:"eo:cloud_cover,omitempty"`
}

type Assets struct {
	Analytic Analytic `json:"analytic,omitempty"`
}

type Analytic struct {
	Href    string `json:"href,omitempty"`
	Title   string `json:"title,omitempty"`
	Product string `json:"product,omitempty"`
}

// {
// 	"type": "Feature",
// 	"id" : "CS3-20160503_132131_05",
// 	"bbox": [-122.59750209, 37.48803556, -122.2880486, 37.613537207],
// 	"geometry": {
// 	  "type": "Polygon",
// 	  "coordinates": [
// 		[
// 		  [-122.308150179, 37.488035566],
// 		  [-122.597502109, 37.538869539],
// 		  [-122.576687533, 37.613537207],
// 		  [-122.288048600, 37.562818007],
// 		  [-122.308150179, 37.488035566]
// 		]
// 	  ]
// 	},
// 	"properties": {
// 	  "datetime": "2016-05-03T13:22:30.040Z",
// 	  "eo:sun_azimuth": 168.7,
// 	  "eo:cloud_cover": 0.12,
// 	  "eo:off_nadir": 1.4,
// 	  "eo:platform": "COOLSAT2",
// 	  "eo:sun_elevation": 33.4,
// 	  "eo:gsd": 0.512,
// 	  "cs:type": "scene",
// 	  "cs:anomalous_pixels": 0.14,
// 	  "cs:earth_sun_distance": 1.0141560,
// 	  "cs:sat_id": "CS3",
// 	  "cs:product_level": "LV1B"
// 	},
// 	"links": [
// 	  {"rel": "self", "href": "http://cool-sat.com/catalog/CS3-20160503_132130_04/CS3-20160503_132130_04.json"},
// 	  {"rel": "thumbnail", "href":"thumbnail.png"},
// 	  {"rel": "root", "href": "http://cool-sat.com/catalog/catalog.json"},
// 	  {"rel": "parent", "href": "http://cool-sat.com/catalog/CS3-20160503_132130_04/catalog.json"},
// 	  {"rel": "collection", "href": "http://cool-sat.com/catalog/CS3-20160503_132130_04/catalog.json"},
// 	  {"rel": "acquisition", "href": "http://cool-sat.com/catalog/acquisitions/20160503_56"}
// 	],
// 	"assets": {
// 	  "analytic": {
// 		"href": "http://cool-sat.com/catalog/CS3-20160503_132130_04/analytic.tif",
// 		"title": "4-Band Analytic",
// 		"product": "http://cool-sat.com/catalog/products/analytic.json"
// 	  },
// 	  "thumbnail": {
// 		"href": "http://cool-sat.com/catalog/CS3-20160503_132130_04/thumbnail.png",
// 		"title": "Thumbnail"
// 	  },
// 	  "udm": {
// 		"href": "http://cool-sat.com/catalog/CS3-20160503_132130_04/UDM.tif",
// 		"title": "Unusable Data Mask",
// 		"product": "http://cool-sat.com/catalog/products/udm.json"
// 	  },
// 	  "json-metadata": {
// 		"href": "http://cool-sat.com/catalog/CS3-20160503_132130_04/extended-metadata.json",
// 		"title": "Extended Metadata",
// 		"product": "http://cool-sat.com/catalog/products/extended-metadata.json"
// 	  },
// 	  "ephemeris": {
// 		"href": "http://cool-sat.com/catalog/CS3-20160503_132130_04/S3-20160503_132130_04.EPH",
// 		"title": "Satellite Ephemeris Metadata",
// 		"product": "http://cool-sat.com/catalog/products/satellite-ephemeris.json"
// 	  }
// 	}

//   }
