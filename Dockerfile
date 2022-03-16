FROM golang:1.17.8-alpine
LABEL maintainer="nekinci <niyaziekinci5050@gmail.com>"

ADD . /go/src/app
WORKDIR /go/src/app
RUN go get .
RUN go build

RUN addgroup -g 1001 -S spg && adduser -u 1001 -S spg  -G spg
USER spg
RUN echo $HOME

ENTRYPOINT ["/go/bin/spg"]

