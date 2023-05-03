## Kompas 3D files Reader

Package to read spc data from the Kompas files.

### Usage

Add `kompassreader` to project:

```shell
go get github.com/kukymbr/kompasreader
```

Import package `github.com/kukymbr/kompasreader` and 
read Kompas file:

```go
package main

import (
	"fmt"

	"github.com/kukymbr/kompasreader"
)

func main() {
	filepath := "testdata/example.spw"

	reader, err := kompasreader.New(filepath)
	if err != nil {
		panic(err)
	}

	doc, err := reader.Read()
	if err != nil {
		panic(err)
	}
	
	fmt.Println(doc.FileInfo.Author)
}
```