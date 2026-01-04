package models

import "strings"

type SlideMetadata struct {
	Notes      string
	Transition string
	Background string
}

type Slide struct {
	Index         int
	RawContent    string
	RenderedCache string
	Metadata      SlideMetadata
}

func NewSlide(index int, content string) *Slide {
	return &Slide{
		Index:      index,
		RawContent: strings.TrimSpace(content),
		Metadata:   SlideMetadata{},
	}
}

func (s *Slide) IsEmpty() bool {
	return strings.TrimSpace(s.RawContent) == ""
}

func (s *Slide) Content() string {
	return s.RawContent
}

func (s *Slide) SetRenderedCache(rendered string) {
	s.RenderedCache = rendered
}

func (s *Slide) GetRenderedCache() string {
	return s.RenderedCache
}

func (s *Slide) HasCache() bool {
	return s.RenderedCache != ""
}

func (s *Slide) ClearCache() {
	s.RenderedCache = ""
}
