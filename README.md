# LatestPointsGoServer
Simple Go web server that retrieves the latest gps coordinates from the OneStepGPS API

The Go server requires an API key command line argument to startup. After checking that the key is valid, the server will listen for requests on port :8080. 

This server has a single endpoint at http://localhost:8080/latest. Upon requesting this endpoint, the server will retrieve the latest points from devices associated with the provided API key, and return this JSON data to the client. The data is retrieved  from this link: https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=[APIKEY]

# Usage:

`go run main.go <api-key>`

or with executable

`./main <api-key>`
