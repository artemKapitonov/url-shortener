package convertor

import (
	"github.com/artemKapitonov/url-shortener/internal/entity"
	url_shortener_v1 "github.com/artemKapitonov/url-shortener/pkg/url-shortener_v1"
	"sync"
)

type EntityConvertor struct {
	r sync.RWMutex
}

func New() *EntityConvertor {
	return &EntityConvertor{}
}

func (c *EntityConvertor) Convert(url *url_shortener_v1.ShortURL) entity.ShortURL {
	return entity.ShortURL{Url: "http.localhost:8080"}
}
