<h1 align="center">gnhentai</h1>

<p align="center">gnhentai — nhentai.net parser for Go.</p>
<p align="center">
    <a href="https://codecov.io/gh/tdakkota/gnhentai"><img src="https://codecov.io/gh/tdakkota/gnhentai/branch/master/graph/badge.svg"/></a>
	<a href="https://goreportcard.com/report/github.com/tdakkota/gnhentai"><img src="https://goreportcard.com/badge/github.com/tdakkota/gnhentai"></a>
	<a href="https://pkg.go.dev/github.com/tdakkota/gnhentai"><img src="https://pkg.go.dev/badge/github.com/tdakkota/gnhentai"></a>
	<a href="https://opensource.org/licenses/BSD-3-Clause"><img src="https://img.shields.io/badge/License-MIT-blue.svg"></a>
</p>

---

## Getting Started

This library is packaged using [Go modules](https://go.dev/ref/mod). You can get it via:

```
go get github.com/tdakkota/gnhentai
```

There are two implementations of `gnhentai.Client`:

- `gnhentai.API` which uses [ogen](https://github.com/ogen-go/ogen)-generated nhentai.net API
- `parser.Parser` which scrapes nhentai.net web pages

I recommend you to use API version, it is more stable and returns more data.

[Example](https://github.com/tdakkota/gnhentai/tree/master/examples/download-random-cover/main.go)

## Related

- [Unofficial API Docs](https://edgyboi2414.github.io/nhentai-api)
