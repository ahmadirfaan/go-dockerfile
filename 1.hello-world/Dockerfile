FROM golang:bullseye as build

WORKDIR /go/src/app
ADD ./hello.go /go/src/app/

RUN go mod init
RUN go get -d -v ./...
RUN go build -o /go/bin/app

#COPY Distroless image
FROM gcr.io/distroless/base-debian10
ENV NAMA="Ahmad Irfaan"
COPY --from=build /go/bin/app /
CMD ["/app"]

