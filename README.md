
## Build

Fetch newest skipper library.

```
% go get github.com/zalando/skipper
go: downloading github.com/zalando/skipper v0.13.218
...
```

Build your custom skipper

```
% make
mkdir -p bin
go build -ldflags "-X main.version=0beb7bd -X main.commit=0beb7bd" -o bin/skipper .
```

## Add filter

[custom Filter code](./filters/custom.go)

## Run

Run your customer skipper with `myFilter()`.

```
% ./bin/skipper -inline-routes='* -> myFilter() -> status(200) -> <shunt>'
[APP]INFO[0000] Expose metrics in codahale format
[APP]INFO[0000] support listener on :9911
[APP]INFO[0000] proxy listener on :9090
[APP]INFO[0000] TLS settings not found, defaulting to HTTP
[APP]INFO[0000] route settings, reset, route: : * -> myFilter() -> status(200) -> <shunt>        <--- custom filter loaded
[APP]INFO[0000] route settings received
[APP]INFO[0000] route settings applied
::1 - - [20/Jun/2022:15:45:07 +0200] "GET /goo HTTP/1.1" 200 0 "-" "curl/7.49.0" 0 localhost:9090 - -
^C

% curl localhost:9090/foo -v
*   Trying ::1...
* Connected to localhost (::1) port 9090 (#0)
> GET /foo HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.49.0
> Accept: */*
>
< HTTP/1.1 200 OK
< My-Filter: response                      <-------- our response filter
< Server: Skipper
< Date: Mon, 20 Jun 2022 13:45:07 GMT
< Transfer-Encoding: chunked
<
* Connection #0 to host localhost left intact
```
