## Kompas 3D files Reader

Package to read specifications data from the Kompas files.

### Usage

Add library to project:

```shell
go get github.com/kukymbr/kompas
```

Import packages and read Kompas file:

```go
package main

import (
	"fmt"

	"github.com/kukymbr/kompas"
	"github.com/kukymbr/kompas/kompasconv"
)

func main() {
	var doc *kompas.Document
	var err error
	
	filepath := "testdata/example.spw"

	reader, err := kompasconv.NewReader(filepath)
	if err != nil {
		panic(err)
	}

	doc, err = reader.Read()
	if err != nil {
		panic(err)
	}
	
	fmt.Println(doc.FileInfo.Author)
}
```