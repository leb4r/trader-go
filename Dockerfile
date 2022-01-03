ARG GOVERSION="1.15.6"

FROM docker.io/library/golang:${GOVERSION}-alpine

WORKDIR $GOPATH/src/github.com/leb4r/trader-go
COPY . .

RUN apk add --no-cache build-base && \
    go install && \
    apk del build-base

WORKDIR $GOPATH
ENTRYPOINT [ "trader-go" ]
