package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var greenLevels = []string{
	"#2d333b", // 0: no commits (visible gray on dark bg)
	"#0e4429", // 1: low
	"#006d32", // 2: medium-low
	"#26a641", // 3: medium-high
	"#39d353", // 4: high
}

const block = "█"

func commitLevel(count, max int) int {
	if count == 0 || max == 0 {
		return 0
	}
	ratio := float64(count) / float64(max)
	switch {
	case ratio <= 0.25:
		return 1
	case ratio <= 0.50:
		return 2
	case ratio <= 0.75:
		return 3
	default:
		return 4
	}
}

func colorBlock(level int) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(greenLevels[level])).Render(block)
}

type heatmapCell struct {
	date  time.Time
	count int
	valid bool
}

func RenderHeatmap(dates []string) string {
	counts := make(map[string]int)
	for _, d := range dates {
		counts[d]++
	}

	today := time.Now().Truncate(24 * time.Hour)
	oneYearAgo := today.AddDate(0, 0, -364)

	maxCommits := 0
	for d := oneYearAgo; !d.After(today); d = d.AddDate(0, 0, 1) {
		if c := counts[d.Format("2006-01-02")]; c > maxCommits {
			maxCommits = c
		}
	}

	// Align start to Sunday
	start := oneYearAgo
	for start.Weekday() != time.Sunday {
		start = start.AddDate(0, 0, -1)
	}

	// Build grid: 7 rows (Sun..Sat) x N week-columns
	var weeks [][]heatmapCell
	d := start
	for !d.After(today) {
		var week []heatmapCell
		for i := 0; i < 7; i++ {
			inRange := !d.Before(oneYearAgo) && !d.After(today)
			c := 0
			if inRange {
				c = counts[d.Format("2006-01-02")]
			}
			week = append(week, heatmapCell{date: d, count: c, valid: inRange})
			d = d.AddDate(0, 0, 1)
		}
		weeks = append(weeks, week)
	}

	// Build month labels at correct character positions
	// Each week column = 2 chars wide (block + space)
	// Weekday label column = 5 chars wide ("Mon  ")
	monthLabelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#8B949E"))
	dayLabelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#8B949E"))

	// Figure out which week index each month starts at
	type monthPos struct {
		name string
		col  int // week index
	}
	var monthPositions []monthPos
	lastMonth := time.Month(0)
	for wi, week := range weeks {
		for _, c := range week {
			if c.valid && c.date.Month() != lastMonth {
				lastMonth = c.date.Month()
				monthPositions = append(monthPositions, monthPos{c.date.Format("Jan"), wi})
				break
			}
		}
	}

	// Build month label row using character positions
	// Each week = 2 chars. Labels are 3 chars. Place them so they don't overlap.
	totalWidth := 5 + len(weeks)*2
	monthLine := make([]byte, totalWidth)
	for i := range monthLine {
		monthLine[i] = ' '
	}
	lastPos := -10 // track last label position to prevent overlap
	for _, mp := range monthPositions {
		pos := 5 + mp.col*2
		name := mp.name
		// Skip if it would overlap with previous label (need 4 char gap)
		if pos-lastPos < 4 {
			continue
		}
		if pos+len(name) <= totalWidth {
			copy(monthLine[pos:], name)
			lastPos = pos
		}
	}
	monthRow := monthLabelStyle.Render(string(monthLine))

	// Weekday labels (Mon, Wed, Fri like GitHub)
	weekdayLabels := map[int]string{
		1: "Mon",
		3: "Wed",
		5: "Fri",
	}

	var rows []string
	rows = append(rows, monthRow)

	for dayIdx := 0; dayIdx < 7; dayIdx++ {
		var row strings.Builder

		if label, ok := weekdayLabels[dayIdx]; ok {
			row.WriteString(dayLabelStyle.Render(fmt.Sprintf(" %s ", label)))
		} else {
			row.WriteString("     ")
		}

		for _, week := range weeks {
			if dayIdx < len(week) {
				c := week[dayIdx]
				if c.valid {
					level := commitLevel(c.count, maxCommits)
					row.WriteString(colorBlock(level) + " ")
				} else {
					row.WriteString("  ")
				}
			}
		}

		rows = append(rows, row.String())
	}

	// Legend
	legendStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#8B949E"))
	var legend strings.Builder
	legend.WriteString("     ")
	legend.WriteString(legendStyle.Render("Less "))
	for i := 0; i < 5; i++ {
		legend.WriteString(colorBlock(i) + " ")
	}
	legend.WriteString(legendStyle.Render("More"))

	rows = append(rows, "")
	rows = append(rows, legend.String())

	heatmapTitle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6CB6FF")).Render("Commit Activity (past year)")

	return fmt.Sprintf("\n%s\n%s", heatmapTitle, strings.Join(rows, "\n"))
}
