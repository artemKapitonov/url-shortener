package dto

import "github.com/artemKapitonov/url-shortener/internal/entity"

type FullURL struct {
	Url string `json:"url,omitempty"`
}

func (u *FullURL) ConvertToEntity() entity.URL {
	return entity.URL{
		FullURL: u.Url,
	}
}

type ShortURL struct {
	Url string `json:"url,omitempty"`
}

func (u *ShortURL) ConvertToEntity() entity.URL {
	return entity.URL{
		ShortURL: u.Url,
	}
}
