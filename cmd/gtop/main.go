package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/fayssal-elmofatiche/gtop/internal/git"
	"github.com/fayssal-elmofatiche/gtop/internal/ui"
)

var version = "dev"

func getVersion() string {
	if version != "dev" {
		return version
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return version
}

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Println("gtop", getVersion())
		return
	}

	gitInfo, err := git.GetInfo()
	if err != nil {
		fmt.Println("Not a git repository")
		os.Exit(1)
	}

	languages := git.GetLanguageStats()
	size, fileCount := git.GetRepoSize()
	contributors := git.GetContributors(5)
	loc := git.GetLinesOfCode()
	lastActivity := git.GetLastActivity()
	velocity := git.GetVelocity()
	depManager, depCount := git.GetDependencyCount()
	health := git.GetBranchHealth()
	hotFiles := git.GetHotFiles(5)

	logo := ui.RenderLogo()
	info := ui.RenderInfo(gitInfo, size, fileCount, languages, loc, lastActivity, velocity, depManager, depCount, health)
	fmt.Println(ui.RenderLayout(logo, info))

	fmt.Println(ui.RenderLanguageBar(languages, 50))

	if len(contributors) > 0 {
		fmt.Println(ui.RenderContributors(contributors))
	}

	if len(hotFiles) > 0 {
		fmt.Println(ui.RenderHotFiles(hotFiles))
	}

	dates, _ := git.GetCommitDates()
	if len(dates) > 0 {
		fmt.Println(ui.RenderHeatmap(dates))
	}
}
