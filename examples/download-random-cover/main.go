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
