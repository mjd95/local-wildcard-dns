To install dependencies, `go get github.com/miekg/dns`.

To run the DNS server, do
```
sudo -E bash -c 'go run dns_server.go'
```
and add the entry `nameserver 127.0.0.1` to your `/etc/resolv.conf`.

Then, to run the example, do `go run example/server/server.go` in one terminal, and
`go run example/client/client.go` in another.

See [this post](https://mjd95.github.io/2018/09/27/local-wildcard-DNS/) for motivation.
