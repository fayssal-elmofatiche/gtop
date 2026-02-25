# Contributing to gfetch

Thanks for your interest in contributing to gfetch!

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/gfetch.git`
3. Create a branch: `git checkout -b my-feature`
4. Make your changes
5. Build and test: `go build ./cmd/gfetch && ./gfetch`
6. Commit your changes
7. Push and open a pull request

## Development

**Requirements**: Go 1.23+

```bash
# Build
go build -o gfetch ./cmd/gfetch

# Run
./gfetch
```

## Project Structure

| Package | Role |
| --- | --- |
| `cmd/gfetch/` | Entry point — orchestrates data gathering and output |
| `internal/git/` | All git data extraction via `exec.Command("git", ...)` |
| `internal/ui/` | Terminal rendering — layout, heatmap, bar charts, logos |

## Guidelines

- Keep it simple — gfetch is a zero-config, single-binary tool
- No external git libraries — we shell out to `git` intentionally
- All sections are conditionally rendered — if data is empty, skip the section
- Test your changes on multiple repos before submitting

## Adding a New Language Logo

Logos live in `internal/ui/logos.go`. Each logo uses `{N}` tokens for color:

1. Add the ASCII art to the `logos` map
2. Add color hex codes to the `colors` slice
3. Add file extensions to `extToLang` in `internal/git/git.go` if needed

## Reporting Issues

Open an issue on GitHub with:
- Your OS and terminal emulator
- The output of `gfetch --version`
- What you expected vs what happened
