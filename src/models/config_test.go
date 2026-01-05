package models

import (
	"testing"
)

func TestNewDefaultConfig(t *testing.T) {
	config := NewDefaultConfig()

	// Test theme defaults
	if config.Theme.Mode != "auto" {
		t.Errorf("Expected theme mode 'auto', got '%s'", config.Theme.Mode)
	}

	if !config.Theme.ShowProgress {
		t.Error("Expected ShowProgress to be true by default")
	}

	if !config.Theme.ShowSlideNum {
		t.Error("Expected ShowSlideNum to be true by default")
	}

	// Test presentation defaults
	if config.Presentation.WordWrap != 80 {
		t.Errorf("Expected WordWrap 80, got %d", config.Presentation.WordWrap)
	}

	if config.Presentation.Margin != 2 {
		t.Errorf("Expected Margin 2, got %d", config.Presentation.Margin)
	}

	if config.Presentation.Padding != 1 {
		t.Errorf("Expected Padding 1, got %d", config.Presentation.Padding)
	}

	// Test keybinding defaults
	if len(config.Keybindings.Next) == 0 {
		t.Error("Expected Next keybindings to be set")
	}

	if len(config.Keybindings.Previous) == 0 {
		t.Error("Expected Previous keybindings to be set")
	}

	if len(config.Keybindings.Quit) == 0 {
		t.Error("Expected Quit keybindings to be set")
	}
}

func TestConfigMerge(t *testing.T) {
	config := NewDefaultConfig()

	// Create another config with different values
	other := &Config{
		Theme: ThemeConfig{
			Mode:         "dark",
			GlamourStyle: "dracula",
			ShowProgress: false,
			ShowSlideNum: false,
		},
		Presentation: PresentationConfig{
			WordWrap: 100,
			Margin:   4,
			Padding:  2,
		},
		Keybindings: KeybindingConfig{
			Next:     []string{"n"},
			Previous: []string{"p"},
			First:    []string{"f"},
			Last:     []string{"l"},
			Quit:     []string{"x"},
		},
	}

	// Merge
	config.Merge(other)

	// Check that values were merged
	if config.Theme.Mode != "dark" {
		t.Errorf("Expected theme mode 'dark' after merge, got '%s'", config.Theme.Mode)
	}

	if config.Theme.GlamourStyle != "dracula" {
		t.Errorf("Expected glamour style 'dracula' after merge, got '%s'", config.Theme.GlamourStyle)
	}

	if config.Theme.ShowProgress {
		t.Error("Expected ShowProgress to be false after merge")
	}

	if config.Presentation.WordWrap != 100 {
		t.Errorf("Expected WordWrap 100 after merge, got %d", config.Presentation.WordWrap)
	}

	if len(config.Keybindings.Next) != 1 || config.Keybindings.Next[0] != "n" {
		t.Error("Expected Next keybindings to be merged")
	}
}

func TestConfigMergeWithNil(t *testing.T) {
	config := NewDefaultConfig()
	originalMode := config.Theme.Mode

	// Merge with nil should not change anything
	config.Merge(nil)

	if config.Theme.Mode != originalMode {
		t.Error("Expected config to remain unchanged when merging with nil")
	}
}

func TestConfigMergeEmptyValues(t *testing.T) {
	config := NewDefaultConfig()
	originalMode := config.Theme.Mode

	// Merge with empty values should not override
	other := &Config{
		Theme: ThemeConfig{
			Mode: "", // Empty mode should not override
		},
	}

	config.Merge(other)

	if config.Theme.Mode != originalMode {
		t.Error("Expected empty theme mode to not override existing value")
	}
}

func TestThemeConfig(t *testing.T) {
	theme := ThemeConfig{
		Mode:         "dark",
		GlamourStyle: "dracula",
		ShowProgress: true,
		ShowSlideNum: true,
	}

	if theme.Mode != "dark" {
		t.Errorf("Expected mode 'dark', got '%s'", theme.Mode)
	}

	if theme.GlamourStyle != "dracula" {
		t.Errorf("Expected glamour style 'dracula', got '%s'", theme.GlamourStyle)
	}

	if !theme.ShowProgress {
		t.Error("Expected ShowProgress to be true")
	}

	if !theme.ShowSlideNum {
		t.Error("Expected ShowSlideNum to be true")
	}
}

func TestPresentationConfig(t *testing.T) {
	presentation := PresentationConfig{
		WordWrap: 120,
		Margin:   3,
		Padding:  2,
	}

	if presentation.WordWrap != 120 {
		t.Errorf("Expected WordWrap 120, got %d", presentation.WordWrap)
	}

	if presentation.Margin != 3 {
		t.Errorf("Expected Margin 3, got %d", presentation.Margin)
	}

	if presentation.Padding != 2 {
		t.Errorf("Expected Padding 2, got %d", presentation.Padding)
	}
}

func TestKeybindingConfig(t *testing.T) {
	keybindings := KeybindingConfig{
		Next:     []string{"right", "space"},
		Previous: []string{"left"},
		First:    []string{"home"},
		Last:     []string{"end"},
		Quit:     []string{"q", "esc"},
	}

	if len(keybindings.Next) != 2 {
		t.Errorf("Expected 2 Next keybindings, got %d", len(keybindings.Next))
	}

	if keybindings.Next[0] != "right" {
		t.Errorf("Expected first Next keybinding to be 'right', got '%s'", keybindings.Next[0])
	}

	if len(keybindings.Quit) != 2 {
		t.Errorf("Expected 2 Quit keybindings, got %d", len(keybindings.Quit))
	}
}
