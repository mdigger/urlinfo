package urlinfo

import (
	"encoding/json"
	"os"
	"testing"
)

func TestRCA(t *testing.T) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	urls := []string{
		"http://appleinsider.com/articles/16/09/24/review-apple-watch-series-2-is-a-great-improvement-but-watchos-3-steals-the-show",
		"http://ya.ru",
		"https://godoc.org/golang.org/x/net/html",
		"https://github.com/mdigger/yrca/blob/master/client.go",
		"https://www.tumblr.com/dashboard",
		"http://mdigger.tumblr.com",
		"https://65.media.tumblr.com/e2906df46bf289df90625afdf625c187/tumblr_oe23mgNqLX1qazdu9o1_1280.jpg",
		"http://lenta.ru",
		"http://www.livejournal.com/media/843446.html",
		"http://localhost:8099/",
		"http://flibusta.net",
	}
	for _, url := range urls {
		enc.Encode(url)
		enc.Encode(Get(url))
	}
}
