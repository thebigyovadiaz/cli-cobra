package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "op-math",
	Short: "op-math is a CLI tool for interacting with mathematical operations",
	Long:  "op-math is a CLI tool for interacting with mathematical operations - addition, multiplication, division, and substraction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing op-math '%s'\n", err)
		os.Exit(1)
	}
}
