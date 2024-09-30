package main

import (
	"go.trulyao.dev/lito/pkg/types"
	"go.trulyao.dev/mirror/v2"
	"go.trulyao.dev/mirror/v2/config"
	"go.trulyao.dev/mirror/v2/generator/typescript"
)

// TODO: define API routes in a separate file as "<path>": "<request-payload-type>"
func main() {
	m := mirror.New(config.Config{
		Enabled:              true,
		EnableParserCache:    true,
		FlattenEmbeddedTypes: true,
	})

	m.AddSources(
		types.Proxy{},
		types.Admin{},
		types.DomainStatusDNS{},
		types.DomainStatus{},
		types.Domain{},
		types.Service{},
		types.Config{},
	)

	target := typescript.DefaultConfig().
		SetFileName("types.ts").
		SetOutputPath("generated/typescript")

	m.AddTarget(target)

	if err := m.GenerateAndSaveAll(); err != nil {
		panic(err)
	}
}
