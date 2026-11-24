package main

import (
	"go/build"
	"os"
	"path/filepath"
	"testing"
)

func TestMatchFileHonorsGOOSFilename(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "only_windows.go"), []byte("package sample\n"), 0600); err != nil { t.Fatal(err) }
	ctx := build.Default
	ctx.GOOS, ctx.GOARCH = "windows", "amd64"
	match, err := ctx.MatchFile(dir, "only_windows.go")
	if err != nil || !match { t.Fatalf("expected windows match: match=%v err=%v", match, err) }
	ctx.GOOS = "linux"
	match, err = ctx.MatchFile(dir, "only_windows.go")
	if err != nil || match { t.Fatalf("expected windows file excluded: match=%v err=%v", match, err) }
}
