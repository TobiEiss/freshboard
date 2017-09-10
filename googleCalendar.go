package freeboardBackend

import (
	"net/http"

	ics "github.com/PuloV/ics-golang"
	"github.com/labstack/echo"
)

// GoogleCalendar is the trsuct for this module
type GoogleCalendar struct {
	countOfUpcomingEvents int
	googleCalendars       []string
}

func InitGoogleCalendar(countOfUpcomingEvents int, googleCalendars []string) *GoogleCalendar {
	return &GoogleCalendar{countOfUpcomingEvents, googleCalendars}
}

// UpcomingEvents returns upcoming events
func (googleCalendar *GoogleCalendar) UpcomingEvents(context echo.Context) error {
	// parse calendar
	parser := ics.New()
	inputChan := parser.GetInputChan()
	for _, calendarAddress := range googleCalendar.googleCalendars {
		inputChan <- calendarAddress
	}

	//  wait for the calendar to be parsed
	parser.Wait()

	// get all calendars in this parser
	cal, err := parser.GetCalendars()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, cal[0].GetUpcomingEvents(googleCalendar.countOfUpcomingEvents))
}
