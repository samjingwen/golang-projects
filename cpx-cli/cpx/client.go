package cpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Server struct {
	Name   string `json:"service"`
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

var httpClient = &http.Client{}

func init() {
	tr := &http.Transport{
		MaxConnsPerHost: MAX_CONCURRENT_REQUESTS,
		MaxIdleConns:    100,
		IdleConnTimeout: 30 * time.Second,
	}
	httpClient = &http.Client{Transport: tr}
}

func getAllServers() ([]string, error) {
	res, err := httpClient.Get(fmt.Sprintf("%s/servers", BASE_URL))

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ipAddresses []string
	if err := json.NewDecoder(res.Body).Decode(&ipAddresses); err != nil {
		return nil, err
	}

	io.Copy(io.Discard, res.Body)
	return ipAddresses, nil
}

func getServerUsageStatistics(ipAddress string) (Server, error) {
	res, err := httpClient.Get(fmt.Sprintf("%s/%s", BASE_URL, ipAddress))
	if err != nil {
		return Server{}, err
	}
	defer res.Body.Close()

	var server Server
	if err := json.NewDecoder(res.Body).Decode(&server); err != nil {
		return Server{}, err
	}

	io.Copy(io.Discard, res.Body)
	return server, nil
}
