package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Info struct {
	Branch            string
	CommitHash        string
	CommitCount       string
	UserName          string
	UserEmail         string
	RemoteURL         string
	LastCommitMessage string
	Status            string
	RepoName          string
	Created           string
}

type LanguageStat struct {
	Name       string
	Percentage float64
	Color      string
}

type Contributor struct {
	Name    string
	Commits int
}

type HotFile struct {
	Path    string
	Changes int
}

type BranchHealth struct {
	TotalBranches int
	StaleBranches int // >30 days without commits
	AheadBehind   string
}

type Velocity struct {
	PerWeek   float64
	Sparkline string
	Trend     string // "↑", "↓", "→"
}

func runGit(args ...string) (string, error) {
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func GetInfo() (Info, error) {
	var info Info
	var err error

	info.Branch, err = runGit("rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return info, err
	}

	hash, _ := runGit("rev-parse", "HEAD")
	if len(hash) > 7 {
		info.CommitHash = hash[:7]
	} else {
		info.CommitHash = hash
	}

	info.CommitCount, _ = runGit("rev-list", "--count", "HEAD")
	info.UserName, _ = runGit("config", "user.name")
	info.UserEmail, _ = runGit("config", "user.email")
	info.RemoteURL, _ = runGit("config", "--get", "remote.origin.url")
	info.LastCommitMessage, _ = runGit("log", "-1", "--pretty=%s")
	info.RepoName = extractRepoName(info.RemoteURL)
	info.Created = getRepoAge()
	info.Status = getStatusSummary()

	return info, nil
}

func extractRepoName(remoteURL string) string {
	if remoteURL == "" {
		if dir, err := os.Getwd(); err == nil {
			return filepath.Base(dir)
		}
		return "unknown"
	}
	u := strings.TrimSuffix(remoteURL, ".git")

	// Handle git@host:owner/repo
	if strings.HasPrefix(u, "git@") {
		if idx := strings.Index(u, ":"); idx >= 0 {
			return u[idx+1:]
		}
	}

	// Handle any https URL — extract last two path segments (owner/repo)
	// Works for github.com, gitlab.com, bitbucket.org, self-hosted, etc.
	u = strings.TrimPrefix(u, "https://")
	u = strings.TrimPrefix(u, "http://")
	parts := strings.Split(u, "/")
	if len(parts) >= 3 {
		// For Bitbucket SCM URLs like host/bitbucket/scm/project/repo
		// or standard host/owner/repo
		return parts[len(parts)-2] + "/" + parts[len(parts)-1]
	}
	if len(parts) == 2 {
		return parts[1]
	}
	return u
}

func getRepoAge() string {
	out, err := runGit("log", "--reverse", "--format=%ci", "--diff-filter=A")
	if err != nil {
		return "unknown"
	}
	lines := strings.Split(out, "\n")
	if len(lines) == 0 || lines[0] == "" {
		return "unknown"
	}
	t, err := time.Parse("2006-01-02 15:04:05 -0700", lines[0])
	if err != nil {
		return "unknown"
	}
	return timeAgo(t)
}

func timeAgo(t time.Time) string {
	dur := time.Since(t)
	days := int(dur.Hours() / 24)
	switch {
	case days < 1:
		return "today"
	case days == 1:
		return "1 day ago"
	case days < 30:
		return fmt.Sprintf("%d days ago", days)
	case days < 60:
		return "1 month ago"
	case days < 365:
		return fmt.Sprintf("%d months ago", days/30)
	case days < 730:
		return "1 year ago"
	default:
		return fmt.Sprintf("%d years ago", days/365)
	}
}

func getStatusSummary() string {
	out, _ := runGit("status", "--short")
	if out == "" {
		return "clean"
	}
	lines := strings.Split(out, "\n")
	modified, added, deleted, untracked := 0, 0, 0, 0
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		code := line[:2]
		switch {
		case strings.Contains(code, "?"):
			untracked++
		case strings.Contains(code, "D"):
			deleted++
		case strings.Contains(code, "A"):
			added++
		default:
			modified++
		}
	}
	var parts []string
	if modified > 0 {
		parts = append(parts, fmt.Sprintf("%d modified", modified))
	}
	if added > 0 {
		parts = append(parts, fmt.Sprintf("%d added", added))
	}
	if deleted > 0 {
		parts = append(parts, fmt.Sprintf("%d deleted", deleted))
	}
	if untracked > 0 {
		parts = append(parts, fmt.Sprintf("%d untracked", untracked))
	}
	return strings.Join(parts, ", ")
}

// Language color mapping
var languageColors = map[string]string{
	"Go":         "#00ADD8",
	"Python":     "#3572A5",
	"JavaScript": "#F7DF1E",
	"TypeScript": "#3178C6",
	"Rust":       "#DEA584",
	"Java":       "#B07219",
	"C":          "#555555",
	"C++":        "#F34B7D",
	"C#":         "#239120",
	"Ruby":       "#CC342D",
	"PHP":        "#4F5D95",
	"Swift":      "#FA7343",
	"Kotlin":     "#A97BFF",
	"Shell":      "#89E051",
	"HTML":       "#E34C26",
	"CSS":        "#563D7C",
	"Lua":        "#000080",
	"Dart":       "#00B4AB",
	"Zig":        "#EC915C",
	"Haskell":    "#5E5086",
	"Elixir":     "#6E4A7E",
	"Scala":      "#DC322F",
	"Other":      "#8B8B8B",
}

var extToLang = map[string]string{
	".go":    "Go",
	".py":    "Python",
	".js":    "JavaScript",
	".ts":    "TypeScript",
	".tsx":   "TypeScript",
	".jsx":   "JavaScript",
	".rs":    "Rust",
	".java":  "Java",
	".c":     "C",
	".h":     "C",
	".cpp":   "C++",
	".cc":    "C++",
	".hpp":   "C++",
	".cs":    "C#",
	".rb":    "Ruby",
	".php":   "PHP",
	".swift": "Swift",
	".kt":    "Kotlin",
	".sh":    "Shell",
	".bash":  "Shell",
	".zsh":   "Shell",
	".html":  "HTML",
	".css":   "CSS",
	".scss":  "CSS",
	".lua":   "Lua",
	".dart":  "Dart",
	".zig":   "Zig",
	".hs":    "Haskell",
	".ex":    "Elixir",
	".exs":   "Elixir",
	".scala": "Scala",
}

func GetLanguageStats() []LanguageStat {
	// Use git ls-files to only count tracked files
	out, err := runGit("ls-files")
	if err != nil {
		return nil
	}

	langBytes := make(map[string]int64)
	var totalBytes int64

	for _, file := range strings.Split(out, "\n") {
		if file == "" {
			continue
		}
		ext := strings.ToLower(filepath.Ext(file))
		lang, ok := extToLang[ext]
		if !ok {
			continue // skip non-code files
		}
		info, err := os.Stat(file)
		if err != nil {
			continue
		}
		size := info.Size()
		langBytes[lang] += size
		totalBytes += size
	}

	if totalBytes == 0 {
		return nil
	}

	var stats []LanguageStat
	for lang, bytes := range langBytes {
		pct := float64(bytes) / float64(totalBytes) * 100
		color := languageColors[lang]
		if color == "" {
			color = languageColors["Other"]
		}
		stats = append(stats, LanguageStat{Name: lang, Percentage: pct, Color: color})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Percentage > stats[j].Percentage
	})

	return stats
}

func GetRepoSize() (string, int) {
	out, err := runGit("ls-files")
	if err != nil {
		return "0 B", 0
	}
	files := strings.Split(strings.TrimSpace(out), "\n")
	if len(files) == 1 && files[0] == "" {
		return "0 B", 0
	}

	var totalSize int64
	for _, f := range files {
		if info, err := os.Stat(f); err == nil {
			totalSize += info.Size()
		}
	}

	var sizeStr string
	switch {
	case totalSize < 1024:
		sizeStr = fmt.Sprintf("%d B", totalSize)
	case totalSize < 1024*1024:
		sizeStr = fmt.Sprintf("%.1f KB", float64(totalSize)/1024)
	default:
		sizeStr = fmt.Sprintf("%.1f MB", float64(totalSize)/(1024*1024))
	}

	return sizeStr, len(files)
}

func GetContributors(max int) []Contributor {
	out, err := runGit("shortlog", "-sn", "--no-merges", "HEAD")
	if err != nil {
		return nil
	}
	var contributors []Contributor
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// Format: "  123\tName"
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) != 2 {
			continue
		}
		count := 0
		fmt.Sscanf(strings.TrimSpace(parts[0]), "%d", &count)
		contributors = append(contributors, Contributor{
			Name:    strings.TrimSpace(parts[1]),
			Commits: count,
		})
		if len(contributors) >= max {
			break
		}
	}
	return contributors
}

func GetLinesOfCode() int {
	out, err := runGit("ls-files")
	if err != nil {
		return 0
	}
	total := 0
	for _, file := range strings.Split(out, "\n") {
		if file == "" {
			continue
		}
		ext := strings.ToLower(filepath.Ext(file))
		if _, ok := extToLang[ext]; !ok {
			continue // only count code files
		}
		data, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		lines := strings.Count(string(data), "\n")
		if len(data) > 0 && data[len(data)-1] != '\n' {
			lines++ // count last line if no trailing newline
		}
		total += lines
	}
	return total
}

func GetLastActivity() string {
	out, err := runGit("log", "-1", "--format=%ci")
	if err != nil {
		return "unknown"
	}
	t, err := time.Parse("2006-01-02 15:04:05 -0700", out)
	if err != nil {
		return "unknown"
	}
	return timeAgo(t)
}

var generatedFiles = map[string]bool{
	"Cargo.lock":        true,
	"package-lock.json": true,
	"yarn.lock":         true,
	"pnpm-lock.yaml":   true,
	"go.sum":            true,
	"Gemfile.lock":      true,
	"composer.lock":     true,
	"Pipfile.lock":      true,
	"poetry.lock":       true,
}

func isGenerated(path string) bool {
	base := filepath.Base(path)
	if generatedFiles[base] {
		return true
	}
	if strings.HasSuffix(base, "-hashes.json") || strings.HasSuffix(base, ".min.js") || strings.HasSuffix(base, ".min.css") {
		return true
	}
	return false
}

func GetHotFiles(max int) []HotFile {
	// Most frequently changed files in the last 90 days
	out, err := runGit("log", "--since=90 days ago", "--pretty=format:", "--name-only")
	if err != nil {
		return nil
	}

	counts := make(map[string]int)
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || isGenerated(line) {
			continue
		}
		counts[line]++
	}

	type kv struct {
		path  string
		count int
	}
	var sorted []kv
	for p, c := range counts {
		sorted = append(sorted, kv{p, c})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].count > sorted[j].count
	})

	var result []HotFile
	for i, item := range sorted {
		if i >= max {
			break
		}
		result = append(result, HotFile{Path: item.path, Changes: item.count})
	}
	return result
}

func GetVelocity() Velocity {
	// Get weekly commit counts for the last 8 weeks
	var weeklyCounts []int
	now := time.Now()
	for i := 7; i >= 0; i-- {
		weekEnd := now.AddDate(0, 0, -7*i)
		weekStart := weekEnd.AddDate(0, 0, -7)
		out, err := runGit("rev-list", "--count", "--since="+weekStart.Format("2006-01-02"), "--until="+weekEnd.Format("2006-01-02"), "HEAD")
		if err != nil {
			weeklyCounts = append(weeklyCounts, 0)
			continue
		}
		count := 0
		fmt.Sscanf(out, "%d", &count)
		weeklyCounts = append(weeklyCounts, count)
	}

	// Calculate average
	total := 0
	for _, c := range weeklyCounts {
		total += c
	}
	avg := float64(total) / float64(len(weeklyCounts))

	// Build sparkline
	sparkChars := []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}
	maxCount := 0
	for _, c := range weeklyCounts {
		if c > maxCount {
			maxCount = c
		}
	}
	var spark strings.Builder
	for _, c := range weeklyCounts {
		idx := 0
		if maxCount > 0 {
			idx = int(float64(c) / float64(maxCount) * 7)
			if idx > 7 {
				idx = 7
			}
		}
		spark.WriteRune(sparkChars[idx])
	}

	// Trend: compare last 4 weeks vs first 4 weeks
	trend := "→"
	if len(weeklyCounts) == 8 {
		firstHalf := 0
		secondHalf := 0
		for i := 0; i < 4; i++ {
			firstHalf += weeklyCounts[i]
			secondHalf += weeklyCounts[i+4]
		}
		if secondHalf > firstHalf+2 {
			trend = "↑"
		} else if firstHalf > secondHalf+2 {
			trend = "↓"
		}
	}

	return Velocity{PerWeek: avg, Sparkline: spark.String(), Trend: trend}
}

func GetDependencyCount() (string, int) {
	// Find repo root so we look for dep files in the right place
	root, err := runGit("rev-parse", "--show-toplevel")
	if err != nil {
		root = "."
	}

	// Detect package manager and count dependencies
	depFiles := []struct {
		file    string
		manager string
		counter func(string) int
	}{
		{"go.mod", "Go modules", countGoMod},
		{"package.json", "npm", countPackageJSON},
		{"requirements.txt", "pip", countLines},
		{"Pipfile", "pipenv", countLines},
		{"Cargo.toml", "Cargo", countCargoToml},
		{"Gemfile", "Bundler", countLines},
		{"composer.json", "Composer", countComposerJSON},
		{"pyproject.toml", "pyproject", countPyprojectToml},
	}

	for _, dep := range depFiles {
		path := filepath.Join(root, dep.file)
		if data, err := os.ReadFile(path); err == nil {
			count := dep.counter(string(data))
			if count > 0 {
				return dep.manager, count
			}
		}
	}
	return "", 0
}

func countGoMod(content string) int {
	count := 0
	inRequire := false
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line == "require (" {
			inRequire = true
			continue
		}
		if line == ")" {
			inRequire = false
			continue
		}
		if inRequire && line != "" && !strings.HasPrefix(line, "//") {
			count++
		}
	}
	return count
}

func countPackageJSON(content string) int {
	// Simple count of lines in "dependencies" and "devDependencies" blocks
	count := 0
	inDeps := false
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, `"dependencies"`) || strings.Contains(trimmed, `"devDependencies"`) {
			inDeps = true
			continue
		}
		if inDeps && strings.HasPrefix(trimmed, "}") {
			inDeps = false
			continue
		}
		if inDeps && strings.Contains(trimmed, `"`) {
			count++
		}
	}
	return count
}

func countLines(content string) int {
	count := 0
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			count++
		}
	}
	return count
}

func countCargoToml(content string) int {
	count := 0
	inDeps := false
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "[dependencies]") || strings.Contains(trimmed, "[dev-dependencies]") {
			inDeps = true
			continue
		}
		if strings.HasPrefix(trimmed, "[") {
			inDeps = false
			continue
		}
		if inDeps && trimmed != "" && !strings.HasPrefix(trimmed, "#") && strings.Contains(trimmed, "=") {
			count++
		}
	}
	return count
}

func countComposerJSON(content string) int {
	return countPackageJSON(content) // same structure
}

func countPyprojectToml(content string) int {
	count := 0
	inDeps := false
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "dependencies") && strings.Contains(trimmed, "[") {
			inDeps = true
			continue
		}
		if inDeps && strings.HasPrefix(trimmed, "]") {
			inDeps = false
			continue
		}
		if inDeps && strings.HasPrefix(trimmed, `"`) {
			count++
		}
	}
	return count
}

func GetBranchHealth() BranchHealth {
	var health BranchHealth

	// Total branches
	out, err := runGit("branch", "-a")
	if err != nil {
		return health
	}
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line != "" && !strings.Contains(line, "->") {
			health.TotalBranches++
		}
	}

	// Stale branches (local branches with no commits in 30+ days)
	out, err = runGit("branch", "--format=%(refname:short) %(committerdate:iso)")
	if err == nil {
		thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
		for _, line := range strings.Split(out, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			parts := strings.SplitN(line, " ", 2)
			if len(parts) < 2 {
				continue
			}
			dateStr := strings.TrimSpace(parts[1])
			if t, err := time.Parse("2006-01-02 15:04:05 -0700", dateStr); err == nil {
				if t.Before(thirtyDaysAgo) {
					health.StaleBranches++
				}
			}
		}
	}

	// Ahead/behind default branch
	defaultBranch := getDefaultBranch()
	currentBranch, _ := runGit("rev-parse", "--abbrev-ref", "HEAD")
	if defaultBranch != "" && currentBranch != defaultBranch {
		out, err = runGit("rev-list", "--left-right", "--count", defaultBranch+"...HEAD")
		if err == nil {
			parts := strings.Fields(out)
			if len(parts) == 2 {
				health.AheadBehind = fmt.Sprintf("↑%s ↓%s vs %s", parts[1], parts[0], defaultBranch)
			}
		}
	}

	return health
}

func getDefaultBranch() string {
	// Try origin/HEAD first
	out, err := runGit("symbolic-ref", "refs/remotes/origin/HEAD")
	if err == nil {
		return strings.TrimPrefix(out, "refs/remotes/origin/")
	}
	// Fallback: check if main or master exists
	for _, branch := range []string{"main", "master"} {
		if _, err := runGit("rev-parse", "--verify", branch); err == nil {
			return branch
		}
	}
	return ""
}

func GetLicense() string {
	root, err := runGit("rev-parse", "--show-toplevel")
	if err != nil {
		root = "."
	}

	licenseFiles := []string{"LICENSE", "LICENSE.md", "LICENSE.txt", "LICENCE", "LICENCE.md", "COPYING", "COPYING.md"}
	for _, name := range licenseFiles {
		data, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			continue
		}
		content := strings.ToLower(string(data))
		switch {
		case strings.Contains(content, "mit license") || strings.Contains(content, "permission is hereby granted, free of charge"):
			return "MIT"
		case strings.Contains(content, "apache license") && strings.Contains(content, "version 2.0"):
			return "Apache-2.0"
		case strings.Contains(content, "gnu general public license") && strings.Contains(content, "version 3"):
			return "GPL-3.0"
		case strings.Contains(content, "gnu general public license") && strings.Contains(content, "version 2"):
			return "GPL-2.0"
		case strings.Contains(content, "gnu lesser general public license"):
			return "LGPL"
		case strings.Contains(content, "bsd 3-clause") || strings.Contains(content, "redistribution and use in source and binary forms"):
			return "BSD-3-Clause"
		case strings.Contains(content, "bsd 2-clause"):
			return "BSD-2-Clause"
		case strings.Contains(content, "mozilla public license") && strings.Contains(content, "2.0"):
			return "MPL-2.0"
		case strings.Contains(content, "unlicense"):
			return "Unlicense"
		case strings.Contains(content, "isc license"):
			return "ISC"
		default:
			return "Custom"
		}
	}
	return ""
}

func GetLatestTag() string {
	tag, err := runGit("describe", "--tags", "--abbrev=0")
	if err != nil {
		return ""
	}
	return tag
}

func CleanURL(raw string) string {
	u := strings.TrimSuffix(raw, ".git")
	u = strings.TrimPrefix(u, "https://")
	u = strings.TrimPrefix(u, "http://")
	return u
}

func GetCommitDates() ([]string, error) {
	out, err := runGit("log", "--pretty=format:%cd", "--date=short")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out, "\n")
	var dates []string
	for _, line := range lines {
		if line != "" {
			dates = append(dates, line)
		}
	}
	return dates, nil
}
