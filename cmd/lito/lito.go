package main

import (
	"github.com/aosasona/lito"
)

func main() {
	lito, err := lito.New(lito.Config{
		ConfigDir: "example",
	})
	if err != nil {
		panic("failed to create instance: " + err.Error())
	}

	if err := lito.Run(); err != nil {
		panic(err)
	}
}
