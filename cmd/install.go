/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/8ideaz/kion/internal/manager"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install <package-name>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var packageName string
		if len(args) > 0 {
			packageName = args[0]
			if packageName == "" {
				err := fmt.Errorf("no package name supplied")
				return err
			}
		} else {
			err := fmt.Errorf("no package name supplied")
			return err
		}
		err := manager.Install(packageName)
		if err != nil {
			fmt.Printf("Failed to install %s: %v\n", packageName, err)
			err = fmt.Errorf("failed to install %s: %w", packageName, err)
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
