FROM golang:1.12 AS builder

EXPOSE 3000

WORKDIR /go/src/rproxy
COPY . .

RUN go get -d -v ./...
RUN go build -o /rproxy -ldflags "-linkmode external -extldflags -static" -a *.go

# We use the distroless/static image since it includes a list of CAs and tzinfo, but is also very slim
FROM gcr.io/distroless/static:8bef63d2c8654ff89358430c7df5778162ab6027
COPY --from=builder /rproxy /rproxy
CMD ["/rproxy"]