# Lito

# Features

- [ ] Web dashboard
- [ ] File-based config and services definition
- [ ] REST API to modify config
- [ ] Watch and reload proxy and services config
- [ ] Add pre- and post- request hooks
- [ ] Add pre- and post- response hooks

# Running

If you don't want to run the binary and you want to run the proxy in your own file or add some extra config, you can do it like this:

```go
package main

import (
	"github.com/aosasona/lito"
)

func main() {
	lito, err := lito.New(lito.Config{
		DataDir:   "example/data",
		ConfigDir: "example",
	})
	if err != nil {
		panic("failed to create instance: " + err.Error())
	}

	if err := lito.Run(); err != nil {
		panic(err)
	}
}
```
