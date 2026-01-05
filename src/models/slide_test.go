package models

import (
	"testing"
)

func TestNewSlide(t *testing.T) {
	index := 5
	content := "# Test Slide\n\nSome content"

	slide := NewSlide(index, content)

	if slide.Index != index {
		t.Errorf("Expected Index %d, got %d", index, slide.Index)
	}

	expectedContent := "# Test Slide\n\nSome content"
	if slide.RawContent != expectedContent {
		t.Errorf("Expected content '%s', got '%s'", expectedContent, slide.RawContent)
	}
}

func TestNewSlideTrimsWhitespace(t *testing.T) {
	content := "  \n  # Test Slide  \n  "
	slide := NewSlide(0, content)

	expectedContent := "# Test Slide"
	if slide.RawContent != expectedContent {
		t.Errorf("Expected trimmed content '%s', got '%s'", expectedContent, slide.RawContent)
	}
}

func TestSlideIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected bool
	}{
		{"Empty string", "", true},
		{"Whitespace only", "   \n\t  ", true},
		{"Has content", "# Test", false},
		{"Single character", "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slide := NewSlide(0, tt.content)
			result := slide.IsEmpty()
			if result != tt.expected {
				t.Errorf("Expected IsEmpty() to be %v for content '%s', got %v", tt.expected, tt.content, result)
			}
		})
	}
}

func TestContent(t *testing.T) {
	content := "# Test Slide"
	slide := NewSlide(0, content)

	if slide.Content() != content {
		t.Errorf("Expected Content() to return '%s', got '%s'", content, slide.Content())
	}
}

func TestRenderedCache(t *testing.T) {
	slide := NewSlide(0, "# Test")

	// Initially should not have cache
	if slide.HasCache() {
		t.Error("Expected slide to not have cache initially")
	}

	if slide.GetRenderedCache() != "" {
		t.Error("Expected empty cache initially")
	}

	// Set cache
	rendered := "Rendered content"
	slide.SetRenderedCache(rendered)

	if !slide.HasCache() {
		t.Error("Expected slide to have cache after setting")
	}

	if slide.GetRenderedCache() != rendered {
		t.Errorf("Expected cache '%s', got '%s'", rendered, slide.GetRenderedCache())
	}

	// Clear cache
	slide.ClearCache()

	if slide.HasCache() {
		t.Error("Expected slide to not have cache after clearing")
	}

	if slide.GetRenderedCache() != "" {
		t.Error("Expected empty cache after clearing")
	}
}

func TestSlideMetadata(t *testing.T) {
	slide := NewSlide(0, "# Test")

	// Default metadata should be empty
	if slide.Metadata.Notes != "" {
		t.Error("Expected empty notes by default")
	}

	if slide.Metadata.Transition != "" {
		t.Error("Expected empty transition by default")
	}

	if slide.Metadata.Background != "" {
		t.Error("Expected empty background by default")
	}

	// Set metadata
	slide.Metadata.Notes = "Test notes"
	slide.Metadata.Transition = "fade"
	slide.Metadata.Background = "#000000"

	if slide.Metadata.Notes != "Test notes" {
		t.Errorf("Expected notes 'Test notes', got '%s'", slide.Metadata.Notes)
	}

	if slide.Metadata.Transition != "fade" {
		t.Errorf("Expected transition 'fade', got '%s'", slide.Metadata.Transition)
	}

	if slide.Metadata.Background != "#000000" {
		t.Errorf("Expected background '#000000', got '%s'", slide.Metadata.Background)
	}
}
