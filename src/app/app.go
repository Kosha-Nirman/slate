package app

import (
	"fmt"

	"github.com/Kosha-Nirman/slate/src/config"
	"github.com/Kosha-Nirman/slate/src/display"
	"github.com/Kosha-Nirman/slate/src/models"
	"github.com/Kosha-Nirman/slate/src/theme"
	tea "github.com/charmbracelet/bubbletea"
)

// * ViewMode represents different view modes
type ViewMode int

const (
	ViewPresentation ViewMode = iota
	ViewHelp
)

// * BubbleTea model for App
type App struct {
	config *models.Config

	presentation *models.Presentation
	viewMode     ViewMode

	theme    *theme.Manager
	renderer *display.Renderer

	width  int
	height int

	err   error
	ready bool
}

func New(filePath string) (*App, error) {
	configLoader := config.New()
	cfg, err := configLoader.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	themeManager := theme.NewManager(&cfg.Theme)

	return &App{
		config:   cfg,
		viewMode: ViewPresentation,
		theme:    themeManager,
	}, nil
}

func (a *App) Init() tea.Cmd {
	return nil
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return a, nil
}

func (a *App) View() string {
	return ""
}

func Run(filepath string) error {
	app, err := New(filepath)
	if err != nil {
		return err
	}

	p := tea.NewProgram(
		app,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running program: %w", err)
	}

	return nil
}
