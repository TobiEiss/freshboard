package core

type Config struct {
	Endpoints []interface{} `json:"endpoints"`
}

type Endpoint struct {
	Path string `json:"path"`
}

type CalendarConfig struct {
	Endpoint
	GoogleCalendars       []string `json:"googleCalendars"`
	CountOfUpcomingEvents int      `json:"countOfUpcomingEvents"`
}
