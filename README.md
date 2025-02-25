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

This library is packaged using [Go modules][go-modules]. You can get it via:

```
go get github.com/tdakkota/gnhentai
```

## Use as lib

There are two implementations of `gnhentai.Client`:

- `api.Client` which uses NHentai API
- `parser.Parser` which parses NHentai web pages using `goquery`

I recommend you to use API version, it is more stable.

[Example](https://github.com/tdakkota/gnhentai/tree/master/examples/download-random-cover/main.go)

## gnhentai-cli

Install and run (`GOBIN` should be in `PATH`), it will download random book into current dir

```
gnhentai-cli download
```

or

```
gnhentai-cli download --id=<your_manga_id>
```

## Use API server

```
gnhentai-server run --bind=<bind_addr, default is :8080>
```

## Related

- [Swagger file](https://gist.github.com/tdakkota/6efa100de2000549027617b1a1088d78)
- [Unofficial API Docs](https://edgyboi2414.github.io/nhentai-api)

[go-modules]: https://github.com/golang/go/wiki/Modules
