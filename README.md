<h1 align="center">gnhentai</h1>

<p align="center">gnhentai — nhentai.net parser for Go.</p>
<p align="center">
    <a href="https://travis-ci.org/github/tdakkota/gnhentai"><img src="https://travis-ci.org/tdakkota/gnhentai.svg?branch=master"></a>
    <a href="https://codecov.io/gh/tdakkota/gnhentai"><img src="https://codecov.io/gh/tdakkota/gnhentai/branch/master/graph/badge.svg" /> </a>
	<a href="https://goreportcard.com/report/github.com/tdakkota/gnhentai"><img src="https://goreportcard.com/badge/github.com/tdakkota/gnhentai"></a>
	<a href="https://www.codefactor.io/repository/github/tdakkota/gnhentai"><img src="https://www.codefactor.io/repository/github/tdakkota/gnhentai/badge"></a>
	<a href="https://godoc.org/github.com/tdakkota/gnhentai"><img src="https://godoc.org/github.com/tdakkota/gnhentai?status.svg"></a>
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
```go
package main

import (
	"fmt"
	"github.com/tdakkota/gnhentai"
	"github.com/tdakkota/gnhentai/api"
	"io"
	"os"
)

func main() {
	c := api.NewClient()

	doujinshi, err := c.Random()
	if err != nil {
		panic(err)
	}

	fmt.Println("Downloading", doujinshi.Name())
	fmt.Println("Tags:")
	for _, tag := range doujinshi.Tags {
		fmt.Println(" - ", tag.Name)
	}

	format := gnhentai.FormatFromImage(doujinshi.Images.Cover)
	cover, err := c.Cover(doujinshi.MediaID, format)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(fmt.Sprintf("cover_%d.%s", doujinshi.MediaID, format))
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, cover)
	if err != nil {
		panic(err)
	}
}

```

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