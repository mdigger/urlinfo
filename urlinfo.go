package urlinfo

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var (
	// UserAgent используется как строка с указанием названия веб-браузера.
	UserAgent = "mdigger/2.0"
	// ParseLimit накладывает ограничение на размер анализируемых данных. Слишком
	// большой размер избыточен, т.к. мы анализируем только заголовок HTML.
	ParseLimit = int64(64 << 10)
	// KeywordMaxLength ограничивает максимально допустимую длину ключевой фразы.
	KeywordMaxLength = 32
)

// Get возвращает информацию с описание указанного URL. Для запроса используется
// http.DefaultClient.
func Get(rawurl string) *Info {
	// проверяем, что URL правильный и содержит полный путь.
	purl, err := url.Parse(rawurl)
	if err != nil || !purl.IsAbs() || purl.Host == "" {
		return nil
	}
	req, err := http.NewRequest("GET", rawurl, nil)
	if UserAgent != "" {
		req.Header.Set("User-Agent", UserAgent)
	}
	// указываем, что соединение должно быть закрыто по окончании
	req.Close = true
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &Info{
			URL:         rawurl,
			Status:      http.StatusServiceUnavailable,
			Description: err.Error(),
		}
	}
	// разбираем полученный ответ
	info := Parse(resp)                        // разбираем ответ
	io.CopyN(ioutil.Discard, resp.Body, 2<<10) // пропускаем остальное
	resp.Body.Close()                          // все закрываем
	return info                                // возвращаем информацию
}

// Info описывает информацию, возвращаемую после разбора ответа на HTTP-запрос.
type Info struct {
	URL         string   `json:"url"`                   // ссылка
	Status      int      `json:"status"`                // статус ответа при обращении
	Length      int64    `json:"length,omitempty"`      // длина ответа
	ContentType string   `json:"mediaType,omitempty"`   // content-type
	Title       string   `json:"title,omitempty"`       // заголовок
	Description string   `json:"description,omitempty"` // краткое описание
	Image       string   `json:"image,omitempty"`       // иллюстрация
	Keywords    []string `json:"keywords,omitempty"`    // список ключевых слов
	Type        string   `json:"type,omitempty"`        // тип содержимого
	Video       string   `json:"video,omitempty"`       // ссылка на видео
	Locale      string   `json:"locale,omitempty"`      // язык
	Site        string   `json:"site,omitempty"`        // название сайта
}

// Parse разбирает ответ на HTTP-запрос и на основании него отдает краткую
// сводную информацию о странице. В основном, конечно, она будет интересна
// и разнообразна для HTML-страниц, которые поддерживают теги для разных
// социальных сетей, типа Facebook и Twitter.
func Parse(resp *http.Response) *Info {
	// заполняем основные поля из ответа сервера
	var info = &Info{
		URL:         resp.Request.URL.String(),
		Status:      resp.StatusCode,
		Length:      resp.ContentLength,
		ContentType: resp.Header.Get("Content-Type"),
	}
	// если сервер не отдавал длину ответа, то она будет -1
	if info.Length == -1 {
		info.Length = 0
	}
	// проверяем, что это HTML-страница, и что ответ имеет нормальный код,
	// чтобы не разбирать страницы с HTML-описанием ошибок и заглушек сайтов
	if info.Status == 200 && strings.HasPrefix(info.ContentType, "text/html") {
		// разбираем заголовок HTML
		info.parse(resp.Body)
		// хоть и редко, но бывает что картинка указана с относительным путем
		if info.Image != "" {
			// разбираем URL в контексте URL запроса
			if url, err := resp.Request.URL.Parse(info.Image); err == nil {
				info.Image = url.String()
			}
		}
	}
	return info
}

// parse разбирает непосредственно HTML ответ страницы и заполняет поля, которые
// находит в заголовке ответа.
func (i *Info) parse(r io.Reader) {
	doc := html.NewTokenizer(io.LimitReader(r, ParseLimit))
	var (
		unique             map[string]struct{} // список уникальных ключевых слов
		list               []string            // реальный список ключевых слов
		none               = struct{}{}        // используется для списка уникальных слов
		description, title bool                // флаг, что описание и заголовок уже добавлены
	)
	for {
		switch doc.Next() {
		case html.ErrorToken:
			return // больше разбирать нечего
		case html.StartTagToken, html.SelfClosingTagToken:
			break // это наш случай - разбираем
		default:
			continue // это нам не интересно
		}
		// разбираем по типам тегов HTML
		switch token := doc.Token(); token.DataAtom {
		case atom.Body:
			return // дошли до тела HTML - дальше делать нечего
		case atom.Title:
			// стандартный заголовок документа
			if doc.Next() == html.TextToken {
				i.Title = string(doc.Text())
			}
		case atom.Meta:
			// самая интересная нам часть
			// бывает два варианта:
			// 	<meta name="keywords" content="юмор,цитаты,чиполлино,интроверт,еда">
			// 	<meta property="og:title" content="Копилка ненужных вещей">
			var name, content string
			for _, attr := range token.Attr {
				switch attr.Key {
				case "name", "property":
					name = strings.TrimSpace(attr.Val)
				case "content":
					content = strings.TrimSpace(attr.Val)
				}
			}
			// игнорируем все пустышки
			if len(name) == 0 {
				continue
			}
			// теперь смотрим на то, что же за данные мы нашли
			switch name {
			case "title", "og:title":
				i.Title = content
				title = true
			case "twitter:title":
				// эти заголовки бывают часто обрезанными, поэтому используем их
				// только в тех случаях, когда больше не из чего выбирать
				if i.Title == "" && !title {
					i.Title = content
				}
			case "og:description":
				i.Description = content
				description = true // взводим флаг, чтобы не перезаписывать в twitter
			case "description", "twitter:description":
				// очень часто там, где описание намеренно сброшено, для
				// twitter оно заполнено, например, дублирование заголовка
				if i.Description == "" && !description {
					i.Description = content
				}
			case "image", "og:image":
				i.Image = content
			case "twitter:image", "twitter:image:src":
				// здесь картинки могут быть меньшего размера
				if i.Image == "" {
					i.Image = content
				}
			case "keywords":
				// разделяем все на отдельные ключевые слова
				keywords := strings.Split(content, ",")
				if unique == nil || list == nil {
					unique = make(map[string]struct{}, len(keywords))
					list = make([]string, 0, len(keywords))
				}
				for _, keyword := range keywords {
					keyword = strings.TrimSpace(keyword)
					// далее, отсекаем дублирующиеся ключевые слова
					if _, ok := unique[keyword]; ok ||
						keyword == "" ||
						len(keyword) > KeywordMaxLength {
						continue
					}
					unique[keyword] = none
					// сохраняем очередность как в оригинале
					list = append(list, keyword)
				}
				i.Keywords = list
			case "article:tag", "video:tag", "og:video:tag":
				if unique == nil || list == nil {
					unique = make(map[string]struct{})
					list = make([]string, 0, 20)
				}
				keyword := strings.TrimSpace(content)
				if _, ok := unique[keyword]; ok ||
					keyword == "" ||
					len(keyword) > KeywordMaxLength {
					continue
				}
				unique[keyword] = none
				list = append(list, keyword)
				i.Keywords = list

			case "og:type":
				i.Type = content
			case "og:video", "og:video:url", "twitter:player":
				// оставляем первый найденный вариант, потому что он обычно лучший
				if i.Video == "" {
					i.Video = content
				}
			case "og:locale":
				i.Locale = content
			case "og:site_name":
				i.Site = content
			case "twitter:site":
				// это обычно название компании в Twitter
				if i.Site == "" {
					i.Site = content
				}
			}
		}
	}
}
