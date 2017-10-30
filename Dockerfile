FROM golang:1.9

WORKDIR /go/src/app

COPY . .

RUN mkdir -p /go/src/github.com/plutov/
RUN ln -s /go/src/app/ /go/src/github.com/plutov/go-snake-telnet

RUN go-wrapper install

ENTRYPOINT ["go-wrapper", "run"]
CMD []
