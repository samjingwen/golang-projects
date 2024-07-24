package cmd

import (
	"cpx-cli/cpx"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"text/tabwriter"
	"time"
)

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Print CPU/Memory of all instances of a given service over an interval",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		if _, ok := cpx.SERVICES[serviceName]; !ok {
			log.Fatalf("Invalid service: %v", serviceName)
		}

		interval, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Invalid interval: %v", err)
		}

		ticker := time.NewTicker(time.Duration(interval) * time.Second)
		done := make(chan bool)
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			printServiceStatistics(serviceName)
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					printServiceStatistics(serviceName)
				}
			}
		}()

		<-sigs
		ticker.Stop()
		done <- true
	},
}

func printServiceStatistics(serviceName string) {
	servers, err := cpx.GetServerUsageStatistics()
	if err != nil {
		log.Printf("Failed to get service stats: %v", err)
	} else {
		var instances []cpx.InstanceInfo
		for _, server := range servers {
			if server.Name == serviceName {
				instances = append(instances, server)
			}
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w, "Service: %s\n", serviceName)
		fmt.Fprintf(w, "Current Time: %s\n", time.Now().Format("2006-01-02 15:04:05.000000"))
		fmt.Fprintln(w, "IP\tStatus\tCPU\tMemory")
		for _, instance := range instances {
			fmt.Fprintf(w, "%s\t%s\t%d%%\t%d%%\n", instance.IP, instance.Status, instance.CPU, instance.Memory)
		}
		w.Flush()
	}

}
