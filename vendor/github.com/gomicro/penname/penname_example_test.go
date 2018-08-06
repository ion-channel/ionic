package penname_test

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gomicro/penname"
)

func ExamplePenName() {
	mockWrite := penname.New()
	mw := io.MultiWriter(os.Stdout, mockWrite)

	mw.Write([]byte("A random line to write, "))

	if strings.Contains(string(mockWrite.Written()), "random") {
		fmt.Println("Found a random")
	}
	// Output: A random line to write, Found a random
}
