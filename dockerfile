FROM golang:1.19

WORKDIR /SE_MIM22_WEBSHOP_MONO

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

RUN go mod download

ENTRYPOINT go build  && ./SE_MIM22_WEBSHOP_MONO