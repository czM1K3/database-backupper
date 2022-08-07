FROM golang:1.18-alpine as build
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./backup .


FROM alpine

COPY --from=build /app/backup /db-backup

RUN apk add --no-cache mongodb-tools ca-certificates && \
	rm -rf /var/cache/apk/*

ENV DB_TYPE MONGODB
ENV BACKUP_PATH /backup

VOLUME ["/backup"]

CMD ["/db-backup"]
