package models

type TopicType string

const (
	TopicTypeAttached TopicType = "attached"
	TopicTypeDetached TopicType = "detached"
	TopicTypeSummary  TopicType = "summary"
	TopicTypeCallout  TopicType = "callout"
	TopicTypeRoot     TopicType = "root"
)

type Topic struct {
	IDComponent
	Title          string `json:"title,omitempty"`
	Style          *Style `json:"style,omitempty"`
	Class          string `json:"class,omitempty"`
	Position       string `json:"position,omitempty"`
	StructureClass string `json:"structureClass,omitempty"`
	Branch         string `json:"branch,omitempty"`
	Width          int64  `json:"width,omitempty"`
	Labels         string `json:"labels,omitempty"`
	// "href": "xap:resources/<HASH>.doc" (attachment reference)
	// "href": "xmind:#<ID>"  (object hyperlink)
	// "href": "http://www.google.com"  (web hyperlink)
	// "href": "file:///Users/user/Documents/test.doc"  (file hyperlink)
	Href      string              `json:"href,omitempty"`
	Notes     *Notes              `json:"notes,omitempty"`
	Image     *TopicImage         `json:"image,omitempty"`
	Children  map[string][]*Topic `json:"children,omitempty"`
	Summaries []*Summary          `json:"summaries,omitempty"`
}

func NewTopic(title string) *Topic {
	topic := &Topic{
		Title: title,
	}
	topic.GenID()
	return topic
}

func (t *Topic) GetChildren(topicType TopicType) []*Topic {
	if t.Children == nil {
		return nil
	}
	return t.Children[string(topicType)]
}

func (t *Topic) GetAttachedChildren() []*Topic {
	return t.GetChildren(TopicTypeAttached)
}

func (t *Topic) AddChildTopic(topic *Topic, topicType TopicType) {
	if t.Children == nil {
		t.Children = make(map[string][]*Topic)
	}
	topics := t.Children[string(topicType)]
	topics = append(topics, topic)
	t.Children[string(topicType)] = topics
}

func (t *Topic) AddAttachedChildTopic(topic *Topic) {
	t.AddChildTopic(topic, TopicTypeAttached)
}
