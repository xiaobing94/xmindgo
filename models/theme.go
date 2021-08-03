package models

type Theme struct {
	IDComponent
	Title                string `json:"title,omitempty"`
	Map                  *Style `json:"map,omitempty"`
	CentralTopic         *Style `json:"centralTopic,omitempty"`
	MainTopic            *Style `json:"mainTopic,omitempty"`
	SubTopic             *Style `json:"subTopic,omitempty"`
	FloatingTopic        *Style `json:"floatingTopic,omitempty"`
	CentralFloatingTopic *Style `json:"centralFloatingTopic,omitempty"`
	Boundary             *Style `json:"boundary,omitempty"`
	Relationship         *Style `json:"relationship,omitempty"`
	SummaryTopic         *Style `json:"summaryTopic,omitempty"`
	Summary              *Style `json:"summary,omitempty"`
}

func (theme *Theme) SetDefault() {
	theme.CentralTopic = GetDefaultCentralTopicStyle()
	theme.MainTopic = GetDefaultMainTopicStyle()
	theme.SubTopic = GetDefaultSubTopicStyle()
}
