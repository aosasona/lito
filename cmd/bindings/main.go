package main

import (
	"log"

	"github.com/aosasona/gots/v2"
	"go.trulyao.dev/lito/pkg/types"
)

func main() {
	g := gots.Init(gots.Config{
		Enabled:           gots.Bool(true),
		PreferUnknown:     gots.Bool(true),
		UseTypeForObjects: gots.Bool(true),
		OutputFile:        gots.String("../lito-ts/lito.ts"),
	})

	g.Register(
		types.Admin{},
		types.Service{},
		types.DomainStatusCert{},
		types.DomainStatusDNS{},
		types.DomainStatus{},
		types.Retry{},
		types.Domain{},
		types.Proxy{},
		types.Config{},
	)

	if err := g.Execute(); err != nil {
		log.Fatalf("Error generating typescript types: %v", err)
	}
}
