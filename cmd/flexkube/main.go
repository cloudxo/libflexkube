package main

import (
	"os"

	"github.com/flexkube/libflexkube/cli/flexkube"
)

func main() {
	os.Exit(flexkube.Run(os.Args))
}
