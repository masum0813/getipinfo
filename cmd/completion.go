/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate the autocompletion script for the specified shell",
	Long:  `Generate the autocompletion script for the specified shell`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("completion called")
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
