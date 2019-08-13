FROM golang:alpine
RUN apk add git



COPY . /go/src/web_server/
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/GeertJohan/go.rice
RUN go get -u github.com/GeertJohan/go.rice/rice
RUN go get -u github.com/GeertJohan/go.rice/embedded
RUN go get /go/src/web_server

RUN echo "installed web server"
WORKDIR /go/src/web_server
RUN rice embed-go

EXPOSE 80

ENTRYPOINT /go/bin/web_server