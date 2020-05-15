FROM golang:buster as builder

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s' -o chatshare main.go

FROM alpine as runner

COPY --from=builder /src/chatshare /usr/bin/chatshare
COPY server.cer .
COPY server.key .
EXPOSE 8080

CMD ["/usr/bin/chatshare"]