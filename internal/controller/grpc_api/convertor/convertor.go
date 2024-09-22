package convertor

import (
	"sync"

	"github.com/artemKapitonov/url-shortener/internal/entity"
	url_shortener_v1 "github.com/artemKapitonov/url-shortener/pkg/url-shortener_v1"
)

type EntityConvertor struct {
	r sync.RWMutex
}

func New() *EntityConvertor {
	return &EntityConvertor{}
}

func (c *EntityConvertor) Convert(url *url_shortener_v1.ShortURL) entity.URL {
	return entity.URL{FullURL: "http.localhost:8080"}
}
