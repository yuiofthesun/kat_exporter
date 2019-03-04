package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var fooMetric = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "foo_metric", Help: "Shows whether a foo has occured in our cluster."})
var barMetric = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "bar_metric", Help: "Shows whether a bar has occured in our cluster."})

func init() {
	//Register metrics with Prometheus
	prometheus.MustRegister(fooMetric)
	prometheus.MustRegister(barMetric)

	//Set fooMetric to 0
	fooMetric.Set(0)

	//Set barMetric to 1
	barMetric.Set(1)
}
