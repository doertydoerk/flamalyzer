package main

import (
	"flamingo.me/dingo"
	"github.com/doertydoerk/flamalyzer/src/analyzers/architecture"
	dingoAnalyzer "github.com/doertydoerk/flamalyzer/src/analyzers/dingo"
	"github.com/doertydoerk/flamalyzer/src/flamalyzer"
)

func main() {
	flamalyzer.Run(
		[]dingo.Module{
			new(dingoAnalyzer.Module),
			new(architecture.Module),
		},
	)
}
