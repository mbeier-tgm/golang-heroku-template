FROM golang:1.15

WORKDIR /go/src/go-webapp
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-webapp"]