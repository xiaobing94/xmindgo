package models

type Image struct {
	Src   string `json:"src,omitempty"` //"xap:resources/<HASH>.png",
	Align string `json:"align,omitempty"`
}

type TopicImage struct {
	Image
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}
