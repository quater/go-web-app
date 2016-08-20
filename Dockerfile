FROM golang:1.6-onbuild

RUN go get github.com/appleboy/gofight && go get gopkg.in/appleboy/gofight.v1

RUN go build && go test

EXPOSE 8484

CMD ["app"]
