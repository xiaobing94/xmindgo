package models

type NoteSpan interface {
	IsNoteSpan() bool
}

type TextSpan struct {
	Style *Style `json:"style"`
	Text  string `json:"text"`
	Class string `json:"class"`
}

func (s *TextSpan) IsNoteSpan() bool {
	return true
}

type ImageSpan struct {
	Style *Style `json:"style,omitempty"`
	Class string `json:"class,omitempty"`
	Image string `json:"image,omitempty"` // "xap:resources/<HASH>.jpg"
}

func (s *ImageSpan) IsNoteSpan() bool {
	return true
}

type HyperlinkSpan struct {
	Style *Style     `json:"style,omitempty"`
	Class string     `json:"class,omitempty"`
	Href  string     `json:"href,omitempty"`
	Spans []NoteSpan `json:"spans,omitempty"`
}

func (s *HyperlinkSpan) IsNoteSpan() bool {
	return true
}

type NotePlain struct {
	Content string `json:"content,omitempty"`
}

type NoteHtmlParagraph struct {
	Style *Style     `json:"style,omitempty"`
	Spans []NoteSpan `json:"spans,omitempty"`
}

type NoteHtmlContent struct {
	Paragraphs []NoteHtmlParagraph `json:"paragraphs,omitempty"`
}

type NoteHtml struct {
	Content NoteHtmlParagraph `json:"content,omitempty"`
}

type Notes struct {
	Plain NotePlain `json:"plain,omitempty"`
	Html  NoteHtml  `json:"html,omitempty"`
}
