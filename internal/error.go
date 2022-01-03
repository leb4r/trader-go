package internal

import (
	"fmt"
	"os"
)

func ThrowError(err error) {
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
