FROM golang:1.17.8-alpine

LABEL maintainer="nekinci <niyaziekinci5050@gmail.com>"
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get .
RUN go build

ENTRYPOINT ["/go/bin/springbootyamlgenerator"]

