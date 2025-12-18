package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg struct{}
type finishedMsg struct{}

type poller struct {
	count int
}

func (p *poller) Init() tea.Cmd {
	p.count = 1
	return nil
}

func (p *poller) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return p, tea.Quit
		}
	case tickMsg:
		p.count++
	case finishedMsg:
		return p, tea.Quit
	}

	return p, poll(p)
}

func (p *poller) View() string {
	return fmt.Sprintf("Count: %d", p.count)
}

func poll(p *poller) tea.Cmd {
	if p.count < 10 {
		return tea.Tick(1*time.Second, func(time.Time) tea.Msg {
			return tickMsg{}
		})
	}
	return tea.Tick(1*time.Second, func(time.Time) tea.Msg {
		return finishedMsg{}
	})
}

func main() {
	tea.NewProgram(&poller{}).Run()
}
