// +build tools

package tools // import "github.com/AlekSi/applehealth/tools"

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/quasilyte/go-consistent"
	_ "golang.org/x/perf/cmd/benchstat"
	_ "mvdan.cc/gofumpt/gofumports"
)
