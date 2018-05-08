# Pen Name
[![Build Status](https://travis-ci.org/gomicro/penname.svg)](https://travis-ci.org/gomicro/penname)
[![GoDoc](https://godoc.org/github.com/gomicro/penname?status.png)](https://godoc.org/github.com/gomicro/penname)

A mock that implements the Closer & Writer interfaces for testing.

# Example

```go
import(
	"fmt"
	"io"
	"os"

	"github.com/gomicro/penname"
)

func main(){
	mockWrite := penname.New()
	mw := io.MultiWriter(os.Stdout, mockWrite)

	mw.Write("A random line to write")

	if strings.Contains( string(mockWrite.Written), "random" ){
		fmt.Println("Found a random")
	}
}
```
