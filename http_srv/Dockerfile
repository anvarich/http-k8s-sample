FROM golang:1.11.1 AS build-env
RUN mkdir /src
ADD http_srv.go /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a http_srv.go

# final stage
FROM scratch
COPY --from=build-env /src/http_srv /app/http_srv
ENTRYPOINT ["/app/http_srv"]
EXPOSE 8080