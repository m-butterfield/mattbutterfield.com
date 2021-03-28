FROM golang:1.16.2-alpine
RUN mkdir /app
ADD . /go/src/github.com/m-butterfield/mattbutterfield.com
WORKDIR /go/src/github.com/m-butterfield/mattbutterfield.com
RUN go build -o bin/server server.go
CMD ["bin/server"]
