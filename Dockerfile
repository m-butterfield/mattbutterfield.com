FROM golang:1.17.1-alpine AS builder
WORKDIR /go/src/github.com/m-butterfield/mattbutterfield.com
COPY go.* ./
RUN go mod download
ADD . /go/src/github.com/m-butterfield/mattbutterfield.com
RUN go build -o bin/server server.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /go/src/github.com/m-butterfield/mattbutterfield.com/bin/ ./bin/
CMD ["bin/server"]
