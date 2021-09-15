package responses

type HealthCheck struct {
	Status string `json:"status"`
}

func Health() *HealthCheck {
	return &HealthCheck{Status: "OK"}
}
