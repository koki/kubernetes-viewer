package main

import (
	"os"

	"github.com/koki/kubernetes-viewer/server/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
