FROM golang:latest AS builder

COPY . $GOPATH/src/github.com/dylankilkenny/watch-cash/server
WORKDIR $GOPATH/src/github.com/dylankilkenny/watch-cash/server

# install dependencies
RUN set -x && \ 
    go get github.com/githubnemo/CompileDaemon && \
    go get github.com/golang/dep/cmd/dep && \
    dep init && \
    dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /watch-cash-server
EXPOSE 3001

# #second stage to obtain a very small image
# FROM alpine
# COPY --from=builder /watch-cash-server .
# EXPOSE 3001

# check if db is ready
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh
RUN apk --update --no-cache add  bash
