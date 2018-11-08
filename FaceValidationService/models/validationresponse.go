package models

// ServiceRunningResponse defines response to indicate the service is running
type ServiceRunningResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}
