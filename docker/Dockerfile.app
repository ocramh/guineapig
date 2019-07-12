# build
FROM golang:alpine as builder
RUN apk update && apk upgrade && \
    apk add --no-cache \
      bash \
      git \ 
      openssh

WORKDIR /go/src/github.com/ocramh/guineapig
COPY . .
ENV GO111MODULE=on 
RUN go mod tidy  
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o guineapig

# release
FROM alpine:3.8
COPY --from=builder /go/src/github.com/ocramh/guineapig/guineapig .
EXPOSE 8080
ENTRYPOINT ./guineapig
 