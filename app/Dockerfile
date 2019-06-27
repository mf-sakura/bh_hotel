FROM golang:1.12.6-alpine3.10 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/mf-sakura/bh_hotel/app
COPY . /go/src/github.com/mf-sakura/bh_hotel/app

RUN apk update \
    && apk add --no-cache git \
    && go get -u github.com/golang/dep/cmd/dep \
    && dep ensure


RUN go build -o bh_hotel /go/src/github.com/mf-sakura/bh_hotel/app/cmd

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/mf-sakura/bh_hotel/app/bh_hotel /app/bh_hotel
WORKDIR /app
EXPOSE 5001
ENTRYPOINT ["/app/bh_hotel"]