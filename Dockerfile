# Builder

FROM golang:1.12 as builder

RUN go get github.com/gin-contrib/multitemplate
RUN go get github.com/gin-contrib/sessions
RUN go get github.com/gin-gonic/gin
RUN go get github.com/sirupsen/logrus
RUN go get github.com/stretchr/testify/assert

WORKDIR /go/src/github.com/semiscone/website
COPY src /go/src/github.com/semiscone/website
RUN go get -d .

RUN go install

# Production

FROM ubuntu:bionic
COPY --from=builder /go/bin/website /website

ARG WORKDIR="/site"
WORKDIR ${WORKDIR}
VOLUME  ${WORKDIR}

COPY src/static/alithon static/alithon
COPY src/static/AdminLTE/dist static/dist
COPY src/static/AdminLTE/plugins static/plugins
COPY src/templates templates

EXPOSE 80

ENTRYPOINT [ "/website" ]