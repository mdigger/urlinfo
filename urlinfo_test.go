package urlinfo

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

var urls = []string{
	"https://www.youtube.com/watch?v=OHegEgC8uwY",
	"http://appleinsider.com/articles/16/09/24/review-apple-watch-series-2-is-a-great-improvement-but-watchos-3-steals-the-show",
	"http://ya.ru",
	"https://godoc.org/golang.org/x/net/html",
	"https://github.com/mdigger/yrca/blob/master/client.go",
	"https://www.tumblr.com/dashboard",
	"http://mdigger.tumblr.com",
	"http://lenta.ru",
	"http://www.livejournal.com/media/843446.html",
	"http://flibusta.net",

	"https://changelog.com/gotime/52",
	"https://www.youtube.com/watch?v=dYrYCt2dTkw",
	"https://ru.insider.pro/technologies/2017-08-12/kak-dva-brata-prevratili-sem-strok-koda-v-startap-stoimostyu-92-mlrd/",
	"http://mi3ch.livejournal.com/3859447.html",
	"http://mi3ch.livejournal.com/3858977.html",
	"http://fishki.net/2356648-muzej-slavjanskoj-kulytury-im-konstantina-vasilyeva.html",
	"http://kak-eto-sdelano.livejournal.com/682219.html",
	"http://ibigdan.livejournal.com/20762448.html",
	"https://vc.ru/p/travis-investors-fight?from=rss",
	"https://hi-news.ru/entertainment/ii-algoritm-pobedil-odnogo-iz-luchshix-v-mire-igrokov-v-dota-2.html",
	"https://www.anekdot.ru/id/901123/",
	"https://lifehacker.ru/2017/08/12/music-quiz-2/",
	"http://www.prostomac.com/2017/08/apple-xochet-dat-sotrudnikam-vozmozhnost-rabotat-iz-doma/",

	"http://mdigger.tumblr.com/post/163937210475/самой-экстремистской-книгой-является-сказка",
	"http://mdigger.tumblr.com/post/163908327000/спасение-интровертов-дело-рук-самих-интровертов",
	"http://mdigger.tumblr.com/tagged/юмор",
	"http://mdigger.tumblr.com/post/163714356225/то-неловкое-чувство-когда-дочь-твоего-начальника",
	"http://mdigger.tumblr.com/post/163640518165/чего-хотят-женщины",

	"http://mdigger.tumblr.com/post/163603836585/альтернативную-энергетику-похоронили-еще-40-лет",
	"http://mdigger.tumblr.com/post/163209666740",
	"http://mdigger.tumblr.com/post/162586203000/собеседование",
	"http://mdigger.tumblr.com/post/162468560805/ты-за-что-сидишь-не-смог-дозвониться-по",
	"http://mdigger.tumblr.com/post/162012991890/как-создавались-спецэффекты-в-эпоху-немого-кино",
	"http://mdigger.tumblr.com/post/162004835065/15-французских-фраз-для-секса-которые-вы-не",
	"http://mdigger.tumblr.com/post/161943272635/булгаков-в-кино-каким-вы-его-не-видели",
	"http://mdigger.tumblr.com/post/161135192660/дореволюціонный-совѣтчикъ-я-матросъ",
	"http://mdigger.tumblr.com/post/161129428905/шекспир-и-ук-рф",
	"http://mdigger.tumblr.com/post/160943645600/geek",
	"http://mdigger.tumblr.com/post/160229340310/кастинг-короткометражка",

	"https://www.youtube.com/watch?v=WpBpbCF2dtw",

	"http://arzamas.academy/materials/1175",
	"http://5respublika.com/kultura/sex-expressions-francais.html",
	"https://twizz.ru/kak-sozdavalis-speceffekty-v-epoxu-nemogo-kino/",
	"http://rusogolik.livejournal.com/409.html",
	"http://masterok.livejournal.com/3781902.html",
	"https://ria.ru/religion/20170519/1494628490.html",
	"https://rns.online/economy/Gref-shkolnoe-obrazovanie-v-Rossii-ostalos-na-urovne-XIX-veka--2017-05-19/",
}

func TestParse(t *testing.T) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	for _, url := range urls {
		enc.Encode(Get(url))
		fmt.Println(strings.Repeat("-", 80))
	}
}
