package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aosasona/lito/cmd/lito"
	"github.com/aosasona/lito/pkg/config"
	"github.com/spf13/viper"
)

func main() {
	configPath := flag.String("config", ".", "Path to lito's config file")

	flag.Parse()

	err := config.Load(*configPath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	proxy := lito.Proxy{
		Host: viper.GetString("proxy.host"),
		Port: viper.GetInt("proxy.port"),
	}

	fmt.Print(proxy)
}
