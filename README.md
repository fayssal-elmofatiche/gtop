# gtop

A fast, terminal-based git repository information tool. Think [onefetch](https://github.com/o2sh/onefetch) meets [neofetch](https://github.com/dylanaraps/neofetch) — displays a summary of any git repo at a glance.

Built in Go with [lipgloss](https://github.com/charmbracelet/lipgloss) for styled terminal output.

## Features

- Repository info (branch, head, remote URL, status)
- Language breakdown with colored bar
- Top contributors
- Lines of code
- Repo age and last activity
- GitHub-style commit activity heatmap (past year)

## Install

### From source (requires Go 1.23+)

```bash
go install github.com/fayssalelmofatiche/gtop/cmd/gtop@latest
```

### Build locally

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
                                  Authors:      Alice (98), Bob (44)
                                  URL:          github.com/owner/repo.git
                                  Status:       clean

██████████████████████████████████████████████████
● Go 85.2%  ● Shell 14.8%

Commit Activity (past year)
     Mar     Apr     May     Jun     Jul     ...
 Mon █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
 Wed █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
 Fri █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
     Less █ █ █ █ █ More
```

## Supported Languages

Go, Python, JavaScript, TypeScript, Rust, Java, C, C++, C#, Ruby, PHP, Swift, Kotlin, Shell, HTML, CSS, Lua, Dart, Zig, Haskell, Elixir, Scala

## License

MIT
