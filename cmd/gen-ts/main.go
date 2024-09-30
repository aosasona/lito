package main

import (
	"go.trulyao.dev/lito/pkg/types"
	"go.trulyao.dev/mirror/v2"
	"go.trulyao.dev/mirror/v2/config"
	"go.trulyao.dev/mirror/v2/generator/typescript"
	"go.trulyao.dev/mirror/v2/parser"
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

	m.Parser().
		OnParseItem(func(sourceName string, target parser.Item) error {
			if target, ok := target.(*parser.Struct); ok {
				// loop through all fields in the struct to make them not nullable since they are already optional
				for _, field := range target.Fields {
					if field.BaseItem.IsNullable() {
						switch baseItem := field.BaseItem.(type) {
						case *parser.Scalar:
							baseItem.Nullable = false
						case *parser.Map:
							baseItem.Nullable = false
						case *parser.Struct:
							baseItem.Nullable = false
						case *parser.List:
							baseItem.Nullable = false
						}
					}
				}
			}

			return nil
		})

	if err := m.GenerateAndSaveAll(); err != nil {
		panic(err)
	}
}
