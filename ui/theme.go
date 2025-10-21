package ui

import (
	"gitAi/ui/themes"
	"os"
	"path/filepath"
	"strings"
)

var currentTheme *themes.Theme

func init() {
	loadedTheme := LoadTheme()
	if loadedTheme != "" {
		SetTheme(loadedTheme)
	} else {
		currentTheme = themes.Catppuccin()
	}
}

func SetTheme(name string) bool {
	var theme *themes.Theme
	switch name {
	case "catppuccin":
		theme = themes.Catppuccin()
	default:
		return false
	}
	currentTheme = theme
	return true
}

func GetCurrentTheme() *themes.Theme {
	return currentTheme
}

func GetAvailableThemes() []string {
	return []string{
		"catppuccin",
	}
}

func GetDarkThemes() []string {
	return []string{
		"catppuccin",
	}
}

func SaveTheme(themeName string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dinyDir := filepath.Join(homeDir, ".config", "diny")
	if err := os.MkdirAll(dinyDir, 0755); err != nil {
		return err
	}

	themePath := filepath.Join(dinyDir, "theme")
	return os.WriteFile(themePath, []byte(themeName), 0644)
}

func LoadTheme() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	themePath := filepath.Join(homeDir, ".config", "diny", "theme")
	data, err := os.ReadFile(themePath)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}