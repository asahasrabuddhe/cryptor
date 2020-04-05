package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "cryptor",
}

func Execute() {
	rootCmd.Execute()
}
