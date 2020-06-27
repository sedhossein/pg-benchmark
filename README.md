# pg-benchmark

### prometheus
download it, extract, and go to directory. after that:
```
docker run \
    -p 9090:9090 \
    -v ./prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus
```


```
docker run -d -p 3000:3000 grafana/grafana
```

```
docker run --net=host -e DATA_SOURCE_NAME="postgresql://admin:secret@localhost:5432/postgres?sslmode=disable" wrouesnel/postgres_exporter  
```



