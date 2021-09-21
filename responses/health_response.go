package responses

import (
	"os"

	"gizumon.com/life-manager-api/constants"
)

type HealthCheck struct {
	Status     string `json:"status"`
	ApiVersion string `json:"api_version"`
}

func Health() *HealthCheck {
	return &HealthCheck{Status: "OK", ApiVersion: os.Getenv(constants.API_PREFIX_KEY)}
}
