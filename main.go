package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nstratos/go-myanimelist/mal"
)

func main() {
	publicInfoClient := &http.Client{
		// Create client ID from https://myanimelist.net/apiconfig.
		Transport: &clientIDTransport{ClientID: "<ClientID>"},
	}

	c := mal.NewClient(publicInfoClient)

	ctx := context.Background()

	animeDetailsCtx, animeDetailsID, animeDetailsOption := c.Anime.Details(ctx, 51009)

	fmt.Println(animeDetailsCtx, animeDetailsID, animeDetailsOption)
}

type clientIDTransport struct {
	Transport http.RoundTripper
	ClientID  string
}

func (c *clientIDTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.Transport == nil {
		c.Transport = http.DefaultTransport
	}
	req.Header.Add("X-MAL-CLIENT-ID", c.ClientID)
	return c.Transport.RoundTrip(req)
}
