# rproxy

rproxy is a dead-simple reverse proxy. You give it a host to forward to and it will forward all requests it receives to that host. No configuration, no tinkering, no BS.

## Usage

Set `RPROXY_TO_HOST` to the host you want to forward. For example, `https://example.com`, or `http://192.168.1.12:3000`. Since this uses Go's `NewSingleHostReverseProxy`, you can use anything it supports. Check the docs [here](https://golang.org/pkg/net/http/httputil/#NewSingleHostReverseProxy).

It listens on port 3000. This is not configurable.

## Docker Image

`tjhorner/rproxy`

## License

MIT