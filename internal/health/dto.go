package health

type HealthResponse struct {
	Status       string     `json:"status"`
	Dependencies HealthData `json:"dependencies"`
}

type HealthData struct {
	Postgres bool `json:"postgres"`
	Cache    bool `json:"cache"`
}
