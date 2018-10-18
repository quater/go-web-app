FROM golang:1.11

WORKDIR /go/src/app
COPY . .

RUN go get github.com/appleboy/gofight
RUN go build && go test

EXPOSE 8484

CMD ["/go/src/app/app"]
