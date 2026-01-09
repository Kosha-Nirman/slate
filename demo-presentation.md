---
title: Slate Demo
author: Slate Team
date: 2025-12-25
---

# 🎪 Welcome to Slate

A terminal-based presentation tool

Press **→** to continue

---

## What is Slate?

Slate transforms markdown files into beautiful terminal presentations.

**Key Features:**
- 📝 Markdown rendering
- 🎨 Beautiful themes
- ⌨️  Keyboard navigation
- 🚀 Fast & lightweight

---

## Why Use Slate?

### For Developers
- Stay in your terminal workflow
- Version control your presentations
- Write slides in markdown

### For Everyone
- Simple and distraction-free
- No need for heavy presentation software
- Cross-platform support

---

## Code Examples

Slate supports syntax highlighting:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello from Slate!")

    // Present anywhere, anytime
    presentation := NewPresentation()
    presentation.Show()
}
```

---

## Markdown Features

You can use all standard markdown:

- **Bold text**
- *Italic text*
- `inline code`
- [Links](https://github.com)

### Lists
1. Numbered lists
2. Work great
3. For sequential content

---

## Tables Support

| Feature | Status | Description |
|---------|--------|-------------|
| Markdown | ✅ | Full support |
| Themes | ✅ | Dark & Light |
| Navigation | ✅ | Smooth & Fast |
| Config | ✅ | YAML based |

---

## Quick Start

### 1. Create a presentation

```bash
slate init my-slides.md
```

### 2. Edit your slides

Use your favorite editor to write markdown

### 3. Present!

```bash
slate present my-slides.md
```

---

## Navigation Shortcuts

| Action | Keys |
|--------|------|
| Next slide | →, Space, L |
| Previous slide | ←, H |
| First slide | Home, G |
| Last slide | End, Shift+G |
| Help | ? |
| Quit | Q, Esc |

**Try pressing `?` now for help!**

---

## Configuration

Customize everything via `~/.config/slate/slate.yaml`:

- 🎨 Themes (dark, light, dracula)
- ⌨️  Keybindings
- 📏 Word wrap & margins
- 📊 Progress indicators

```bash
slate config init  # Create config
slate config show  # View config
```

---

## Built With Love

Slate uses amazing open-source libraries:

- **Bubble Tea** - TUI framework
- **Glamour** - Markdown rendering
- **Lipgloss** - Styling
- **Cobra** - CLI framework

All from the incredible [Charm](https://charm.sh) team! 💜

---

## Get Started Today

### Install
```bash
go install github.com/Kosha-Nirman/slate@latest
```

### Create
```bash
slate init
```

### Present
```bash
slate present presentation.md
```

---

# 🎉 Thank You!

## Try it yourself:
- Create your first presentation
- Customize themes
- Share with your team

### Questions?
Press **?** for help
Press **Q** to quit

**Happy Presenting! 🚀**
