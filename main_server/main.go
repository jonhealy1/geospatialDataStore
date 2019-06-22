package main

import (
	"encoding/json"
	"flag"
	"fmt"

	//rejson "go-rejson"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
	rejson "github.com/nitishm/go-rejson"
)

// to run redis and rejson
// docker run -p 6379:6379 --name redis-redisjson redislabs/rejson:latest
//

// Name - student name
type Name struct {
	First  string `json:"first,omitempty"`
	Middle string `json:"middle,omitempty"`
	Last   string `json:"last,omitempty"`
}

// Student - student object
type Student struct {
	Name Name `json:"name,omitempty"`
	Rank int  `json:"rank,omitempty"`
}

func Example_JSONSet(rh *rejson.Handler) {

	student := Student{
		Name: Name{
			"Mark",
			"S",
			"Pronto",
		},
		Rank: 1,
	}
	res, err := rh.JSONSet("student", ".", student)
	if err != nil {
		log.Fatalf("Failed to JSONSet")
		return
	}

	if res.(string) == "OK" {
		fmt.Printf("Success: %s\n", res)
	} else {
		fmt.Println("Failed to Set: ")
	}

	studentJSON, err := redis.Bytes(rh.JSONGet("student", "."))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return
	}

	readStudent := Student{}
	err = json.Unmarshal(studentJSON, &readStudent)
	if err != nil {
		log.Fatalf("Failed to JSON Unmarshal")
		return
	}

	fmt.Printf("Student read from redis : %#v\n", readStudent)
}

func main() {
	redisurl := os.Getenv("REDIS_URL")
	var addr = flag.String("Server", redisurl, "Redis server address")

	rh := rejson.NewReJSONHandler()
	flag.Parse()

	fmt.Println("Hello!")

	// Redigo Client
	conn, err := redis.Dial("tcp", *addr)
	if err != nil {
		log.Fatalf("Failed to connect to redis-server @ %s", *addr)
	}
	defer func() {
		_, err = conn.Do("FLUSHALL")
		err = conn.Close()
		if err != nil {
			log.Fatalf("Failed to communicate to redis-server @ %v", err)
		}
	}()
	rh.SetRedigoClient(conn)
	fmt.Println("Executing Example_JSONSET for Redigo Client")
	Example_JSONSet(rh)

	/* NET/HTTP */

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Hello World from path: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":"+PORT, nil)
}
