FROM golang:1.10.8-alpine3.9
RUN apk update && apk add gcc libc-dev
RUN mkdir /app
ADD . /go/src/github.com/m-butterfield/mattbutterfield.com
WORKDIR /go/src/github.com/m-butterfield/mattbutterfield.com
RUN go build -o bin/server server.go
CMD ["bin/server"]
