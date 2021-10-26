package main

import "time"

func main() {
	apiClient := Default("", time.Now())
}

type ApiClient interface {
}

type apiClient struct {
	endpoint string
	timeout  time.Time
}

func Default(endpoint string, timeout time.Time) ApiClient {
	return apiClient{
		endpoint,
		timeout,
	}
}
