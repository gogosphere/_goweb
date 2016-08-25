# docker build -t local/golang-web .
FROM golang:alpine

ADD . /go/src/app

RUN apk add --update go git
RUN cd /go/src/app/ && go get -v
RUN cd /go/src/app/ && go run ./main.go

# EXPOSE 8080:8080

# docker commit ->containerid:0f83dd309341<- whancock/goweb
