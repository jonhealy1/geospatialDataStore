install docker

run from the main folder that contains the docker-compose.yml file  
$ docker-compose build  
$ docker-compose up  

1. navigate to localhost:3030 - should see the hello message  
navigate to localhost:3030/catalogs - should see some catalogs data

2. should see this in the terminal window:

my-golang-app-run_1  | Executing Example_JSONSET for Redigo Client  
my-golang-app-run_1  | Success: OK  
my-golang-app-run_1  | Catalog read from redis : main.Catalog{CatalogB:map[string]string{"description":"This is a very basic sample catalog.", "id":"sample", "stac_version":"0.6.1", "title":"Sample catalog"}, Links:main.Links{Href:"http://www.example.com/sample-catalog/catalog.json", Rel:"root"}}


TODO: validate the json schema client side in javascript

Things to know:
-   Redis: especially RedisJSON - https://oss.redislabs.com/redisjson/
-   Docker-compose: there is a docker-compose.yml file and then a dockerfile in
    the main_server folder
-   Go: gorilla/mux and redigo libraries