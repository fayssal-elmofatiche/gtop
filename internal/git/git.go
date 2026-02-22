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
		// Fall back to directory name
		if dir, err := os.Getwd(); err == nil {
			return filepath.Base(dir)
		}
		return "unknown"
	}
	// Handle https://github.com/owner/repo.git
	u := strings.TrimSuffix(remoteURL, ".git")
	if idx := strings.Index(u, "github.com"); idx >= 0 {
		return strings.TrimPrefix(u[idx:], "github.com/")
	}
	// Handle git@host:owner/repo.git
	if idx := strings.LastIndex(u, ":"); idx >= 0 {
		return u[idx+1:]
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
