FROM golang:1.16 as build

WORKDIR /go/src/github.com/s8sg/performance_test/go_gin

COPY go.mod         go.mod
COPY go.sum         go.sum
COPY main.go        .

RUN CGO_ENABLED=1  go build -a -installsuffix cgo -o go_gin .

CMD ["./go_gin"]