package main


import (

    "fmt"
    "log"
    "io"
    "bytes"
    "time"
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"

    exporter "github.com/karthikeyanrathore/nginx-prometheus-exporter"
)


func HealthCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Println("check check 123")
}

func main() {
    // Nginx prometheus exporter port
    exporterPort := fmt.Sprintf(":%d", 6777)
    // Target Host and Port 
    // Pull metrics from target host
    targetHost := fmt.Sprintf("127.0.0.1")
    targetPort := 8080
    targetPath := fmt.Sprintf("/status")

    // Todo: include Flags
    fmt.Println("Nginx Prometheus Exporter")

    TARGET_NGINX_URI := fmt.Sprintf("http://%s:%d%s", targetHost, targetPort, targetPath)
    nginxStats := func() (int) {
        netClient := &http.Client{
            Timeout: time.Second * 10,
        }
        resp, err := netClient.Get(TARGET_NGINX_URI)
        if err != nil {
            log.Fatalf("netClient error: %s", err)
        }
        defer resp.Body.Close()
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatalf("io.ReadAll error: %s", err)
        }
        r := bytes.NewReader(body)
        exporter.ScanNginxStats(r)
        return 0
    }
    // fmt.Println(nginxStats())
    nginxStats()
    reg := prometheus.NewRegistry()

    promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
    http.Handle("/metrics", promHandler )

    // http.HandleFunc("/", HealthCheck)
    log.Printf("starting nginx exporter on %s/metrics", exporterPort)
    http.ListenAndServe(exporterPort, nil)
}