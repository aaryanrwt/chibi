package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	sparklineStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))
)

// RenderDashboard returns a V2 live telemetry dashboard string.
func RenderDashboard() string {
	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().Bold(true).Render("Predictive Telemetry (V2)"))
	b.WriteString("\nCPU Usage Trend:\n")
	b.WriteString(sparklineStyle.Render(" ▂▃▄▅▆▇█▇▆▅▄▃▂ "))
	b.WriteString("\n")

	return b.String()
}
