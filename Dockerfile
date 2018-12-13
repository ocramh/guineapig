# build
FROM golang:alpine as builder

WORKDIR /go/src/github.com/ocramh/guineapig
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o guineapig

# release
FROM alpine:3.8
COPY --from=builder /go/src/github.com/ocramh/guineapig/guineapig .
EXPOSE 8080
ENTRYPOINT ./guineapig
