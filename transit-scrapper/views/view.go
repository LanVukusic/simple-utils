package views

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type Item string
type LoadingItem struct {
	Msg string
}
type ExitAlert struct {
	Msg string
}

func (i Item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
}

type model struct {
	spinner  spinner.Model
	list     list.Model
	items    []Item
	msg      string
	Quitting bool
	main     bool
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case ExitAlert:
		m.Quitting = true
		m.msg = msg.Msg
		return m, tea.Quit

	case LoadingItem:
		m.msg = msg.Msg
		return m, nil

	case []Item:
		m.items = msg

		temp := []list.Item{}
		for _, i := range m.items {
			temp = append(temp, i)
		}
		m.list.SetItems(temp)

	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "ctrl+d", "q", "enter":
			m.Quitting = true
			os.Exit(0)
			return m, tea.Quit
		}

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	var cmd tea.Cmd
	m.list.SetShowTitle(true)
	m.list.Title = m.msg
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.Quitting {
		return quitTextStyle.Render(fmt.Sprintf("%s", m.msg))
	}

	if m.main {
		return quitTextStyle.Render("Srečno pot")
	}

	if len(m.items) == 0 {
		ms := ""
		if m.msg == "" {
			ms = "Loading results ..."
		} else {
			ms = m.msg
		}
		return fmt.Sprintf("\n\n   %s %s\n\n", m.spinner.View(), ms)
	}
	return "\n" + m.list.View()
}

func InitialModel() model {
	// create spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	const defaultWidth = 20

	l := list.New(nil, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Prevozi: "
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle

	return model{spinner: s, list: l}
}
