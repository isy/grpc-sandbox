FROM golang:1.14.2-alpine AS build
WORKDIR /go/src/github.com/isy/grpc-sandbox/payment

RUN apk add --no-cache ca-certificates git make gcc g++ \
	&& GO111MODULE=off go get -u github.com/oxequa/realize

COPY . .

RUN make mod \
	&& make build


FROM alpine:3.10 AS app

COPY --from=build /go/src/github.com/isy/grpc-sandbox/payment/build/payment /usr/local/bin/payment

RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["payment"]