# HTTP Proxy

This is an HTTP proxy to log requests and responses.

## Installation

### Unix users (Linux, BSDs and MacOSX)

Unix users may download and install latest *http-proxy* release with command:

```bash
sh -c "$(curl https://sweetohm.net/dist/http-proxy/install)"
```

If *curl* is not installed on you system, you might run:

```bash
sh -c "$(wget -O - https://sweetohm.net/dist/http-proxy/install)"
```

**Note:** Some directories are protected, even as *root*, on **MacOSX** (since *El Capitan* release), thus you can't install *http-proxy* in */usr/bin* for instance.

### Binary package

Otherwise, you can download latest binary at <https://github.com/c4s4/http-proxy/releases>. Get binary for your platform, move it somewhere in your *PATH* and rename it *http-proxy*.

### Go developer

If Go is installed on your machine, you can build and install *http-proxy* typing:

```
$ go install github.com/c4s4/http-proxy@latest
```

Note that in this case, version returned with `http-proxy -version` will be *UNKNOWN*.

## Usage

To run proxy, type:

```
$ http-proxy -port 8000 -addr http://127.0.0.1:8080
```

Where:

- **-port** is the port the proxy is listening (defaults to *8000*)
- **-addr** is the address the proxy is forwarding to (defaults to *http://127.0.0.1:8080*)

When performing requests on this proxy, it will print requests and responses in terminal:

```
$ http-proxy
################################################################################
#               REQUEST from 127.0.0.1:8000 at 2022-04-06T13:40:44             #
################################################################################
# REQUEST ######################################################################
GET /health HTTP/1.1
Host: 127.0.0.1:8000
Accept: */*
Accept-Encoding: gzip, deflate, br
Accept-Language: fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7
Connection: keep-alive
Referer: http://127.0.0.1:8000/docs
Sec-Ch-Ua: " Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"
Sec-Ch-Ua-Mobile: ?0
Sec-Ch-Ua-Platform: "Linux"
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: same-origin
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36
# RESPONSE #####################################################################
HTTP/1.1 200 OK
Content-Length: 15
Content-Type: application/json; charset=utf-8
Date: Wed, 06 Apr 2022 13:40:44 GMT

{"status":"ok"}
```

This is very handy to debug HTTP. Note that these logs are not exactly request and response, as described in [this documentation](https://pkg.go.dev/net/http/httputil#DumpRequest).

You can print help with `http-proxy -help` and version with `http-proxy -version`.

*Enjoy!*
