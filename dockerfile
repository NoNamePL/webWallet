FROM golang:1.23.2

WORKDIR /build

COPY . . 

RUN go mod tidy

RUN go build -o main.go cmd/main.go -b 0.0.0.0
CMD [". /main"]
