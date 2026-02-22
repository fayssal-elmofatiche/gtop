package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"sync"

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

	var (
		languages    []git.LanguageStat
		size         string
		fileCount    int
		contributors []git.Contributor
		loc          int
		lastActivity string
		velocity     git.Velocity
		depManager   string
		depCount     int
		health       git.BranchHealth
		hotFiles     []git.HotFile
		dates        []string
	)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		languages = git.GetLanguageStats()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		size, fileCount = git.GetRepoSize()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		contributors = git.GetContributors(5)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		loc = git.GetLinesOfCode()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		lastActivity = git.GetLastActivity()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		velocity = git.GetVelocity()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		depManager, depCount = git.GetDependencyCount()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		health = git.GetBranchHealth()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		hotFiles = git.GetHotFiles(5)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		dates, _ = git.GetCommitDates()
	}()

	wg.Wait()

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

	if len(dates) > 0 {
		fmt.Println(ui.RenderHeatmap(dates))
	}
}
