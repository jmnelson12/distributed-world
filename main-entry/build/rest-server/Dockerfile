# run from source dir: docker build -t jmnelson12/rest-server/test -f ./build/rest-server/Dockerfile .
FROM golang:1.16-alpine AS build

WORKDIR /build/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/rest-server ./cmd/rest-server

#-

FROM alpine:3.13 AS certificates

RUN apk --no-cache add ca-certificates

#-

FROM scratch

COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/bin/rest-server ./bin/rest-server

EXPOSE 8080

CMD ["rest-server"]