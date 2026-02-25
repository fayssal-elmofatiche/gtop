package ui

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/fayssal-elmofatiche/gfetch/internal/git"
)

var (
	labelStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6CB6FF")).Width(14)
	valueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#E6EDF3"))
	dimStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#8B949E"))
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#F0883E"))
)

// RenderParams holds all data needed to render the info panel.
type RenderParams struct {
	Info             git.Info
	Size             string
	FileCount        int
	Languages        []git.LanguageStat
	LOC              int
	LastActivity     string
	Velocity         git.Velocity
	DepManager       string
	DepCount         int
	Health           git.BranchHealth
	License          string
	LatestTag        string
	CICD             []string
	StashCount       int
	Contributors     int
	TestRatio        git.TestRatio
	CommitConvention string
}

func RenderLogo(language string) string {
	logo := getLanguageLogo(language)
	return renderColoredArt(logo.art, logo.colors)
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

func RenderInfo(p RenderParams) string {
	// Build language summary
	var langParts []string
	for i, l := range p.Languages {
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
		row("Repository:", titleStyle.Render(p.Info.RepoName)),
		row("Branch:", fmt.Sprintf("%s %s", p.Info.Branch, dimStyle.Render(fmt.Sprintf("(%s commits)", p.Info.CommitCount)))),
		row("Head:", fmt.Sprintf("%s %s", dimStyle.Render(p.Info.CommitHash), p.Info.LastCommitMessage)),
		row("Author:", fmt.Sprintf("%s %s", p.Info.UserName, dimStyle.Render(fmt.Sprintf("<%s>", p.Info.UserEmail)))),
		row("Created:", p.Info.Created),
		row("Last active:", p.LastActivity),
		row("Languages:", langSummary),
		row("Size:", fmt.Sprintf("%s %s", p.Size, dimStyle.Render(fmt.Sprintf("(%d files)", p.FileCount)))),
		row("Lines:", formatLOC(p.LOC)),
	}

	if p.Info.RemoteURL != "" {
		rows = append(rows, row("URL:", git.CleanURL(p.Info.RemoteURL)))
	}

	if p.Contributors > 0 {
		rows = append(rows, row("Authors:", fmt.Sprintf("%d", p.Contributors)))
	}

	if p.LatestTag != "" {
		rows = append(rows, row("Version:", p.LatestTag))
	}

	if p.License != "" {
		rows = append(rows, row("License:", p.License))
	}

	// Velocity
	if p.Velocity.Sparkline != "" {
		trendStyle := dimStyle
		if p.Velocity.Trend == "↑" {
			trendStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#3FB950"))
		} else if p.Velocity.Trend == "↓" {
			trendStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F85149"))
		}
		rows = append(rows, row("Velocity:", fmt.Sprintf("%.1f/wk %s %s", p.Velocity.PerWeek, p.Velocity.Sparkline, trendStyle.Render(p.Velocity.Trend))))
	}

	// Dependencies
	if p.DepCount > 0 {
		rows = append(rows, row("Deps:", fmt.Sprintf("%d %s", p.DepCount, dimStyle.Render("("+p.DepManager+")"))))
	}

	// Branch health
	if p.Health.TotalBranches > 0 {
		branchStr := fmt.Sprintf("%d", p.Health.TotalBranches)
		if p.Health.StaleBranches > 0 {
			staleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F85149"))
			branchStr += fmt.Sprintf(" %s", staleStyle.Render(fmt.Sprintf("(%d stale)", p.Health.StaleBranches)))
		}
		if p.Health.AheadBehind != "" {
			branchStr += " " + dimStyle.Render(p.Health.AheadBehind)
		}
		rows = append(rows, row("Branches:", branchStr))
	}

	// CI/CD
	if len(p.CICD) > 0 {
		rows = append(rows, row("CI/CD:", strings.Join(p.CICD, ", ")))
	}

	// Test ratio
	if p.TestRatio.TestLines > 0 {
		ratioStr := fmt.Sprintf("%.0f%% %s", p.TestRatio.Ratio*100, dimStyle.Render(fmt.Sprintf("(%s test / %s code)", formatLOC(p.TestRatio.TestLines), formatLOC(p.TestRatio.CodeLines))))
		rows = append(rows, row("Tests:", ratioStr))
	}

	// Commit convention
	if p.CommitConvention != "" {
		rows = append(rows, row("Commits:", p.CommitConvention))
	}

	// Stash
	if p.StashCount > 0 {
		stashStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#D2A8FF"))
		rows = append(rows, row("Stash:", stashStyle.Render(fmt.Sprintf("%d entries", p.StashCount))))
	}

	if p.Info.GitVersion != "" {
		rows = append(rows, row("Git:", p.Info.GitVersion))
	}

	statusStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#3FB950"))
	if p.Info.Status != "clean" {
		statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F85149"))
	}
	rows = append(rows, row("Status:", statusStyle.Render(p.Info.Status)))

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
			legend.WriteString("  ")
		}
		dotStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(lang.Color))
		legend.WriteString(fmt.Sprintf("%s %s %s", dotStyle.Render("●"), lang.Name, dimStyle.Render(fmt.Sprintf("%.1f%%", lang.Percentage))))
	}

	return fmt.Sprintf("\n%s\n%s", bar.String(), legend.String())
}

func RenderContributors(stats git.ContributorStats) string {
	if len(stats.Top) == 0 {
		return ""
	}

	totalLabel := ""
	if stats.Total > len(stats.Top) {
		totalLabel = dimStyle.Render(fmt.Sprintf(" (%d total)", stats.Total))
	}
	header := titleStyle.Render("Top Authors") + totalLabel
	var lines []string
	lines = append(lines, header)

	// Sum all commits across top authors for percentage calculation
	totalCommits := 0
	for _, c := range stats.Top {
		totalCommits += c.Commits
	}

	maxCommits := stats.Top[0].Commits
	barMax := 20
	barStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#6CB6FF"))

	for _, c := range stats.Top {
		w := int(float64(c.Commits) / float64(maxCommits) * float64(barMax))
		if w < 1 {
			w = 1
		}
		bar := barStyle.Render(strings.Repeat("█", w))
		pct := float64(c.Commits) / float64(totalCommits) * 100
		label := dimStyle.Render(fmt.Sprintf("%5.1f%%", pct))
		lines = append(lines, fmt.Sprintf("  %s %s %s %s", label, bar, valueStyle.Render(c.Name), dimStyle.Render(fmt.Sprintf("(%d)", c.Commits))))
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
	barStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F0883E"))

	for _, f := range files {
		w := int(float64(f.Changes) / float64(maxChanges) * float64(barMax))
		if w < 1 {
			w = 1
		}
		bar := barStyle.Render(strings.Repeat("█", w))
		count := dimStyle.Render(fmt.Sprintf("%3d", f.Changes))
		lines = append(lines, fmt.Sprintf("  %s %s %s", count, bar, valueStyle.Render(f.Path)))
	}
	return "\n" + strings.Join(lines, "\n")
}

func RenderReleases(releases []git.Release) string {
	if len(releases) == 0 {
		return ""
	}

	header := titleStyle.Render("Releases")
	var lines []string
	lines = append(lines, header)

	tagStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#3FB950")).Bold(true)
	for _, r := range releases {
		lines = append(lines, fmt.Sprintf("  %s %s", tagStyle.Render(r.Tag), dimStyle.Render(r.Age)))
	}
	return "\n" + strings.Join(lines, "\n")
}

func RenderLayout(logoBlock, info string) string {
	gap := "   "
	return lipgloss.JoinHorizontal(lipgloss.Top, logoBlock, gap, info)
}
