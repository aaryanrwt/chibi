package tui

import (
	"context"
	"fmt"
	"strings"

	"chibi/internal/ai"
	"chibi/internal/k8s"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D56F4")).MarginBottom(1)
	promptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))
	errorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	msgStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
)

type aiResponseMsg string
type errMsg error

type model struct {
	textInput textinput.Model
	messages  []string
	err       error
	engine    *ai.Engine
}

func initialModel(engine *ai.Engine) model {
	ti := textinput.New()
	ti.Placeholder = "Ask Chibi to 'diagnose pod <name>'..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 60
	ti.PromptStyle = promptStyle
	ti.TextStyle = msgStyle

	return model{
		textInput: ti,
		messages:  []string{},
		err:       nil,
		engine:    engine,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			input := m.textInput.Value()
			if input != "" {
				m.messages = append(m.messages, fmt.Sprintf("You: %s", input))
				m.textInput.SetValue("")

				if strings.HasPrefix(input, "diagnose pod ") {
					parts := strings.Split(input, " ")
					if len(parts) >= 3 {
						podName := parts[2]
						namespace := "default" // hardcoded for MVP
						m.messages = append(m.messages, fmt.Sprintf("Chibi: Fetching pod %s and generating AI diagnosis (this may take a few seconds)...", podName))

						// Return a command to fetch data and query AI asynchronously
						return m, func() tea.Msg {
							ctx := context.Background()
							pod, err := k8s.GetPod(ctx, namespace, podName)
							if err != nil {
								return errMsg(fmt.Errorf("Failed to get pod: %w", err))
							}

							builder := ai.NewContextBuilder()
							jsonStr, _ := builder.BuildResourceContext(pod)
							prompt := builder.BuildDiagnosticPrompt("Pod", podName, jsonStr)

							if m.engine == nil || m.engine.ActiveProvider == nil {
								return errMsg(fmt.Errorf("No AI provider configured. Please set OPENAI_API_KEY environment variable"))
							}

							resp, err := m.engine.ActiveProvider.Generate(ctx, "You are an expert Kubernetes SRE.", prompt)
							if err != nil {
								return errMsg(fmt.Errorf("AI Error: %w", err))
							}
							return aiResponseMsg(resp)
						}
					}
				} else {
					m.messages = append(m.messages, "Chibi: I only know how to 'diagnose pod <name>' for now.")
				}
			}
		}

	case aiResponseMsg:
		m.messages = append(m.messages, fmt.Sprintf("Chibi: %s", string(msg)))
		return m, nil

	case errMsg:
		m.err = msg
		return m, nil

	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("🤖 Chibi AI SRE Assistant"))
	b.WriteString("\n\n")

	// V2: Render predictive telemetry dashboard
	b.WriteString(RenderDashboard())
	b.WriteString("\n")

	for _, msg := range m.messages {
		b.WriteString(msgStyle.Render(msg) + "\n")
	}
	b.WriteString("\n")

	b.WriteString(m.textInput.View())
	b.WriteString("\n\n(esc to quit)\n")

	if m.err != nil {
		b.WriteString(errorStyle.Render(fmt.Sprintf("\nError: %v", m.err)))
	}

	return b.String()
}

// StartTUI begins the Bubble Tea application loop.
func StartTUI(engine *ai.Engine) error {
	p := tea.NewProgram(initialModel(engine))
	_, err := p.Run()
	return err
}
