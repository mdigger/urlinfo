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
		"http://ya.ru",
		"https://godoc.org/golang.org/x/net/html",
		"https://github.com/mdigger/yrca/blob/master/client.go",
		"https://www.tumblr.com/dashboard",
		"http://mdigger.tumblr.com",
		"https://65.media.tumblr.com/e2906df46bf289df90625afdf625c187/tumblr_oe23mgNqLX1qazdu9o1_1280.jpg",
		"http://lenta.ru",
		"http://www.livejournal.com/media/843446.html",
	}
	for _, url := range urls {
		info, err := Get(url)
		if err != nil {
			t.Error(err)
		}
		enc.Encode(url)
		enc.Encode(info)
	}
}
