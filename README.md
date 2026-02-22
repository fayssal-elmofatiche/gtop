# gtop

A fast, terminal-based git repository information tool. Think [onefetch](https://github.com/o2sh/onefetch) meets [neofetch](https://github.com/dylanaraps/neofetch) — displays a summary of any git repo at a glance.

Built in Go with [lipgloss](https://github.com/charmbracelet/lipgloss) for styled terminal output.

## Features

- **Repository info** — branch, head commit, remote URL, working tree status
- **Language breakdown** — colored proportional bar with percentages (by file size)
- **Top contributors** — most active authors by commit count
- **Lines of code** — total across all detected source files
- **Repo age & last activity** — human-readable timestamps
- **Commit velocity** — average commits/week with sparkline trend (8 weeks)
- **Dependency count** — auto-detects Go modules, npm, pip, Cargo, Bundler, Composer, and more
- **Branch health** — total branches, stale branch count (>30 days), ahead/behind default branch
- **Hot files** — most frequently changed files in the last 90 days
- **Commit heatmap** — GitHub-style contribution graph for the past year

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

Pre-built binaries for Linux, macOS, and Windows are available on the [Releases](https://github.com/fayssal-elmofatiche/gtop/releases) page.

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

## Supported Languages

Go, Python, JavaScript, TypeScript, Rust, Java, C, C++, C#, Ruby, PHP, Swift, Kotlin, Shell, HTML, CSS, Lua, Dart, Zig, Haskell, Elixir, Scala

## Supported Package Managers

Go modules, npm, pip, pipenv, pyproject, Cargo, Bundler, Composer

## License

MIT
