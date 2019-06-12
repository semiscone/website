FROM golang:1.12 as builder

RUN go get github.com/gin-contrib/multitemplate
RUN go get github.com/gin-contrib/sessions
RUN go get github.com/gin-gonic/gin
RUN go get github.com/sirupsen/logrus
RUN go get github.com/stretchr/testify/assert

ENV PROJ $GOPATH/src/alithon/hummingbird
COPY src ${PROJ}
WORKDIR ${PROJ}

RUN go build


FROM ubuntu:bionic

WORKDIR /app
COPY src/static/alithon static/alithon
COPY src/static/AdminLTE/dist static/dist
COPY src/static/AdminLTE/plugins static/plugins
COPY src/templates templates

COPY --from=builder /go/src/alithon/hummingbird/hummingbird .

EXPOSE 80

ENTRYPOINT [ "/app/hummingbird" ]