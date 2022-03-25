// Package dingo provides checks targeting dingo-specific problems
package dingo

import (
	"flamingo.me/dingo"
	"github.com/doertydoerk/flamalyzer/src/analyzers"
	"github.com/doertydoerk/flamalyzer/src/analyzers/dingo/checks/bind"
	"github.com/doertydoerk/flamalyzer/src/analyzers/dingo/checks/inject"
	"github.com/doertydoerk/flamalyzer/src/flamalyzer/configuration"
	"golang.org/x/tools/go/analysis"
)

// Module to register the dingo checks
type Module struct{}

// The default properties which are used if there is no config-file
var defaultProps = Props{
	CheckPointerReceiver:                   true,
	CheckStrictTagsAndFunctions:            true,
	CheckCorrectInterfaceToInstanceBinding: true,
}

// Props of an analyzer which will be used by the config-module to match the entries
// of a file to these variables. Is used to activate and deactivate checks, for example.
type Props struct {
	Name                                   string
	CheckPointerReceiver                   bool
	CheckStrictTagsAndFunctions            bool
	CheckCorrectInterfaceToInstanceBinding bool
}

// The Analyzer holds a set of checks, uses the config and has props that can be defined to get read by the config
type Analyzer struct {
	checks []*analysis.Analyzer
	config configuration.AnalyzerConfig
	props  Props
}

// Configure DI
func (m *Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(analyzers.Analyzer)).To(new(Analyzer))
}

// Inject dependencies
func (d *Analyzer) Inject(config configuration.AnalyzerConfig) {
	d.props = defaultProps
	d.props.Name = "dingoAnalyzer"
	d.config = config
}

// ChecksToExecute decides which checks to run
func (d *Analyzer) ChecksToExecute() []*analysis.Analyzer {
	analyzers.DecodeAnalyzerConfigurationsToAnalyzerProps(d.props.Name, d.config, &d.props)

	if d.props.CheckPointerReceiver {
		d.checks = append(d.checks, inject.ReceiverAnalyzer)
	}
	if d.props.CheckStrictTagsAndFunctions {
		d.checks = append(d.checks, inject.TagAnalyzer)
	}
	if d.props.CheckCorrectInterfaceToInstanceBinding {
		d.checks = append(d.checks, bind.Analyzer)
	}
	return d.checks
}
