# Bases for building and running the app
FROM golang:1.22-alpine AS builder-base
WORKDIR /go/src/github.com/m-butterfield/mattbutterfield.com
COPY go.* ./
RUN go mod download
ADD . /go/src/github.com/m-butterfield/mattbutterfield.com

FROM alpine:latest AS runner-base
WORKDIR /root

# Run build
FROM builder-base AS server-builder
RUN go build -o bin/server cmd/server/main.go

FROM builder-base AS worker-builder
RUN apk add pkgconfig vips-dev gcc musl-dev lame-dev
RUN go build -o bin/worker cmd/worker/main.go
# tippecanoe
ARG TIPPECANOE_RELEASE="2.53.0"
RUN apk add --no-cache sudo git g++ make libgcc libstdc++ sqlite-libs sqlite-dev zlib-dev bash \
 && cd /root \
 && git clone https://github.com/felt/tippecanoe.git tippecanoe \
 && cd tippecanoe \
 && git checkout tags/$TIPPECANOE_RELEASE \
 && cd /root/tippecanoe \
 && make \
 && make install \
 && cd /root \
 && rm -rf /root/tippecanoe

# Copy the built executable to the runner
FROM runner-base AS server
COPY --from=server-builder /go/src/github.com/m-butterfield/mattbutterfield.com/bin/ ./bin/
CMD ["bin/server"]

FROM runner-base AS worker
RUN apk add lame-dev vips # gdal-tools - if we need to convert CRS
COPY --from=worker-builder /go/src/github.com/m-butterfield/mattbutterfield.com/bin/ ./bin/
COPY --from=worker-builder /usr/local/bin/* /usr/local/bin/
CMD ["bin/worker"]
