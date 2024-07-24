package cmd

import (
	"cpx-cli/cpx"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Print all running services",
	Run: func(cmd *cobra.Command, args []string) {
		services, err := cpx.GetAverageUsageStatistics()
		if err != nil {
			log.Fatalf("Failed to get servers: %v", err)
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintln(writer, "Service\tStatus\tCPU\tMemory")
		for _, service := range services {
			if (healthy && service.Status == "Healthy") || (unhealthy && service.Status == "Unhealthy") || (!healthy && !unhealthy) {
				fmt.Fprintf(writer, "%s\t%s\t%d%%\t%d%%\n", service.Name, service.Status, service.CPU, service.Memory)
			}
		}
		writer.Flush()
	},
}
