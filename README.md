
## Build
```sh
docker-compose build
```
## Run Servers
```sh
docker-compose down && docker-compose up -d
```

## Run Performance Test
## Install K6
```
brew install k6
```

## Run Test
```
k6 run -u 1000 -d 10m spring_performance.js
k6 run -u 1000 -d 10m go_performance.js
```
