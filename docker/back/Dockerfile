FROM golang:1.14.4-stretch as builder
ENV GO111MODULE=on
WORKDIR /build
RUN go get -u github.com/kisielk/errcheck
COPY ./back/go.mod ./back/go.sum ./
RUN go mod download
COPY ./back ./
RUN errcheck ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o back ./main.go

FROM alpine:latest
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*
WORKDIR /root/
COPY --from=builder /build/back .
ENV GIN_MODE=release
CMD ["./back"]
