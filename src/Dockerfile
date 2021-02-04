# https://medium.com/travis-on-docker/multi-stage-docker-builds-for-creating-tiny-go-images-e0e1867efe5a
# Multi-Stage Docker Builds

FROM golang AS builder
RUN apk --no-cache add build-base git
WORKDIR /go/src/bin
ADD ./src ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o application .

FROM golang
WORKDIR /app/src/controle_de_acesso_core
ENV GOPATH=/app
COPY --from=builder /go/src/bin/application /app/src/controle_de_acesso_core
RUN go build webapp.go
EXPOSE 8080
ENTRYPOINT ["./webapp"]
