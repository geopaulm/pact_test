FROM golang:1.19-alpine

ENV GO111MODULE=on
ENV GOFLAGS=-mod=mod

ENV APP_HOME /go/src/server
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
EXPOSE 8080
CMD ["go", "run", "."]