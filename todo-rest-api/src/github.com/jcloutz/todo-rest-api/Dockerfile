FROM golang:1.9.3
WORKDIR /go/src/github.com/jcloutz/todo-rest-api
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/jcloutz/todo-rest-api/app .
CMD ["./app"]
