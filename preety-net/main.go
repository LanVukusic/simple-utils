package main

import (
	"fmt"
	"net"
	"os"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type netAdapter struct {
	title, desc string
}

func localAddresses() (error, []list.Item) {
	ifaces, err := net.Interfaces()
	items := []list.Item{}

	if err != nil {
		fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
		return err, nil
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			switch v := a.(type) {

			case *net.IPNet:
				if !v.IP.IsLoopback() && v.IP.To4() != nil {
					items = append(items, item{title: i.Name, desc: v.IP.String()})
				}

			}

		}
	}
	return nil, items
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {

			i, ok := m.list.SelectedItem().(item)
			if ok {
				fmt.Printf("desc: %s", string(i.desc))
				clipboard.WriteAll(i.desc)
			}
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	var err, items = localAddresses()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create model
	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}

	// title styling
	m.list.Title = "Wifi interfaces and addresses"
	m.list.Styles.Title.Underline(true)
	m.list.Styles.Title.Background(lipgloss.NoColor{})

	//list settings
	m.list.ShowStatusBar()
	m.list.SetFilteringEnabled(false)
	m.list.SetShowHelp(false)

	// run tea
	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
