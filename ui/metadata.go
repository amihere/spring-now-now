package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	springlist "github.com/amihere/spring-now-now/springlist"
	"github.com/amihere/spring-now-now/style"
)

func (m Model) updateMetaData(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter, tea.KeyCtrlS:
			if m.focused == len(m.inputs)-1 || msg.Type == tea.KeyCtrlS {
				input := m.inputs[groupID].Value()
				if input != "" {
					m.GroupID = input
				}

				input = m.inputs[artifactID].Value()
				if input != "" {
					m.ArtifactID = input
				}

				input = m.inputs[applicationName].Value()
				if input != "" {
					m.Name = input
				}

				input = m.inputs[description].Value()
				if input != "" {
					m.Description = input
				}

				input = m.inputs[packageName].Value()
				if input != "" {
					m.PackageName = input
				}

				input = m.inputs[version].Value()
				if input != "" {
					m.Version = input
				}

				m.state = choosePackaging
				m.list = springlist.NewNormalListModel(
					"Packaging",
					getPackagingOptions(m.springBoot),
					m.Packaging,
					m.list.Width(), m.list.Height(),
				)
				return m, nil
			}
			m.nextInput()
		case tea.KeyShiftTab, tea.KeyCtrlP, tea.KeyUp:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN, tea.KeyDown:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) viewMetaData() string {
	return fmt.Sprintf(
		`
    %s

    %s
    %s

    %s
    %s

    %s
    %s

    %s
    %s

    %s
    %s

    %s
    %s
`,
		style.TitleStyle.Render("Project Metadata"),
		style.InputTitleStyle.Width(60).Render("Group ID"),
		m.inputs[groupID].View(),
		style.InputTitleStyle.Width(60).Render("Artifact ID"),
		m.inputs[artifactID].View(),
		style.InputTitleStyle.Width(60).Render("Application Name"),
		m.inputs[applicationName].View(),
		style.InputTitleStyle.Width(60).Render("Description"),
		m.inputs[description].View(),
		style.InputTitleStyle.Width(60).Render("Package Name"),
		m.inputs[packageName].View(),
		style.InputTitleStyle.Width(60).Render("Version"),
		m.inputs[version].View(),
	) + "\n"
}

// nextInput focuses the next input field
func (m *Model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *Model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
