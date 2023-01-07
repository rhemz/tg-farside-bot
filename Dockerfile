FROM golang:1.19.4-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN apk update && apk upgrade && apk add --no-cache ca-certificates git
RUN update-ca-certificates
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tg-farside-bot -ldflags="-w -s" .

FROM scratch
MAINTAINER russ@russ.wtf

ENV LOG_JSON=true
ENV LISTEN_PORT 8080

WORKDIR /app

COPY --from=builder /app/tg-farside-bot .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE $LISTEN_PORT
ENTRYPOINT ["./tg-farside-bot"]