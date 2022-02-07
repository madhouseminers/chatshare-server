FROM golang:buster as builder

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s' -o chatshare cmd/chatshare/main.go

FROM scratch as runner

COPY --from=builder /src/chatshare /usr/bin/chatshare
COPY --from=builder /etc/ssl/certs ./etc/ssl/certs
EXPOSE 8080

CMD ["/usr/bin/chatshare"]