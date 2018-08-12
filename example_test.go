package urlinfo_test

import (
	"encoding/json"
	"os"

	"github.com/mdigger/urlinfo"
)

func Example() {
	info := urlinfo.Get("https://edition.cnn.com/2018/08/11/politics/donald-trump-vacation-new-jersey/index.html")
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(info)
	// Output:
	// {
	// 	"url": "https://edition.cnn.com/2018/08/11/politics/donald-trump-vacation-new-jersey/index.html",
	// 	"status": 200,
	// 	"mediaType": "text/html; charset=utf-8",
	// 	"title": "Peek inside Trump's total non-vacation in New Jersey",
	// 	"description": "For President Donald Trump, the long-standing image of a leader in repose doesn't apply, at least in his own mind. In New Jersey this week, events were arranged to demonstrate his extended stay at a wooded golf resort was, in fact, work.",
	// 	"author": "Kevin Liptak and Jeff Zeleny, CNN",
	// 	"image": "https://cdn.cnn.com/cnnnext/dam/assets/170611140113-trump-golfing-super-tease.jpg",
	// 	"keywords": [
	// 		"politics"
	// 	],
	// 	"type": "article",
	// 	"site": "CNN"
	// }
}
