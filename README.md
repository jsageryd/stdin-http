# stdin-http

[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/jsageryd/stdin-http#license)

This tool slurps data from STDIN into memory and then serves that data for any
HTTP request made to it.

## Installation
```
$ go get -u github.com/jsageryd/stdin-http
```

## Example
```
$ echo 'Hello, World o/' | stdin-http -p 1234
2018/05/25 10:06:18 Accepting data from stdin...
2018/05/25 10:06:18 Serving 16 bytes at :1234...
2018/05/25 10:06:30 [::1]:56189 GET /
```

```
$ http get :1234
HTTP/1.1 200 OK
Accept-Ranges: bytes
Content-Length: 16
Content-Type: text/plain; charset=utf-8
Date: Fri, 25 May 2018 08:06:30 GMT
Last-Modified: Fri, 25 May 2018 08:06:30 GMT

Hello, World o/

```

## License
Copyright (c) 2018 Johan Sageryd <j@1616.se>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
