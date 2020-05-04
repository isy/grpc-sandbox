FROM golang:1.14.2-alpine AS build
WORKDIR /go/src/github.com/isy/grpc-sandbox/user

RUN apk add --no-cache ca-certificates git make gcc g++ \
	&& CGO_ENABLED=1 go get -u github.com/rubenv/sql-migrate/sql-migrate \
	&& GO111MODULE=off go get -u github.com/oxequa/realize

COPY . .

RUN make mod \
	&& make build


FROM alpine:3.10 AS app

COPY --from=build /go/src/github.com/isy/grpc-sandbox/user/build/server /usr/local/bin/user-service

RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["user-service"]