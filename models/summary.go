package models

type Summary struct {
	IDComponent
	Style   *Style  `json:"style,omitempty"`
	Class   string `json:"class,omitempty"`
	Range   string `json:"range,omitempty"`
	TopicID string `json:"topicId,omitempty"`
}
