package models

// Style 样式
type Style struct {
	IDComponent
	Type       string            `json:"type,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

// NewStyle 实例化样式
func NewStyle(styleType string) *Style {
	return &Style{
		Type:       styleType,
		Properties: map[string]string{},
	}
}

// AddProperty 添加属性
func (s *Style) AddProperty(key string, value string) {
	if s.Properties == nil {
		s.Properties = make(map[string]string)
	}
	s.Properties[key] = value
}

// RemoveProperty 移除某个属性
func (s *Style) RemoveProperty(key string) {
	if s.Properties == nil {
		return
	}
	delete(s.Properties, key)
}

// GetDefaultSubTopicStyle 获取默认subTopic主题
func GetDefaultSubTopicStyle() *Style {
	style := &Style{
		Type: "topic",
		Properties: map[string]string{
			"border-line-color": "#558ED5",
			"border-line-width": "3pt",
			"fo:font-family":    "Microsoft YaHei",
			"line-class":        "org.xmind.branchConnection.curve",
			"line-color":        "#558ED5",
			"line-width":        "1pt",
		},
	}
	style.GenID()
	return style
}

// GetDefaultCentralTopicStyle 获取默认centralTopic主题
func GetDefaultCentralTopicStyle() *Style {
	style := &Style{
		Type: "topic",
		Properties: map[string]string{
			"border-line-color": "#558ED5",
			"border-line-width": "5pt",
			"fo:color":          "#376092",
			"fo:font-family":    "Microsoft YaHei",
			"line-class":        "org.xmind.branchConnection.curve",
			"line-color":        "#558ED5",
			"line-width":        "1pt",
			"shape-class":       "org.xmind.topicShape.roundedRect",
			"svg:fill":          "#DCE6F2",
		},
	}
	style.GenID()
	return style
}

// GetDefaultMainTopicStyle 获取默认mainTopic主题
func GetDefaultMainTopicStyle() *Style {
	style := &Style{
		Type: "topic",
		Properties: map[string]string{
			"border-line-color": "#558ED5",
			"border-line-width": "2pt",
			"fo:color":          "#17375E",
			"fo:font-family":    "Microsoft YaHei",
			"line-class":        "org.xmind.branchConnection.curve",
			"line-color":        "#558ED5",
			"line-width":        "1pt",
			"shape-class":       "org.xmind.topicShape.roundedRect",
			"svg:fill":          "#DCE6F2",
		},
	}
	style.GenID()
	return style
}

