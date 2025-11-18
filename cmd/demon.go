/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/demon"
	"github.com/spf13/cobra"
)

// demonCmd represents the demon command
var demonCmd = &cobra.Command{
	Use:   "demon",
	Short: "Start Pomodoro timer demon",
	Long:  `Start Pomodoro timer demon that communicate with tomato cli client`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return demon.Serve()
	},
}

func init() {
	rootCmd.AddCommand(demonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// demonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// demonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
