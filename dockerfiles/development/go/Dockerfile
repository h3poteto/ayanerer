FROM golang:1.7-alpine

# gom installの依存パッケージでbuild-baseが必要になることがある
RUN apk add --no-cache \
    git \
    build-base \
    openssh

RUN mkdir -p /go && \
    mkdir -p /root/ayanerer

ENV GOPATH /go
ENV ECHOROOT /root/ayanerer

ENV HOME /root
WORKDIR ${HOME}

RUN go get github.com/mattn/gom

ENV GO15VENDOREXPERIMENT 0
ENV ECHOENV development

EXPOSE 9090:9090

WORKDIR /root/ayanerer

CMD ["/bin/sh"]
