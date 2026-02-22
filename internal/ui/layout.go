package ui

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/fayssal-elmofatiche/gtop/internal/git"
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
 ╔═╗ ╔╦╗ ╔═╗ ╔═╗
 ║ ╦  ║  ║ ║ ╠═╝
 ╚═╝  ╩  ╚═╝ ╩  `

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

func RenderInfo(info git.Info, size string, fileCount int, languages []git.LanguageStat, loc int, lastActivity string, velocity git.Velocity, depManager string, depCount int, health git.BranchHealth) string {
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

	if info.RemoteURL != "" {
		rows = append(rows, row("URL:", info.RemoteURL))
	}

	// Velocity
	if velocity.Sparkline != "" {
		trendStyle := dimStyle
		if velocity.Trend == "↑" {
			trendStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#3FB950"))
		} else if velocity.Trend == "↓" {
			trendStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F85149"))
		}
		rows = append(rows, row("Velocity:", fmt.Sprintf("%.1f/wk %s %s", velocity.PerWeek, velocity.Sparkline, trendStyle.Render(velocity.Trend))))
	}

	// Dependencies
	if depCount > 0 {
		rows = append(rows, row("Deps:", fmt.Sprintf("%d %s", depCount, dimStyle.Render("("+depManager+")"))))
	}

	// Branch health
	if health.TotalBranches > 0 {
		branchStr := fmt.Sprintf("%d", health.TotalBranches)
		if health.StaleBranches > 0 {
			staleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F85149"))
			branchStr += fmt.Sprintf(" %s", staleStyle.Render(fmt.Sprintf("(%d stale)", health.StaleBranches)))
		}
		if health.AheadBehind != "" {
			branchStr += " " + dimStyle.Render(health.AheadBehind)
		}
		rows = append(rows, row("Branches:", branchStr))
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

func RenderContributors(contributors []git.Contributor) string {
	if len(contributors) == 0 {
		return ""
	}

	header := titleStyle.Render("Top Authors")
	var lines []string
	lines = append(lines, header)

	maxCommits := contributors[0].Commits
	barMax := 20

	for _, c := range contributors {
		w := int(float64(c.Commits) / float64(maxCommits) * float64(barMax))
		if w < 1 {
			w = 1
		}
		bar := lipgloss.NewStyle().Foreground(lipgloss.Color("#6CB6FF")).Render(strings.Repeat("█", w))
		count := dimStyle.Render(fmt.Sprintf("%5d", c.Commits))
		lines = append(lines, fmt.Sprintf("  %s %s %s", count, bar, valueStyle.Render(c.Name)))
	}
	return "\n" + strings.Join(lines, "\n")
}

func RenderHotFiles(files []git.HotFile) string {
	if len(files) == 0 {
		return ""
	}

	header := titleStyle.Render("Hot Files") + dimStyle.Render(" (90 days)")
	var lines []string
	lines = append(lines, header)

	maxChanges := files[0].Changes
	barMax := 20

	for _, f := range files {
		w := int(float64(f.Changes) / float64(maxChanges) * float64(barMax))
		if w < 1 {
			w = 1
		}
		bar := lipgloss.NewStyle().Foreground(lipgloss.Color("#F0883E")).Render(strings.Repeat("█", w))
		count := dimStyle.Render(fmt.Sprintf("%3d", f.Changes))
		lines = append(lines, fmt.Sprintf("  %s %s %s", count, bar, valueStyle.Render(f.Path)))
	}
	return "\n" + strings.Join(lines, "\n")
}

func RenderLayout(logoBlock, info string) string {
	gap := "   "
	return lipgloss.JoinHorizontal(lipgloss.Top, logoBlock, gap, info)
}
