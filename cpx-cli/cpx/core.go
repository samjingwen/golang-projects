package cpx

import (
	"fmt"
	"sync"
)

type InstanceInfo struct {
	IP     string
	Name   string
	Status string
	CPU    int
	Memory int
}

type ServiceInfo struct {
	Name   string
	Status string
	CPU    int
	Memory int
}

func GetServerUsageStatistics() ([]InstanceInfo, error) {
	ips, err := getAllServers()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	instanceInfoChan := make(chan InstanceInfo, len(ips))
	errorChan := make(chan error, len(ips))

	for _, ip := range ips {
		wg.Add(1)

		go func(ip string) {
			defer wg.Done()

			stats, err := getServerUsageStatistics(ip)
			if err != nil {
				errorChan <- err
				return
			}

			cpu, mem := parseUsage(stats.CPU), parseUsage(stats.Memory)

			var status string
			if cpu > CPU_THRESHOLD || mem > MEM_THRESHOLD {
				status = "Unhealthy"
			} else {
				status = "Healthy"
			}

			info := InstanceInfo{
				IP:     ip,
				Name:   stats.Name,
				Status: status,
				CPU:    cpu,
				Memory: mem,
			}

			instanceInfoChan <- info
		}(ip)
	}

	wg.Wait()
	close(instanceInfoChan)
	close(errorChan)

	if len(errorChan) > 0 {
		return nil, <-errorChan
	}

	var infos []InstanceInfo
	for info := range instanceInfoChan {
		infos = append(infos, info)
	}

	return infos, nil
}

func GetAverageUsageStatistics() (map[string]ServiceInfo, error) {
	servers, err := GetServerUsageStatistics()
	if err != nil {
		return nil, err
	}

	stats := make(map[string]ServiceInfo)
	counts := make(map[string]int)
	healthyCounts := make(map[string]int)

	for _, server := range servers {
		if stat, ok := stats[server.Name]; ok {
			stats[server.Name] = ServiceInfo{
				CPU:    stat.CPU + server.CPU,
				Memory: stat.Memory + server.Memory,
			}
		} else {
			stats[server.Name] = ServiceInfo{
				CPU:    server.CPU,
				Memory: server.Memory,
			}
		}

		counts[server.Name]++
		if server.Status == "Healthy" {
			healthyCounts[server.Name]++
		}
	}

	for name, stat := range stats {
		var status string
		if healthyCounts[name] >= COUNT_THRESHOLD {
			status = "Healthy"
		} else {
			status = "Unhealthy"
		}

		stats[name] = ServiceInfo{
			Name:   name,
			Status: status,
			CPU:    stat.CPU / counts[name],
			Memory: stat.Memory / counts[name],
		}
	}

	return stats, nil
}

func parseUsage(usage string) int {
	var val int
	_, _ = fmt.Sscanf(usage, "%d%%", &val)
	return val
}
