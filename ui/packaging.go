package ui

import (
	"github.com/amihere/spring-now-now/springboot"
	springlist "github.com/amihere/spring-now-now/springlist"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func getPackagingOptions(sp springboot.SpringBoot) []list.Item {
	var items []list.Item
	for _, v := range sp.Packaging.Values {
		items = append(items, springlist.NormalListItem{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return items
}

func (m Model) updatePackaging(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			i, ok := m.list.SelectedItem().(springlist.NormalListItem)
			if ok {
				m.Packaging = i.ID
				m.state = chooseJavaVersion
				m.list = springlist.NewNormalListModel(
					"Java Version",
					getJavaVersion(m.springBoot),
					m.JavaVersion,
					m.list.Width(), m.list.Height(),
				)
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 1)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) viewPackaging() string {
	return "\n  " + m.list.View()
}
