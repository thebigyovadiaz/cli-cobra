package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var shouldRoundUp bool
var multiplyCmd = &cobra.Command{
	Use:     "multiply",
	Aliases: []string{"mul", "multi", "multiple"},
	Short:   "Multiply a number from another",
	Long:    "Carry out multiplication operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Subtraction of %s from %s = %s.\n\n", args[1], args[0], Multiply(args[0], args[1], shouldRoundUp))
	},
}

func init() {
	multiplyCmd.Flags().BoolVarP(&shouldRoundUp, "round", "r", false, "Round results up to 2 decimal places")
	rootCmd.AddCommand(multiplyCmd)
}
