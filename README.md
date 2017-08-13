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
	info := urlinfo.Get("https://www.youtube.com/watch?v=OHegEgC8uwY")
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(info)
}

```

Output:

```json
{
  "url": "https://www.youtube.com/watch?v=OHegEgC8uwY",
  "status": 200,
  "mediaType": "text/html; charset=utf-8",
  "title": "БГ - Время наебениться",
  "description": "Позвольте мне прервать ваши вечные споры, Позвольте расшатать скрепы и оковы. Время беспощадно, оно как волчица, Вот мы сидим здесь, а оно мчится. Ох бы жить...",
  "image": "https://i.ytimg.com/vi/OHegEgC8uwY/hqdefault.jpg",
  "keywords": [
    "БГ",
    "время",
    "наебениться",
    "Степанцов",
    "bg",
    "Аквариум",
    "бг аудио",
    "бг филми",
    "филми",
    "филм бг аудио",
    "филми бг аудио",
    "bg audio",
    "бг...",
    "бг субтитри",
    "филми бг субтитри",
    "бг музика",
    "бг песни",
    "маша и мечока",
    "мики маус",
    "гребенщиков",
    "аквариум",
    "город золотой"
  ],
  "type": "video",
  "video": "https://www.youtube.com/embed/OHegEgC8uwY",
  "site": "YouTube"
}
```