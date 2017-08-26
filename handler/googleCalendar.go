package handler

import (
	"net/http"

	ics "github.com/PuloV/ics-golang"
	"github.com/TobiEiss/freeboardBackend/core"
	"github.com/labstack/echo"
)

// GoogleCalendar interacts with google-calendar
type GoogleCalendar struct {
	CalendarConfig core.CalendarConfig
}

type freeboardEvent struct {
	Start string
}

// UpcomingEvents returns upcoming events
func (googleCalendar *GoogleCalendar) UpcomingEvents(c echo.Context) error {
	context := c.(*core.CContext)

	// parse calendar
	parser := ics.New()
	inputChan := parser.GetInputChan()
	for _, calendarAddress := range googleCalendar.CalendarConfig.GoogleCalendars {
		inputChan <- calendarAddress
	}

	//  wait for the calendar to be parsed
	parser.Wait()

	// get all calendars in this parser
	cal, err := parser.GetCalendars()
	if err != nil {
		return context.JSONFFail(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, cal[0].GetUpcomingEvents(googleCalendar.CalendarConfig.CountOfUpcomingEvents))
}
