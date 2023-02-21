package main

import (
	"fmt"
	"os"

	"github.com/ganeshdipdumbare/gmake/internal/makefile"
	"github.com/ganeshdipdumbare/gmake/internal/tui"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	targets := makefile.GetMakeTargets()
	items := []list.Item{}
	for _, v := range targets {
		items = append(items, tui.Item{
			Name: v.Name,
			Desc: v.Description,
		})
	}

	m := tui.Model{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.List.Title = "Select target"

	p := tea.NewProgram(m, tea.WithInputTTY())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
