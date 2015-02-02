// Библиотека для работы с Yandex Rich Content API
// (https://tech.yandex.ru/rca/)
//
// Rich Content API предоставляет доступ к контентной системе Яндекса. В ней хранятся десятки
// миллиардов страниц, и для любой из них можно получить сниппет, заголовок и URL в каноническом виде,
// полный текст, а также список ссылок на размещённые изображения и видео.
//
// Например, API будет полезен для сервисов, чьи пользователи обмениваются ссылками. Благодаря
// Rich Content API они смогут видеть превью любой веб-страницы, не покидая сервис.
package yrca

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const yandexRCAUrl = "http://rca.yandex.com/" // ссылка на сервис

// Client описывает клиента для запросов к Yandex Rich Content API.
type Client struct {
	url    string // Yandex Rich Content API URL with API Key
	client *http.Client
}

// NewClient возвращает инициализированный Client для запросов к Yandex Rich Content API.
// В качестве обязательного параметра необходимо передать ключ для доступа к API, который
// вы получили на сервере Yandex (https://tech.yandex.ru/key/form.xml?service=rca).
func NewClient(apiKey string) *Client {
	return &Client{
		url: yandexRCAUrl + "?key=" + apiKey + "&url=",
		client: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
	}
}

// Get осуществляет обращение к Yandex Rich Content API и возвращает ответ.
func (c *Client) Get(url string) (ctx *Response, err error) {
	var httpresp *http.Response
	httpresp, err = c.client.Get(c.url + url)
	if err != nil {
		return
	}
	switch httpresp.StatusCode {
	case http.StatusUnauthorized: // 401 Unauthorized - Невалидный ключ.
		err = &Error{
			Code:    httpresp.StatusCode,
			Message: "invalid API key",
		}
		return
	case http.StatusForbidden: // 403 Forbidden - Отсутствует обязательный параметр key.
		err = &Error{
			Code:    httpresp.StatusCode,
			Message: "no API key",
		}
	default:
		ctx = new(Response)
		err = json.NewDecoder(httpresp.Body).Decode(ctx)
		httpresp.Body.Close()
	}
	return
}

// GetFull возвращает полный текст страницы вместе с ссылками на изображения.
func (c *Client) GetFull(url string) (ctx *Response, err error) {
	return c.Get(url + "&full=1")
}

// Error описывает формат ошибки, возвращаемой сервисом.
type Error struct {
	Type    string `json:"error_type,omitempty"`    // internal или external
	Code    int    `json:"error_code,string"`       // код ошибки (HTTP)
	Url     string `json:"url,omitempty"`           // запрашиваемый адрес или адрес с ошибкой
	Message string `json:"error_message,omitempty"` // описание ошибки
}

// Error возвращает текстовое описание ошибки.
func (e *Error) Error() string {
	return fmt.Sprintf("yrca [%d]: %s", e.Code, e.Message)
}

// Response описывает формат ответа, возвращаемый сервисом.
type Response struct {
	Url        string   `json:"url"`               // Адрес страницы, извлеченный из запроса.
	FinalUrl   string   `json:"finalurl"`          // Адрес страницы в каноническом виде.
	Title      string   `json:"title,omitempty"`   // Заголовок страницы.
	Content    string   `json:"content,omitempty"` // Краткая аннотация страницы (сниппет) или ее полный текст.
	Img        []string `json:"img,omitempty"`     // Список ссылок на основные или все изображения на странице.
	Video      []*Video `json:"video,omitempty"`   // Список пар (url, duration) для основных видеороликов на странице.
	Mime       string   `json:"mime,omitempty"`    // MIME-тип страницы.
	Confidence struct { // Степень уверенности в качестве выбора: high, medium, low.
		Img     string `json:"img,omitempty"`     // изображений
		Content string `json:"content,omitempty"` // текстовой аннотации страницы
	} `json:"confidence,omitempty"`
}

// Video описывает информацию о видео-файлах, найденных по запрашиваемому адресу.
type Video struct {
	Duration int    `json:"duration,omitempty"` // продолжительность
	Url      string `json:"url,omitempty"`      // ссылка для загрузки
}
