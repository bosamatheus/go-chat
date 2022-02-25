FROM golang:1.17-alpine

ARG CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /gochat

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o gochat cmd/gochat/main.go

CMD ["./gochat"]
