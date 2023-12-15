package main

import (
	gots "github.com/aosasona/gots/v2"
	"github.com/aosasona/gots/v2/config"
	"go.trulyao.dev/lito/pkg/types"
)

// TODO: define API routes in a separate file as "<path>": "<request-payload-type>"
func main() {
	g := gots.Init(config.Config{
		Enabled:           gots.Bool(true),
		OutputFile:        gots.String("ext/typescript/types.ts"),
		UseTypeForObjects: gots.Bool(true),
		ExpandObjectTypes: gots.Bool(false),
	})

	g.AddSource(types.Proxy{})
	g.AddSource(types.Admin{})
	g.AddSource(types.DomainStatusDNS{})
	g.AddSource(types.DomainStatus{})
	g.AddSource(types.Domain{})
	g.AddSource(types.Service{})

	g.AddSource(types.Config{})

	if err := g.Execute(); err != nil {
		panic(err)
	}
}
