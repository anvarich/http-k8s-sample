FROM golang:1.11.1 AS build-env
RUN mkdir /src
ADD http_proxy.go /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a http_proxy.go

# final stage
FROM scratch
COPY --from=build-env /src/http_proxy /app/http_proxy
ENV BACKEND_URL=""
ENTRYPOINT ["/app/http_proxy"]
EXPOSE 8088