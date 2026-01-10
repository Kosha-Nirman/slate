package models

import (
	"testing"
	"time"
)

func TestNewPresentation(t *testing.T) {
	filepath := "/path/to/test.md"
	p := NewPresentation(filepath)

	if p.FilePath != filepath {
		t.Errorf("Expected FilePath %s, got %s", filepath, p.FilePath)
	}

	if p.Slides == nil {
		t.Error("Expected Slides to be initialized")
	}

	if len(p.Slides) != 0 {
		t.Errorf("Expected empty Slides, got %d slides", len(p.Slides))
	}

	if p.Date.IsZero() {
		t.Error("Expected Date to be set")
	}
}

func TestAddSlide(t *testing.T) {
	p := NewPresentation("test.md")
	slide := NewSlide(0, "# Test Slide")

	p.AddSlide(slide)

	if len(p.Slides) != 1 {
		t.Errorf("Expected 1 slide, got %d", len(p.Slides))
	}

	if p.Slides[0] != slide {
		t.Error("Expected slide to be added")
	}
}

func TestGetSlide(t *testing.T) {
	p := NewPresentation("test.md")
	slide1 := NewSlide(0, "# Slide 1")
	slide2 := NewSlide(1, "# Slide 2")

	p.AddSlide(slide1)
	p.AddSlide(slide2)

	// Test valid index
	s, err := p.GetSlide(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if s != slide1 {
		t.Error("Expected to get slide1")
	}

	// Test invalid index (negative)
	_, err = p.GetSlide(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}

	// Test invalid index (too large)
	_, err = p.GetSlide(10)
	if err == nil {
		t.Error("Expected error for out of bounds index")
	}
}

func TestSlideCount(t *testing.T) {
	p := NewPresentation("test.md")

	if p.SlideCount() != 0 {
		t.Errorf("Expected 0 slides, got %d", p.SlideCount())
	}

	p.AddSlide(NewSlide(0, "# Slide 1"))
	if p.SlideCount() != 1 {
		t.Errorf("Expected 1 slide, got %d", p.SlideCount())
	}

	p.AddSlide(NewSlide(1, "# Slide 2"))
	if p.SlideCount() != 2 {
		t.Errorf("Expected 2 slides, got %d", p.SlideCount())
	}
}

func TestMetadata(t *testing.T) {
	p := NewPresentation("test.md")
	p.Title = "Test Presentation"
	p.Author = "Test Author"
	p.Date = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	metadata := p.GetMetadata()

	if metadata["title"] != "Test Presentation" {
		t.Errorf("Expected title 'Test Presentation', got '%s'", metadata["title"])
	}

	if metadata["author"] != "Test Author" {
		t.Errorf("Expected author 'Test Author', got '%s'", metadata["author"])
	}

	if metadata["date"] != "2025-01-01" {
		t.Errorf("Expected date '2025-01-01', got '%s'", metadata["date"])
	}
}

func TestSetMetadata(t *testing.T) {
	p := NewPresentation("test.md")

	metadata := map[string]string{
		"title":  "New Title",
		"author": "New Author",
		"date":   "2025-12-25",
	}

	p.SetMetadata(metadata)

	if p.Title != "New Title" {
		t.Errorf("Expected title 'New Title', got '%s'", p.Title)
	}

	if p.Author != "New Author" {
		t.Errorf("Expected author 'New Author', got '%s'", p.Author)
	}

	expectedDate := time.Date(2025, 12, 25, 0, 0, 0, 0, time.UTC)
	if !p.Date.Equal(expectedDate) {
		t.Errorf("Expected date %v, got %v", expectedDate, p.Date)
	}
}

func TestValidate(t *testing.T) {
	// Test empty file path
	p := NewPresentation("")
	p.AddSlide(NewSlide(0, "# Test"))
	err := p.Validate()
	if err == nil {
		t.Error("Expected error for empty file path")
	}

	// Test no slides
	p = NewPresentation("test.md")
	err = p.Validate()
	if err == nil {
		t.Error("Expected error for no slides")
	}

	// Test all empty slides
	p = NewPresentation("test.md")
	p.AddSlide(NewSlide(0, "   "))
	p.AddSlide(NewSlide(1, ""))
	err = p.Validate()
	if err == nil {
		t.Error("Expected error for all empty slides")
	}

	// Test valid presentation
	p = NewPresentation("test.md")
	p.AddSlide(NewSlide(0, "# Test Slide"))
	err = p.Validate()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestPresentationIsEmpty(t *testing.T) {
	p := NewPresentation("test.md")

	if !p.IsEmpty() {
		t.Error("Expected presentation to be empty")
	}

	p.AddSlide(NewSlide(0, "# Test"))

	if p.IsEmpty() {
		t.Error("Expected presentation to not be empty")
	}
}
