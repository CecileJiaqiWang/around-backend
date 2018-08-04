package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"strconv"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

const (
	// Default range
	DISTANCE = "200km"
)
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
	http.HandleFunc("/search", handlerSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


/**
 * Handle post.
 */
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

/**
 * Handle search.
 */
func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one search request!")

	// Don't care the errors at this point.
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)

	ran := DISTANCE

	if val := r.URL.Query().Get("range"); val != "" {
		ran = val + "km"
	}

	fmt.Println("range is", ran)

	// Test post
	p := &Post{
		User:"1111",
		Message:"test test",
		Location: Location{
			Lat:lat,
			Lon:lon,
		},
	}

	// Convert go struct to json string.
	js, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
