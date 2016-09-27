# Rich Content Info (URL info library)

[![Build Status](https://travis-ci.org/mdigger/urlinfo.svg?branch=master)](https://travis-ci.org/mdigger/urlinfo)
[![GoDoc](https://godoc.org/github.com/mdigger/urlinfo?status.svg)](https://godoc.org/github.com/mdigger/urlinfo)
[![Coverage Status](https://coveralls.io/repos/github/mdigger/urlinfo/badge.svg?branch=master)](https://coveralls.io/github/mdigger/urlinfo?branch=master)

```go
package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/mdigger/urlinfo"
)

func main() {
	info := urlinfo.Get("http://mdigger.tumblr.com")
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(info)
}

```

Output:

```json
{
    "url": "http://mdigger.tumblr.com",
    "status": 200,
    "mediaType": "text/html; charset=UTF-8",
    "title": "Копилка ненужных вещей",
    "description": "У каждого человека должна быть своя \"копилка ненужных вещей\", которой он с радостью готов поделиться с другими. Вы как раз в таком месте...",
    "image": "http://66.media.tumblr.com/avatar_d1b501c5ba67_128.png"
}
```