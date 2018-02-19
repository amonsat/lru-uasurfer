# lru-uasurfer
Golang LRU wrapping for [ua-surfer](https://github.com/avct/uasurfer).
## Use

### Basic Use

```go
package main

import (
	"fmt"

	surfer "github.com/Amonsat/lru-uasurfer"
)

func main() {
	s := surfer.New()
	s.LoadDump("")

	// This is an example how to use it
	r := s.Parse(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Safari/537.36 Edge/13.10586`)

	// Print results
	fmt.Printf("Result: %#v \n", r)

	// if you want to dump cache, do it:
	s.SaveDump("")
}
```