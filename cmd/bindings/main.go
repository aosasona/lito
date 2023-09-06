package main

import (
	"log"

	"github.com/aosasona/gots/v2"
	"go.trulyao.dev/lito"
)

func main() {
	g := gots.Init(gots.Config{
		Enabled:           gots.Bool(true),
		PreferUnknown:     gots.Bool(true),
		UseTypeForObjects: gots.Bool(true),
		OutputFile:        gots.String("interfaces/lito.ts/lito.ts"),
	})

	g.Register(
		lito.Admin{},
		lito.Service{},
		lito.DomainStatusCert{},
		lito.DomainStatusDNS{},
		lito.DomainStatus{},
		lito.Retry{},
		lito.Domain{},
		lito.Proxy{},
		lito.Config{},
	)

	if err := g.Execute(); err != nil {
		log.Fatalf("Error generating typescript types: %v", err)
	}
}
