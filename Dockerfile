FROM golang:1.13.0-alpine3.10 AS builder

WORKDIR /app
COPY . .
ARG GO111MODULE=on
RUN go build -o goproxy

FROM alpine:3.10 AS prod
COPY --from=builder /app/goproxy .
EXPOSE 8080
VOLUME ["/cache"]
CMD ["./goproxy", "-dir", "/cache", "-listen", ":8080"]
