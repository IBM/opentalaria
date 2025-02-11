# Builder image
FROM golang:alpine3.21 AS builder

COPY . .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -buildvcs=false -o /tmp/opentalaria ./

# Runtime image
FROM alpine:latest

WORKDIR /

COPY --from=builder /tmp/opentalaria /openatalaria
COPY server.properties /

EXPOSE 9092

ENTRYPOINT [ "/openatalaria", "-c", "server.properties" ]
