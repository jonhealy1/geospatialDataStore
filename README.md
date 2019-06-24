install docker

run from the main folder that contains the docker-compose.yml file  
$ docker-compose build  
$ docker-compose up  

1. navigate to localhost:3030 - should see the hello message  

navigate to localhost:3030/catalogs - should see some catalogs data  

$ curl -H "Content-Type: application/json" -d '{"id":"New Catalog"}' http://localhost:3030/catalogs

refresh localhost:3030/catalogs - should see new data added

my-golang-app-run_1  | 2019/06/24 17:19:03 GET	/catalogs	CatalogIndex	55.8µs  
my-golang-app-run_1  | 2019/06/24 17:19:08 POST	/catalogs	CatalogCreate	384.8µs  
my-golang-app-run_1  | 2019/06/24 17:19:10 GET	/catalogs	CatalogIndex	33.3µs

2. should see this in the terminal window:

my-golang-app-run_1  | Executing Example_JSONSET for Redigo Client  
my-golang-app-run_1  | Success: OK  
my-golang-app-run_1  | Catalog read from redis : main.Catalog{RId:0, Stac_version:"0.6.1", Id:"sample", Title:"Sample catalog", Description:"This is a very basic sample catalog.", Links:main.Links{Href:"http://www.example.com/sample-catalog/catalog.json", Rel:"item"}}


TODO: validate the json schema client side in javascript

Things to know:
-   Redis: especially RedisJSON - https://oss.redislabs.com/redisjson/
-   Docker-compose: there is a docker-compose.yml file and then a dockerfile in
    the main_server folder
-   Go: gorilla/mux and redigo libraries