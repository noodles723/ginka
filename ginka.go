package ginka

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
)

func NewReporter() []ginkgo.Reporter {
	stenographer := NewStenographer(config.DefaultReporterConfig.NoColor)
	return []ginkgo.Reporter{reporters.NewDefaultReporter(config.DefaultReporterConfig, stenographer)}
}
