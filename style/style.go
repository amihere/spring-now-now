package style

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	// ErrorStyle globally styles errors
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff0033")).
			Bold(true)
		// SuccessStyle
	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5f9f2b")).
			Bold(true)
		// TitleStyle
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#fff")).
			Background(lipgloss.Color("#56972b")).
			PaddingLeft(1).
			PaddingRight(1)
		// ItemStyle
	ItemStyle = lipgloss.NewStyle().PaddingLeft(4)
	// SelectedNormalListItemStyle
	SelectedNormalListItemStyle = lipgloss.NewStyle().
					PaddingLeft(2).
					Bold(true).
					Foreground(greenColor)
		// PaginationStyle
	PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	// HelpStyle
	HelpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	// QuitTextStyle
	QuitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	// ChoosenTitleStyle
	ChoosenTitleStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				BorderForeground(lipgloss.AdaptiveColor{Light: "#84f5a2", Dark: "#3cc962"}).
				Foreground(lipgloss.AdaptiveColor{Light: "#61c77d", Dark: "#61c77d"}).
				Padding(0, 0, 0, 1)
		// ChoosenDescStyle
	ChoosenDescStyle = ChoosenTitleStyle.Copy().
				Foreground(lipgloss.AdaptiveColor{Light: "#84f5a2", Dark: "#84f5a2"})
		// DocStyle
	DocStyle = lipgloss.NewStyle().Margin(1, 2)
	// InputTitleStyle
	InputTitleStyle = lipgloss.NewStyle().Bold(true).Foreground(greenColor)
)
