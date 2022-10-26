package main

import (
	"context"

	"github.com/6missedcalls/lw3-tool/cmd"
)

var (
	Version = "0.0.0"
	Commit  = "NONE"
	Date    = "unknown"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "version", Version)
	ctx = context.WithValue(ctx, "commit", Commit)
	ctx = context.WithValue(ctx, "date", Date)
	cmd.Execute(ctx)
}
