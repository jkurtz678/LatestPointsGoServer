package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var base_url = "https://track.onestepgps.com/v3/api/public"
var api_key = ""

// enable cross origin requests
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// handles http requests for latest device points
func latestPointsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling latest points request...")

	response := requestLatestPoints()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error: failed to read response body!")
		return
	}

	enableCors(&w)
	fmt.Fprintf(w, string(data))
	fmt.Println("Successful response.")
}

// requests latest gps device points from OneStepGPS API
func requestLatestPoints() *http.Response {
	//build request url
	latest_devices := "/device?latest_point=true&api-key="
	request_url := base_url + latest_devices + api_key

	response, err := http.Get(request_url)
	if err != nil {
		fmt.Printf("The HTTP request fails with error %s\n", err)
	}
	return response
}

// determines if api key is valid by sending test request to API
func checkValidAPIKey() bool {
	fmt.Println("Checking if api key is valid...")
	res := requestLatestPoints()
	if res.StatusCode >= 200 && res.StatusCode <= 300 {
		fmt.Println("Valid key!")
		return true
	} else {
		fmt.Println("ERROR: invalid api key!")
		return false
	}
}

func main() {
	//extract api key from cmd line args
	if len(os.Args) != 2 {
		fmt.Println("USAGE: ./main <api-key>")
		return
	}
	api_key = os.Args[1]

	if checkValidAPIKey() {
		//attach request handlers
		http.HandleFunc("/latest", latestPointsHandler)

		//start server
		fmt.Println("Server listening on port 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
