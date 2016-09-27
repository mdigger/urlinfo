package urlinfo_test

import (
	"encoding/json"
	"os"

	"github.com/mdigger/urlinfo"
)

func Example() {
	info := urlinfo.Get("http://mdigger.tumblr.com")
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(info)
	// Output:
	// {
	//     "url": "http://mdigger.tumblr.com",
	//     "status": 200,
	//     "mediaType": "text/html; charset=UTF-8",
	//     "title": "Копилка ненужных вещей",
	//     "description": "У каждого человека должна быть своя \"копилка ненужных вещей\", которой он с радостью готов поделиться с другими. Вы как раз в таком месте...",
	//     "image": "http://66.media.tumblr.com/avatar_d1b501c5ba67_128.png"
	// }
}
