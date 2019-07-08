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
RUN go get github.com/go-sql-driver/mysql

WORKDIR /go/src/github.com/semiscone/website
COPY src /go/src/github.com/semiscone/website

RUN go install


FROM ubuntu:bionic

RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/* \
    && localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

ARG WORKDIR="/website"
WORKDIR ${WORKDIR}

COPY --from=builder /go/bin/website ${WORKDIR}/website
COPY src/static ${WORKDIR}/static
COPY src/templates ${WORKDIR}/templates

# EXPOSE 5000

ENTRYPOINT ["/website/website"]
