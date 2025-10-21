package cmd

import (
	"fmt"
	"gitAi/config"
	"gitAi/ui"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func RunConfigurationSetup() config.UserConfig {
	userConfig := config.UserConfig{
		UseEmoji:        false,
		UseConventional: false,
		Tone:            config.Casual,
		Length:          config.Short,
	}

	err := huh.NewConfirm().
		Title("Use emoji prefixes in commit messages?").
		Description("Add emojis like ✨ feat: or 🐛 fix: to commit messages").
		Affirmative("Yes").
		Negative("No").
		Value(&userConfig.UseEmoji).
		Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = huh.NewConfirm().
		Title("Use Conventional Commits format?").
		Description("Format: type(scope): description").
		Affirmative("Yes").
		Negative("No").
		Value(&userConfig.UseConventional).
		Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = huh.NewSelect[config.Tone]().
		Title("Choose your commit message tone").
		Options(
			huh.NewOption("Professional - formal and matter-of-fact", config.Professional),
			huh.NewOption("Casual - light but clear", config.Casual),
			huh.NewOption("Friendly - warm and approachable", config.Friendly),
		).
		Value(&userConfig.Tone).
		Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = huh.NewSelect[config.Length]().
		Title("Choose your commit message length").
		Options(
			huh.NewOption("Short - subject only (no body)", config.Short),
			huh.NewOption("Normal - subject + optional body (1-4 bullets)", config.Normal),
			huh.NewOption("Long - subject + detailed body (2-6 bullets)", config.Long),
		).
		Value(&userConfig.Length).
		Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	return userConfig
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration with an interactive setup",
	Long: `Initialize configuration with an interactive setup.

This command will guide you through configuring your commit message preferences:
- Emoji: Add emoji prefixes to commit messages
- Format: Conventional commits or free-form messages
- Tone: Professional, casual, or friendly
- Length: Short, normal, or detailed messages

The configuration will be saved to .git/config.json in your git repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		userConfig := RunConfigurationSetup()

		err := config.Save(userConfig)
		if err != nil {
			fmt.Printf("Error saving configuration: %v\n", err)
			os.Exit(1)
		}

		config.PrintConfiguration(userConfig)
		ui.RenderTitle("Configuration saved!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}