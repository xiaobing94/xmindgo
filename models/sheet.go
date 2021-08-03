package models

type TopicPositioningType string

const (
	TopicPositioningFixed TopicPositioningType = "fixed"
	TopicPositioningFree  TopicPositioningType = "free"
)

type TopicOverlappingType string

const (
	TopicOverlappingOverlap TopicOverlappingType = "overlap"
	TopicOverlappingNone    TopicOverlappingType = "none"
)

type Sheet struct {
	IDComponent
	Title            string               `json:"title,omitempty"`
	RootTopic        *Topic               `json:"rootTopic,omitempty"`
	Style            *Style               `json:"style,omitempty"`
	TopicPositioning TopicPositioningType `json:"topicPositioning,omitempty"` // "free" or "fixed"
	TopicOverlapping TopicOverlappingType `json:"topicOverlapping,omitempty"` // "overlap" or "none"
	Theme            *Theme               `json:"theme,omitempty"`
}

func NewSheet() *Sheet {
	sheet := &Sheet{}
	sheet.UseDefaultTheme()
	sheet.GenID()
	topic := &Topic{}
	sheet.AddRootTopic(topic)
	return sheet
}

func (sheet *Sheet) UseDefaultTheme() {
	sheet.Theme = &Theme{}
	sheet.Theme.SetDefault()
}

func (sheet *Sheet) AddRootTopic(topic *Topic) {
	sheet.RootTopic = topic
}
