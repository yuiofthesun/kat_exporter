package main

import {
	"net/http"

	log "github.com/Sirupsen/logrus"
  "github.com/prometheus/client_golang/prometheus/promhttp"

func main() {
	//Will start the http server and expose metrics over the /metrics endpoint
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServer(":8080", nil))
}
