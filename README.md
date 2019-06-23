install docker

install redis

$ docker run -p 6379:6379 --name redis-redisjson redislabs/rejson:latest

run in a separate window - run from the same folder as docker-compose.yml file
$ docker-compose build
$ docker-compose up

1. navigate to localhost:3030 - should see hello world message

2. should see this in the terminal window:

my-golang-app-run_1  | Executing Example_JSONSET for Redigo Client
my-golang-app-run_1  | Success: OK
my-golang-app-run_1  | Catalog read from redis : main.Catalog{Stac_version:"0.6.1", Id:"sample", Title:"Sample catalog", Description:"This is a very basic sample catalog.", Links:main.Links{Href:"item.json", Rel:"item"}}


TODO: store examples from the test folder in redis and fool around with it

Things to know:
-   Redis: especially RedisJSON - https://oss.redislabs.com/redisjson/
-   Docker-compose: there is a docker-compose.yml file and then a dockerfile in
    the main_server folder
-   Go: gorilla/mux and redigo libraries