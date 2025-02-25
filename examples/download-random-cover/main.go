package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"github.com/tdakkota/gnhentai"
)

func downloadFile(ctx context.Context, client *http.Client, srcURL, destPath string) error {
	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = dest.Close()
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, srcURL, http.NoBody)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	if _, err := io.Copy(dest, resp.Body); err != nil {
		return err
	}

	return nil
}

func run(ctx context.Context) error {
	var (
		client = http.DefaultClient
		nh     = gnhentai.NewAPI(gnhentai.APIOptions{
			Client: client,
		})
	)

	doujinshi, err := nh.Random(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Downloading", doujinshi.Name())
	fmt.Println("Tags:")
	for _, tag := range doujinshi.Tags {
		fmt.Println(" - ", tag.Name)
	}

	cover, ok := doujinshi.Images.Cover.Get()
	if !ok {
		fmt.Println("No cover found")
		return nil
	}

	ext, err := cover.Ext()
	if err != nil {
		return err
	}

	link, err := doujinshi.CoverLink()
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("cover_%s.%s", doujinshi.ID.ToString(), ext)
	fmt.Println("Downloading cover:", filename)

	return downloadFile(ctx, client, link, filename)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		os.Exit(1)
	}
}
