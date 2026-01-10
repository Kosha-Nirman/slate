package models

import (
	"errors"
	"time"
)

type Presentation struct {
	FilePath string
	Title    string
	Author   string
	Date     time.Time
	Slides   []*Slide
	Config   *Config
}

func NewPresentation(filePath string) *Presentation {
	return &Presentation{
		FilePath: filePath,
		Date:     time.Now(),
		Slides:   make([]*Slide, 0),
	}
}

func (p *Presentation) AddSlide(slide *Slide) {
	p.Slides = append(p.Slides, slide)
}

func (p *Presentation) GetSlide(index int) (*Slide, error) {
	if index < 0 || index > len(p.Slides) {
		return nil, errors.New("slide index out of bounds")
	}

	return p.Slides[index], nil
}

func (p *Presentation) SlideCount() int {
	return len(p.Slides)
}

func (p *Presentation) IsEmpty() bool {
	return len(p.Slides) == 0
}

func (p *Presentation) SetMetadata(metadata map[string]string) {
	if title, ok := metadata["title"]; ok {
		p.Title = title
	}

	if author, ok := metadata["author"]; ok {
		p.Author = author
	}

	if dateStr, ok := metadata["date"]; ok {
		if date, err := time.Parse("2006-01-02", dateStr); err == nil {
			p.Date = date
		}
	}
}

func (p *Presentation) GetMetadata() map[string]string {
	metadata := make(map[string]string)

	if p.Title != "" {
		metadata["title"] = p.Title
	}
	if p.Author != "" {
		metadata["author"] = p.Author
	}
	if !p.Date.IsZero() {
		metadata["date"] = p.Date.Format("2006-01-02")
	}

	return metadata
}

func (p *Presentation) Validate() error {
	if p.FilePath == "" {
		return errors.New("presentation file path is required")
	}

	if len(p.Slides) == 0 {
		return errors.New("presentation must have at least one slide")
	}

	hasContent := false
	for _, slide := range p.Slides {
		if !slide.IsEmpty() {
			hasContent = true
			break
		}
	}

	if !hasContent {
		return errors.New("presentation must have at least one non-empty slide")
	}

	return nil
}
