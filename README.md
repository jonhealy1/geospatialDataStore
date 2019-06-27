install docker

run from the main folder that contains the docker-compose.yml file  
$ docker-compose build  
$ docker-compose up  

1. navigate to localhost:3030 - should see the hello message  

navigate to localhost:3030/catalogs - should see some catalogs data  

$ curl -H "Content-Type: application/json" -d '{"id":"New Catalog"}' http://localhost:3030/catalogs

refresh localhost:3030/catalogs - should see new data added

main_server  | 2019/06/24 17:19:03 GET	/catalogs	CatalogIndex	55.8µs  
main_server  | 2019/06/24 17:19:08 POST	/catalogs	CatalogCreate	384.8µs 
main_server  | 2019/06/24 17:19:10 GET	/catalogs	CatalogIndex	33.3µs

2. should see this in the terminal window:

main_server  | Executing Example_JSONSET for Redigo Client  
main_server  | Success: OK  
main_server  | Catalog read from redis : main.Catalog{RId:0, Stac_version:"0.6.1", Id:"sample", Title:"Sample catalog", Description:"This is a very basic sample catalog.", Links:main.Links{Href:"http://www.example.com/sample-catalog/catalog.json", Rel:"item"}}


TODO: validate the json schema client side in javascript

Things to know:
-   Redis: especially RedisJSON - https://oss.redislabs.com/redisjson/
-   Docker-compose: there is a docker-compose.yml file and then a dockerfile in
    the main_server folder
-   Go: gorilla/mux and redigo libraries

-----------------------------------------------------  
web_server folder  

The web_server is made using rice to bundle the html, css, and javascript into a golang application.  

Tutorial: https://www.thepolyglotdeveloper.com/2017/03/bundle-html-css-javascript-served-golang-application/  

To rebuild the web stuff:  
$ rice embed.go  

This builds a file called rice-box.go that handles everything automatically.

$ go build  
$ go run main.go  

navigate to localhost:12345  

