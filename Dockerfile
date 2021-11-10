FROM docker.io/library/golang:alpine as builder
WORKDIR /app

COPY . .
RUN go mod download &&\
    go build -o app


# FROM docker.io/library/golang
FROM docker.io/library/alpine
WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8443

ENTRYPOINT ["./app"]
