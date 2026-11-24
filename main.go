package main

import (
	"flag"
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"
)

func inspect(dir, goos, goarch, tags string) error {
	ctx := build.Default
	ctx.GOOS = goos
	ctx.GOARCH = goarch
	ctx.BuildTags = strings.Fields(strings.ReplaceAll(tags, ",", " "))
	entries, err := os.ReadDir(dir)
	if err != nil { return err }
	fmt.Printf("Target: %s/%s  tags=%q\n", ctx.GOOS, ctx.GOARCH, ctx.BuildTags)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") { continue }
		match, err := ctx.MatchFile(dir, entry.Name())
		if err != nil {
			fmt.Printf("[error] %-28s %v\n", entry.Name(), err)
			continue
		}
		state := "excluded"
		if match { state = "selected" }
		fmt.Printf("[%-8s] %s\n", state, entry.Name())
	}
	return nil
}

func main() {
	dir := flag.String("dir", ".", "directory containing Go source files")
	goos := flag.String("goos", build.Default.GOOS, "target operating system")
	goarch := flag.String("goarch", build.Default.GOARCH, "target architecture")
	tags := flag.String("tags", "", "comma- or space-separated build tags")
	flag.Parse()
	abs, err := filepath.Abs(*dir)
	if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
	if err := inspect(abs, *goos, *goarch, *tags); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
