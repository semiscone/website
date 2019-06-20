# Builder

FROM golang:1.12 as builder

RUN go get github.com/gin-contrib/multitemplate
RUN go get github.com/gin-contrib/sessions
RUN go get github.com/gin-gonic/gin
RUN go get github.com/sirupsen/logrus
RUN go get github.com/stretchr/testify/assert
RUN go get github.com/go-pg/pg
RUN go get github.com/jinzhu/gorm
RUN go get github.com/lib/pq

WORKDIR /go/src/github.com/semiscone/website
COPY src /go/src/github.com/semiscone/website

RUN go install


FROM ubuntu:bionic

RUN apt-get update
RUN apt-get upgrade -y

COPY --from=builder /go/bin/website /website

ARG WORKDIR="/site"
WORKDIR ${WORKDIR}
VOLUME  ${WORKDIR}

COPY src/static static

EXPOSE 80

ENTRYPOINT [ "/website" ]