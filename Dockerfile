FROM golang:latest

MAINTAINER HeHuang "huanghe.poly@gmail.com"

WORKDIR $GOPATH/src/github.com/huanghe314/geektime_cloud_native_course

ADD . $GOPATH/src/github.com/huanghe314/geektime_cloud_native_course

RUN go build -o srv_practice

ENTRYPOINT ./srv_practice
