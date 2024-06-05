# nginx-prometheus-exporter


## How to run it?
steps

1. spin up nginx service
```
docker-compose up
```

2. run nginx-prometheus-exporter
```
go run cmd/exporter/main.go
```

3. spin up prometheus and grafana
```
cd prometheus && docker-compose up
```

4. prometheus dashboard: http://localhost:9090/

5. grafana dashboard: http://localhost:3000/

6. to verify nginx-prometheus-exporter working
```
curl http://localhost:6777/metrics
```
