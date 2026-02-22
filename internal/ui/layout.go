package ui

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/fayssalelmofatiche/gtop/internal/git"
)

var (
	labelStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6CB6FF")).Width(14)
	valueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#E6EDF3"))
	dimStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#8B949E"))
	logoStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#F0883E"))
	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#F0883E")).
			Padding(0, 2)
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#F0883E"))
)

const logo = `
   ____ _____ ___  ____
  / ___|_   _/ _ \|  _ \
 | |  _  | || | | | |_) |
 | |_| | | || |_| |  __/
  \____| |_| \___/|_|`

func RenderLogo() string {
	return borderStyle.Render(logoStyle.Render(logo))
}

func row(label, value string) string {
	if value == "" {
		value = "-"
	}
	return labelStyle.Render(label) + valueStyle.Render(value)
}

func formatLOC(loc int) string {
	switch {
	case loc < 1000:
		return fmt.Sprintf("%d", loc)
	case loc < 1_000_000:
		return fmt.Sprintf("%.1fK", float64(loc)/1000)
	default:
		return fmt.Sprintf("%.1fM", float64(loc)/1_000_000)
	}
}

func RenderInfo(info git.Info, size string, fileCount int, languages []git.LanguageStat, contributors []git.Contributor, loc int, lastActivity string) string {
	// Build language summary
	var langParts []string
	for i, l := range languages {
		if i >= 3 {
			break
		}
		langParts = append(langParts, fmt.Sprintf("%s %.1f%%", l.Name, l.Percentage))
	}
	langSummary := strings.Join(langParts, ", ")
	if langSummary == "" {
		langSummary = "-"
	}

	rows := []string{
		row("Repository:", titleStyle.Render(info.RepoName)),
		row("Branch:", fmt.Sprintf("%s %s", info.Branch, dimStyle.Render(fmt.Sprintf("(%s commits)", info.CommitCount)))),
		row("Head:", fmt.Sprintf("%s %s", dimStyle.Render(info.CommitHash), info.LastCommitMessage)),
		row("Created:", info.Created),
		row("Last active:", lastActivity),
		row("Languages:", langSummary),
		row("Size:", fmt.Sprintf("%s %s", size, dimStyle.Render(fmt.Sprintf("(%d files)", fileCount)))),
		row("Lines:", formatLOC(loc)),
	}

	// Contributors
	if len(contributors) > 0 {
		var contribParts []string
		for _, c := range contributors {
			contribParts = append(contribParts, fmt.Sprintf("%s %s", c.Name, dimStyle.Render(fmt.Sprintf("(%d)", c.Commits))))
		}
		rows = append(rows, row("Authors:", strings.Join(contribParts, ", ")))
	}

	if info.RemoteURL != "" {
		rows = append(rows, row("URL:", info.RemoteURL))
	}

	statusStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#3FB950"))
	if info.Status != "clean" {
		statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F85149"))
	}
	rows = append(rows, row("Status:", statusStyle.Render(info.Status)))

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func RenderLanguageBar(languages []git.LanguageStat, width int) string {
	if len(languages) == 0 {
		return ""
	}

	barWidth := width
	if barWidth <= 0 {
		barWidth = 50
	}

	// Build the colored bar
	var bar strings.Builder
	remaining := barWidth
	for i, lang := range languages {
		w := int(math.Round(lang.Percentage / 100.0 * float64(barWidth)))
		if i == len(languages)-1 {
			w = remaining // give last language whatever is left
		}
		if w <= 0 {
			continue
		}
		if w > remaining {
			w = remaining
		}
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(lang.Color))
		bar.WriteString(style.Render(strings.Repeat("█", w)))
		remaining -= w
	}

	// Build the legend
	var legend strings.Builder
	for i, lang := range languages {
		if i >= 5 {
			break
		}
		if i > 0 {
			legend.WriteString(dimStyle.Render("  "))
		}
		dot := lipgloss.NewStyle().Foreground(lipgloss.Color(lang.Color)).Render("●")
		legend.WriteString(fmt.Sprintf("%s %s %s", dot, lang.Name, dimStyle.Render(fmt.Sprintf("%.1f%%", lang.Percentage))))
	}

	return fmt.Sprintf("\n%s\n%s", bar.String(), legend.String())
}

func RenderLayout(logoBlock, info string) string {
	gap := "   "
	return lipgloss.JoinHorizontal(lipgloss.Top, logoBlock, gap, info)
}
