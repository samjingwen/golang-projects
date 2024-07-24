package cmd

import (
	"cpx-cli/cpx"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

var instancesCmd = &cobra.Command{
	Use:   "instances",
	Short: "Print all running instances",
	Run: func(cmd *cobra.Command, args []string) {
		servers, err := cpx.GetServerUsageStatistics()
		if err != nil {
			log.Fatalf("Failed to get servers: %v", err)
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintln(writer, "IP\tService\tStatus\tCPU\tMemory")
		for _, server := range servers {
			if (healthy && server.Status == "Healthy") || (unhealthy && server.Status == "Unhealthy") || (!healthy && !unhealthy) {
				fmt.Fprintf(writer, "%s\t%s\t%s\t%d%%\t%d%%\n", server.IP, server.Name, server.Status, server.CPU, server.Memory)
			}
		}
		writer.Flush()
	},
}
