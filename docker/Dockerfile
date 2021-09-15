FROM golang:1.14

RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work

CMD ["go", "run", "server.go"]
