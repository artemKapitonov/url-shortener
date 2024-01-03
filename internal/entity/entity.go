package entity

type FullURL struct {
	ID  uint64 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
}

type ShortURL struct {
	ID  uint64 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
}
