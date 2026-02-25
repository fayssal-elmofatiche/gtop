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
		codeStats        git.CodeStats
		contribStats     git.ContributorStats
		lastActivity     string
		velocity         git.Velocity
		depManager       string
		depCount         int
		health           git.BranchHealth
		hotFiles         []git.HotFile
		dates            []string
		license          string
		latestTag        string
		cicd             []string
		releases         []git.Release
		stashCount       int
		commitConvention string
	)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		codeStats = git.GetCodeStats()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		contribStats = git.GetContributors(5)
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

	wg.Add(1)
	go func() {
		defer wg.Done()
		license = git.GetLicense()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		latestTag = git.GetLatestTag()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cicd = git.GetCICD()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		releases = git.GetRecentReleases(5)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		stashCount = git.GetStashCount()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		commitConvention = git.GetCommitConvention()
	}()

	wg.Wait()

	primaryLang := ""
	if len(codeStats.Languages) > 0 {
		primaryLang = codeStats.Languages[0].Name
	}
	logo := ui.RenderLogo(primaryLang)
	info := ui.RenderInfo(ui.RenderParams{
		Info:             gitInfo,
		Size:             codeStats.Size,
		FileCount:        codeStats.FileCount,
		Languages:        codeStats.Languages,
		LOC:              codeStats.LOC,
		LastActivity:     lastActivity,
		Velocity:         velocity,
		DepManager:       depManager,
		DepCount:         depCount,
		Health:           health,
		License:          license,
		LatestTag:        latestTag,
		CICD:             cicd,
		StashCount:       stashCount,
		Contributors:     contribStats.Total,
		TestRatio:        codeStats.TestRatio,
		CommitConvention: commitConvention,
	})
	fmt.Println(ui.RenderLayout(logo, info))

	fmt.Println(ui.RenderLanguageBar(codeStats.Languages, 50))

	if len(contribStats.Top) > 0 {
		fmt.Println(ui.RenderContributors(contribStats))
	}

	if len(hotFiles) > 0 {
		fmt.Println(ui.RenderHotFiles(hotFiles))
	}

	if len(releases) > 0 {
		fmt.Println(ui.RenderReleases(releases))
	}

	if len(dates) > 0 {
		fmt.Println(ui.RenderHeatmap(dates))
	}
}
