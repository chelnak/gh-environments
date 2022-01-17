package list

// A simple program demonstrating the paginator component from the Bubbles
// component library.

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"
	"github.com/chelnak/gh-environments/internal/client"
	"github.com/google/go-github/v42/github"
	"github.com/olekukonko/tablewriter"

	tea "github.com/charmbracelet/bubbletea"
)

var githubClient client.Client

type model struct {
	items     *github.EnvResponse
	paginator paginator.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	// if !m.paginator.OnLastPage() {
	// 	log.Fatal("not on last page")
	// 	m.items = get()
	// }

	m.paginator, cmd = m.paginator.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s := strings.Builder{}

	renderTable(m.items.Environments, &s)

	s.WriteString("\nh/l ←/→ page • q: quit\n")
	s.WriteString(m.paginator.View())

	return s.String()
}

func renderTable(environments []*github.Environment, writer io.Writer) {
	table := tablewriter.NewWriter(writer)
	timeFormat := "2006-01-02 15:04:05"
	for _, environment := range environments {
		table.Append([]string{
			fmt.Sprintf("%d", *environment.ID),
			*environment.Name,
			environment.CreatedAt.Local().Format(timeFormat),
			environment.UpdatedAt.Local().Format(timeFormat),
		})
	}

	table.SetHeader([]string{"id", "name", "created", "updated"})

	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetHeaderLine(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoFormatHeaders(false)
	table.SetColumnSeparator("")
	table.SetTablePadding("  ") // two spaces
	table.SetNoWhiteSpace(true)

	table.SetColumnColor(nil,
		nil,
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor})

	table.Render()
}

func newModel(items *github.EnvResponse, perPage int) model {
	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = perPage
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	p.SetTotalPages(items.GetTotalCount())

	return model{
		paginator: p,
		items:     items,
	}
}

func newPaginatedTable(c client.Client, perPage int) {
	githubClient = c
	envResponse, err := githubClient.GetEnvironments()
	if err != nil {
		log.Fatal(err)
	}

	if *envResponse.TotalCount == 0 {
		fmt.Printf("There are no environments in %s/%s\n", githubClient.Owner(), githubClient.Repo())
		return
	}

	fmt.Printf(
		"Showing %d of %d environments in %s/%s\n\n",
		len(envResponse.Environments),
		*envResponse.TotalCount,
		githubClient.Owner(),
		githubClient.Repo(),
	)

	m := newModel(envResponse, perPage)

	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
