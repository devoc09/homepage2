FROM golang:1.17.5-alpine3.15 as builder

WORKDIR /homepage2

ADD . /homepage2
RUN go get
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o webserver .

FROM alpine:latest
RUN apk --no-cache add ca-certificates && mkdir /var/www/
EXPOSE 80 443
WORKDIR /root/homepage2
ADD . /root/homepage2
COPY --from=builder /homepage2/webserver .
CMD ["./webserver"]
