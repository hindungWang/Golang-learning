FROM golang:1.9.4-alpine3.7 as builder
WORKDIR /go/src/github.com/choerodon/
COPY server/ /
COPY vendor/ /usr/local/go/src/


ADD pv.yaml /pv.yaml

RUN \
  apk update &&\
  apk add git

RUN go build /myserver.go

EXPOSE 9090

CMD ["/myserver"]