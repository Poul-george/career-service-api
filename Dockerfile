# goバージョン
FROM golang:latest

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN go install golang.org/x/tools/cmd/goimports@latest

WORKDIR /go/src
COPY go.mod go.sum /go/src/
RUN go mod download
RUN go mod verify

EXPOSE 1324

# CMD ["go", "run", "api/main.go"]