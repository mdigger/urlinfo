package yrca

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ВНИМАНИЕ! Воспользуйтесь своим ключем!
const apiKey = "xx.xxxxxx.xxxx.xxxxxxx"

func TestYRCA(t *testing.T) {
	client := NewClient(apiKey)
	for _, url := range []string{
		"http://mdigger.tumblr.com/",
		"http://facebook.com/",
		"http://yandex.ru/",
		"http://golang.org/",
	} {
		ctx, err := client.GetFull(url)
		if err != nil {
			t.Error(err)
		}
		data, err := json.MarshalIndent(ctx, "", "\t")
		if err != nil {
			t.Error(err)
		}
		fmt.Println(string(data))
	}
}

func Example() {
	client := NewClient(apiKey)
	ctx, err := client.Get("http://golang.org/")
	if err != nil {
		panic(err)
	}
	fmt.Println(ctx.FinalURL)
	fmt.Print(ctx.Title)
	// Output:
	// http://golang.org/
	// The Go Programming Language
}
