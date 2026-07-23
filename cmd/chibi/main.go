package main

import (
	"chibi/internal/cli"
	"chibi/internal/telemetry"
)

func main() {
	defer telemetry.Sync()
	cli.Execute()
}
