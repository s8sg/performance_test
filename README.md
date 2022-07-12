
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

## Install K6
```
brew install k6
```

## Run Test
```
k6 run -u 1000 -d 10m spring_performance.js
k6 run -u 1000 -d 10m go_performance.js
```
