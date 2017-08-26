package core

type Config struct {
	Endpoints []interface{}
}

type Endpoint struct {
	Path string
}

type CalendarConfig struct {
	Endpoint
	GoogleCalendars       []string
	CountOfUpcomingEvents int
}
