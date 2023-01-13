FROM golang:alpine

WORKDIR /SE_MIM22_WEBSHOP_MONO

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./SE_MIM22_WEBSHOP_MONO