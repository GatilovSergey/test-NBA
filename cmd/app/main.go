package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use: filepath.Base(os.Args[0]),
}

func main (){
	if err := Root.Execute(); err != nil && err != context.Canceled {
		os.Exit(1)
	}
}
