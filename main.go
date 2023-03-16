package main

import (
	"flag"
	"sync"

	"github.com/aosasona/lito/cmd/lito"
	"github.com/aosasona/lito/internal/db"
)

func main() {

	port := flag.Int("p", 80, "Port to run Lito on")
	flag.Parse()

	_ = make(chan error)
	db := db.Init()

	lito, _ := lito.New(db, *port)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := lito.Run(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
