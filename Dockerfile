FROM golang:1.17.3-alpine AS builder
RUN apk add gcc musl-dev lame-dev
WORKDIR /go/src/github.com/m-butterfield/mattbutterfield.com
COPY go.* ./
RUN go mod download
ADD . /go/src/github.com/m-butterfield/mattbutterfield.com
RUN go build -o bin/server cmd/server.go
RUN go build -o bin/worker cmd/worker.go

FROM alpine:latest
RUN apk add lame
WORKDIR /root
COPY --from=builder /go/src/github.com/m-butterfield/mattbutterfield.com/bin/ ./bin/
CMD ["bin/server"]
