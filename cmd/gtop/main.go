package main

import (
	"fmt"
	"os"

	"github.com/fayssalelmofatiche/gtop/internal/git"
	"github.com/fayssalelmofatiche/gtop/internal/ui"
)

func main() {
	gitInfo, err := git.GetInfo()
	if err != nil {
		fmt.Println("Not a git repository")
		os.Exit(1)
	}

	languages := git.GetLanguageStats()
	size, fileCount := git.GetRepoSize()
	contributors := git.GetContributors(3)
	loc := git.GetLinesOfCode()
	lastActivity := git.GetLastActivity()
	velocity := git.GetVelocity()
	depManager, depCount := git.GetDependencyCount()
	health := git.GetBranchHealth()
	hotFiles := git.GetHotFiles(5)

	logo := ui.RenderLogo()
	info := ui.RenderInfo(gitInfo, size, fileCount, languages, contributors, loc, lastActivity, velocity, depManager, depCount, health)
	fmt.Println(ui.RenderLayout(logo, info))

	fmt.Println(ui.RenderLanguageBar(languages, 50))

	if len(hotFiles) > 0 {
		fmt.Println(ui.RenderHotFiles(hotFiles))
	}

	dates, _ := git.GetCommitDates()
	if len(dates) > 0 {
		fmt.Println(ui.RenderHeatmap(dates))
	}
}
