FROM golang:latest

ADD . /go/src/mux-api/ 

RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/gorilla/mux
RUN go get github.com/saravanan-babu/go-rest-api

RUN go install mux-api

ENTRYPOINT /go/bin/mux-api

EXPOSE 3000 
