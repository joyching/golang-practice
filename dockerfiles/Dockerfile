FROM golang:1.11-stretch

RUN apt-get update && \
    apt-get install git && \
    go get -u github.com/kardianos/govendor

WORKDIR $GOPATH/src/github.com/joyching/golang-practice
