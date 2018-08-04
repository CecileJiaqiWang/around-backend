package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}


//{
//	"user": "john",
//	"message": "Test",
//	"location": {
//		"lat": 90,
//		"lon": -120
//
//  }
//}
type Post struct {
	// Tell decoder the actual name in JSON.
	User     string   `json:"user"`
	Message  string   `json:"message"`
	Location Location `json:"location"`
}

func main() {
	fmt.Println("Started service!")
	http.HandleFunc("/post", handlerPost)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one post request!")
	// Decode request body.
	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err != nil {
		// On error. Stops goroutine.
		panic(err)
	}
	fmt.Fprintf(w, "Post received %s\n", p.Message)
}
