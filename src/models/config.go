package models

type ThemeConfig struct {
	Mode         string
	GlamourStyle string
	ShowProgress bool
	ShowSlideNum bool
}

type PresentationConfig struct {
	WordWrap int
	Margin   int
	Padding  int
}

type KeybindingConfig struct {
	Next     []string
	Previous []string
	First    []string
	Last     []string
	Quit     []string
}

type Config struct {
	Theme        ThemeConfig
	Presentation PresentationConfig
	Keybindings  KeybindingConfig
}

func NewDefaultConfig() *Config {
	return &Config{
		Theme: ThemeConfig{
			Mode:         "auto",
			GlamourStyle: "",
			ShowProgress: true,
			ShowSlideNum: true,
		},
		Presentation: PresentationConfig{
			WordWrap: 80,
			Margin:   2,
			Padding:  1,
		},
		Keybindings: KeybindingConfig{
			Next:     []string{"right", "space", "l"},
			Previous: []string{"left", "h"},
			First:    []string{"home", "g"},
			Last:     []string{"end", "G"},
			Quit:     []string{"q", "esc", "ctrl+c"},
		},
	}
}

func (c *Config) Merge(other *Config) {
	if other == nil {
		return
	}

	// * Merge theme config
	if other.Theme.Mode != "" {
		c.Theme.Mode = other.Theme.Mode
	}
	if other.Theme.GlamourStyle != "" {
		c.Theme.GlamourStyle = other.Theme.GlamourStyle
	}
	c.Theme.ShowProgress = other.Theme.ShowProgress
	c.Theme.ShowSlideNum = other.Theme.ShowSlideNum

	// * Merge presentation config
	if other.Presentation.WordWrap > 0 {
		c.Presentation.WordWrap = other.Presentation.WordWrap
	}
	if other.Presentation.Margin >= 0 {
		c.Presentation.Margin = other.Presentation.Margin
	}
	if other.Presentation.Padding >= 0 {
		c.Presentation.Padding = other.Presentation.Padding
	}

	// * Merge keybindings
	if len(other.Keybindings.Next) > 0 {
		c.Keybindings.Next = other.Keybindings.Next
	}
	if len(other.Keybindings.Previous) > 0 {
		c.Keybindings.Previous = other.Keybindings.Previous
	}
	if len(other.Keybindings.First) > 0 {
		c.Keybindings.First = other.Keybindings.First
	}
	if len(other.Keybindings.Last) > 0 {
		c.Keybindings.Last = other.Keybindings.Last
	}
	if len(other.Keybindings.Quit) > 0 {
		c.Keybindings.Quit = other.Keybindings.Quit
	}
}
