
## Build
```sh
docker-compose build
```
## Run Servers
```sh
docker-compose up -d
```

## Run Performance Test
```sh
docker run --rm -v /etc/hosts:/etc/hosts \
  skandyla/wrk -t5 -c10 -d30  http://192.168.0.100:8081/mockserver/status
```