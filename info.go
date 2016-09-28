// Package urlinfo allows you to obtain summary information about the content at
// the specified URL.
package urlinfo

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Info describes the information located at the specified URL.
type Info struct {
	URL         string `json:"url"`
	Status      int    `json:"status"`
	Length      int64  `json:"length,omitempty"`
	Type        string `json:"mediaType,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
}

func (i *Info) parseHTML(r io.Reader) {
	if i.Status != 200 || !strings.HasPrefix(i.Type, "text/html") {
		return
	}
	doc := html.NewTokenizer(io.LimitReader(r, ParseLimit))
parse:
	for {
		switch doc.Next() {
		case html.ErrorToken:
			break parse
		case html.StartTagToken, html.SelfClosingTagToken:
			break
		default:
			continue
		}

		switch token := doc.Token(); token.DataAtom {
		case atom.Body:
			break parse // end parse on body
		case atom.Title:
			if doc.Next() == html.TextToken {
				i.Title = string(doc.Text())
			}
		case atom.Meta:
			var name, content string
			for _, attr := range token.Attr {
				switch attr.Key {
				case "name", "property":
					name = attr.Val
				case "content":
					content = attr.Val
				}
			}
			// remove the possible prefixes in the name
			if indx := strings.LastIndexByte(name, byte(':')); indx > 0 {
				name = name[indx+1:]
			}
			switch name {
			case "description":
				if i.Description == "" {
					i.Description = content
				}
			case "title":
				if i.Title == "" {
					i.Title = content
				}
			case "image":
				if i.Image == "" {
					i.Image = content
				}
			}
		}
	}
}

// NewInfo returns information about the contents of the URL in response to the
// HTTP request.
func NewInfo(resp *http.Response) *Info {
	var info = &Info{
		Status: resp.StatusCode,
		Length: resp.ContentLength,
		URL:    resp.Request.URL.String(),
		Type:   resp.Header.Get("Content-Type"),
	}
	if info.Length == -1 {
		info.Length = 0
	}
	info.parseHTML(resp.Body) // parse HTML
	// parse image URL & bringing it to an absolute
	if info.Image != "" {
		url, err := resp.Request.URL.Parse(info.Image)
		if err == nil {
			info.Image = url.String()
		}
	}
	return info
}

var (
	UserAgent  = "mdigger/2.0"
	ParseLimit = int64(64 << 10)
)
var httpClient = &http.Client{
	Timeout: time.Second * 60,
}

// Get возвращает краткое описание информации, возвращаемое по указанному URL.
func Get(urlstr string) *Info {
	// проверяем, что URL правильный и содержит полный путь.
	purl, err := url.Parse(urlstr)
	if err != nil || !purl.IsAbs() || purl.Host == "" {
		return nil
	}
	req, err := http.NewRequest(http.MethodGet, urlstr, nil)
	if UserAgent != "" {
		req.Header.Set("User-Agent", UserAgent)
	}
	req.Close = true // close connection, not to produce open
	resp, err := httpClient.Do(req)
	if err != nil {
		return &Info{
			URL:         urlstr,
			Status:      http.StatusServiceUnavailable,
			Description: err.Error(),
		}
	}
	info := NewInfo(resp)
	io.CopyN(ioutil.Discard, resp.Body, 2<<10) // skip full body response
	resp.Body.Close()
	return info
}
