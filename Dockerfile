# https://medium.com/travis-on-docker/multi-stage-docker-builds-for-creating-tiny-go-images-e0e1867efe5a
# Multi-Stage Docker Builds

FROM golang:alpine AS builder
RUN apk --no-cache add build-base git
WORKDIR /go/src/bin
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o application src/webapp.go

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/bin/application /app/
EXPOSE 8000
ENTRYPOINT ./application