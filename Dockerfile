FROM golang:1.9.3

ADD . /go/src/app
WORKDIR /go/src/app
RUN go get app
RUN go install

CMD [ "/go/bin/app" ]