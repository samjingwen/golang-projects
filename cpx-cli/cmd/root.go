package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cpx",
	Short: "CPX CLI is a tool to manage and monitor cloud services",
}

var (
	healthy   bool
	unhealthy bool
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	instancesCmd.Flags().BoolVar(&healthy, "healthy", false, "Filter to show only healthy services")
	instancesCmd.Flags().BoolVar(&unhealthy, "unhealthy", false, "Filter to show only unhealthy services")

	servicesCmd.Flags().BoolVar(&healthy, "healthy", false, "Filter to show only healthy services")
	servicesCmd.Flags().BoolVar(&unhealthy, "unhealthy", false, "Filter to show only unhealthy services")

	rootCmd.AddCommand(instancesCmd)
	rootCmd.AddCommand(servicesCmd)
	rootCmd.AddCommand(trackCmd)
}
