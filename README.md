# go-aguid

A **G**lobally **U**nique **ID**entifier (GUID) generator in Go.

[![GoDoc](https://godoc.org/github.com/stevelacy/go-aguid?status.svg)](https://godoc.org/github.com/stevelacy/go-aguid)
![](https://img.shields.io/badge/license-MIT-blue.svg)


### Example

```go
import (
	"fmt"
	"github.com/stevelacy/go-aguid"
)

// Return a deterministic reproducible UUID
func Example() {
	uuid, err := aguid.Aguid("input-string")
	if err != nil {
		panic(err)
	}
	fmt.Printf("uuid: %s \n", uuid)
}

// Return a unique random UUID.V4()
func Example_random() {
	uuid, err := aguid.Aguid("")
	if err != nil {
		panic(err)
	}
	fmt.Printf("uuid: %s \n", uuid)
}
```

### License [MIT](./LICENSE)
