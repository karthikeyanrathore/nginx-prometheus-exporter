package exporter

import (
    _ "fmt"
    "github.com/prometheus/client_golang/prometheus"
)

type NginxCollector struct {

    ActiveConnection *prometheus.Desc
    Connections *prometheus.Desc

    stats func() ([]NginxStub)
    
}

func NewNginxCollector(stats func() ([]NginxStub)) (prometheus.Collector) {
    return &NginxCollector{
        ActiveConnection: prometheus.NewDesc(
            "active_connection",
            "help_message",
            nil,
            nil,
        ),
        Connections: prometheus.NewDesc(
            "connections",
            "help_message",
            nil,
            nil,
        ),
        stats: stats,
    }
}

func (nc *NginxCollector) Collect(ch chan<- prometheus.Metric) {
    stats := nc.stats()
    for _, s := range stats {
        ch <- prometheus.MustNewConstMetric(
            nc.ActiveConnection,
            prometheus.GaugeValue,
            s.ActiveConnection,
        )
    }
}

func (nc *NginxCollector) Describe(ch chan<- *prometheus.Desc) {
    // Gather metadata about each metric.
	ds := []*prometheus.Desc{
		nc.ActiveConnection,
	}
	for _, d := range ds {
		ch <- d
	}
}