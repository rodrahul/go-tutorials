## A hello world example with Go
#FROM golang:1.8-onbuild
#MAINTAINER rod.rahul@gmail.com

FROM golang:1.12.4

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]