FROM golang:1.20.3-alpine3.17 as builder

WORKDIR /usr/src/app

COPY go.mod ./

COPY . .
RUN go build -v -o /usr/local/bin/app ./

FROM alpine:3.17

COPY --from=builder /usr/local/bin/app /app
RUN chmod +x /app
CMD ["/app"]