package cpx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		ips := []string{"10.58.1.1", "10.58.1.2", "10.58.1.3"}
		json.NewEncoder(w).Encode(ips)
	})
	handler.HandleFunc("/10.58.1.1", func(w http.ResponseWriter, r *http.Request) {
		service := Server{Name: "PermissionsService", CPU: "20%", Memory: "30%"}
		json.NewEncoder(w).Encode(service)
	})
	handler.HandleFunc("/10.58.1.2", func(w http.ResponseWriter, r *http.Request) {
		service := Server{Name: "PermissionsService", CPU: "40%", Memory: "50%"}
		json.NewEncoder(w).Encode(service)
	})
	handler.HandleFunc("/10.58.1.3", func(w http.ResponseWriter, r *http.Request) {
		service := Server{Name: "AuthService", CPU: "90%", Memory: "95%"}
		json.NewEncoder(w).Encode(service)
	})

	return httptest.NewServer(handler)
}

func TestGetServerUsageStatistics(t *testing.T) {
	server := mockServer()
	defer server.Close()

	// Override the base URL to use the mock server
	oldBaseURL := BASE_URL
	BASE_URL = server.URL
	defer func() { BASE_URL = oldBaseURL }()

	instances, err := GetServerUsageStatistics()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(instances) != 3 {
		t.Fatalf("Expected 3 instances, got %d", len(instances))
	}

	expected := []InstanceInfo{
		{IP: "10.58.1.1", Name: "PermissionsService", Status: "Healthy", CPU: 20, Memory: 30},
		{IP: "10.58.1.2", Name: "PermissionsService", Status: "Healthy", CPU: 40, Memory: 50},
		{IP: "10.58.1.3", Name: "AuthService", Status: "Unhealthy", CPU: 90, Memory: 95},
	}

	for i, instance := range instances {
		if instance != expected[i] {
			t.Errorf("Expected service %v, got %v", expected[i], instance)
		}
	}
}

func TestGetAverageUsageStatistics(t *testing.T) {
	server := mockServer()
	defer server.Close()

	// Override the base URL to use the mock server
	oldBaseURL := BASE_URL
	BASE_URL = server.URL
	defer func() { BASE_URL = oldBaseURL }()

	stats, err := GetAverageUsageStatistics()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := map[string]ServiceInfo{
		"PermissionsService": {Name: "PermissionsService", Status: "Healthy", CPU: 30.0, Memory: 40.0},
		"AuthService":        {Name: "AuthService", Status: "Unhealthy", CPU: 90.0, Memory: 95.0},
	}

	for name, stat := range stats {
		expectedStat, ok := expected[name]
		if !ok {
			t.Errorf("Unexpected service %s", name)
		}

		if stat != expectedStat {
			t.Errorf("Expected stats for %s: %+v, got %+v", name, expectedStat, stat)
		}
	}
}
