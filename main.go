package main

import (
	"github.com/leb4r/trader-go/cmd"
	"github.com/leb4r/trader-go/internal"
)

func main() {
	if err := cmd.Execute(); err != nil {
		internal.ThrowError(err)
	}
}
