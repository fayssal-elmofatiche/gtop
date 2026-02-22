# gtop

A fast, zero-config, terminal-based git repository information tool. Think [onefetch](https://github.com/o2sh/onefetch) meets [neofetch](https://github.com/dylanaraps/neofetch) — displays a rich summary of any git repo at a glance.

## Features

- **Repository info** — branch, head commit, remote URL, working tree status
- **Language breakdown** — colored proportional bar with percentages (weighted by file size)
- **Top contributors** — bar chart of most active authors by commit count
- **Lines of code** — total across all detected source files
- **Repo age & last activity** — human-readable timestamps
- **Commit velocity** — average commits/week with sparkline trend over the last 8 weeks
- **Dependency count** — auto-detects Go modules, npm, pip, Cargo, Bundler, Composer, and more
- **Branch health** — total branches, stale branch count (>30 days), ahead/behind default branch
- **Hot files** — most frequently changed files in the last 90 days with proportional bars
- **Commit heatmap** — GitHub-style contribution graph for the past year (7-row daily grid, 5 intensity levels)

## Install

### Homebrew (macOS/Linux)

```bash
brew install fayssal-elmofatiche/tap/gtop
```

### Go install (requires Go 1.23+)

```bash
go install github.com/fayssal-elmofatiche/gtop/cmd/gtop@latest
```

### Download binary

Pre-built binaries for Linux, macOS, and Windows (amd64/arm64) are available on the [Releases](https://github.com/fayssal-elmofatiche/gtop/releases) page.

### Build from source

```bash
git clone https://github.com/fayssal-elmofatiche/gtop.git
cd gtop
go build -o gtop ./cmd/gtop
```

## Usage

```bash
cd /path/to/any/git/repo
gtop
```

No flags, no config. Just run `gtop` inside a git repository.

```bash
gtop --version   # print version
```

## Example Output

```
╭─────────────────────────────╮   Repository:   owner/repo
│                             │   Branch:       main (142 commits)
│     ____ _____ ___  ____    │   Head:         a1b2c3d add new feature
│    / ___|_   _/ _ \|  _ \   │   Created:      8 months ago
│   | |  _  | || | | | |_) |  │   Last active:  today
│   | |_| | | || |_| |  __/   │   Languages:    Go 85.2%, Shell 14.8%
│    \____| |_| \___/|_|      │   Size:         48.3 KB (12 files)
╰─────────────────────────────╯   Lines:        1.2K
                                  URL:          github.com/owner/repo.git
                                  Velocity:     2.4/wk ▂▃▅▇█▆▃▄ ↑
                                  Deps:         12 (Go modules)
                                  Branches:     5 (1 stale)
                                  Status:       clean

██████████████████████████████████████████████████
● Go 85.2%  ● Shell 14.8%

Top Authors
    98 ████████████████████ Alice
    44 ████████             Bob
    12 ██                   Charlie

Hot Files (90 days)
   18 ████████████████████ internal/git/git.go
   12 █████████████        cmd/gtop/main.go
    7 ███████              internal/ui/layout.go

Commit Activity (past year)
     Mar     Apr     May     Jun     Jul     Aug     ...
 Mon █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
 Wed █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
 Fri █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
     Less █ █ █ █ █ More
```

## How It Works

gtop is a single-binary CLI tool written in Go. It gathers all data by shelling out to `git` commands — no external libraries for git interaction, no indexing, no caching. This keeps the tool simple and ensures it works with any git version.

**Architecture:**

| Package | Role |
| --- | --- |
| `cmd/gtop/` | Entry point — orchestrates data gathering and output |
| `internal/git/` | All git data extraction via `exec.Command("git", ...)` |
| `internal/ui/` | Terminal rendering — layout, heatmap, bar charts |

**Tech stack:**

- **[Go](https://go.dev/)** — fast compilation, single static binary, cross-platform
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** — styled terminal output (colors, borders, layout) from the Charm ecosystem
- **[goreleaser](https://goreleaser.com/)** — cross-compilation and release automation (Linux/macOS/Windows, amd64/arm64)
- **git CLI** — all repository data is gathered through standard `git` commands (`log`, `shortlog`, `ls-files`, `status`, `rev-list`, etc.)

**Key design decisions:**

- No interactive TUI — gtop is a one-shot display tool (like neofetch), not a persistent UI (like htop)
- No external git library — shelling out to `git` keeps the binary small and avoids CGO dependencies
- Language detection by file extension weighted by byte size — simple heuristic, no tree-sitter or deep parsing
- All sections are conditionally rendered — if there are no contributors, deps, or hot files, those sections are silently omitted

## Supported Languages

Go, Python, JavaScript, TypeScript, Rust, Java, C, C++, C#, Ruby, PHP, Swift, Kotlin, Shell, HTML, CSS, Lua, Dart, Zig, Haskell, Elixir, Scala

## Supported Package Managers

| File | Manager |
| --- | --- |
| `go.mod` | Go modules |
| `package.json` | npm |
| `requirements.txt` | pip |
| `Pipfile` | pipenv |
| `pyproject.toml` | pyproject |
| `Cargo.toml` | Cargo |
| `Gemfile` | Bundler |
| `composer.json` | Composer |

## Requirements

- **git** must be installed and available in `PATH`
- Works on any terminal that supports ANSI colors (most modern terminals)
- Colors are automatically disabled when output is piped

## License

MIT
