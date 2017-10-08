package plugins

import (
	"net/http"

	ics "github.com/TobiEiss/ics-golang"
	"github.com/labstack/echo"
)

// ICSCalendar is the trsuct for this module
type ICSCalendar struct {
	CountOfUpcomingEvents int    `json:"countOfUpcomingEvents"`
	Calendar              string `json:"calendar"`
}

// UpcomingEvents returns upcoming events
func UpcomingEvents(context echo.Context) error {
	icsCalendar := new(ICSCalendar)
	if err := context.Bind(icsCalendar); err != nil {
		return err
	}

	// parse calendar
	ics.RepeatRuleApply = true
	parser := ics.New()
	inputChan := parser.GetInputChan()
	inputChan <- icsCalendar.Calendar
	parser.Wait()

	// get all calendars in this parser
	cal, err := parser.GetCalendars()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, cal[0].GetUpcomingEvents(icsCalendar.CountOfUpcomingEvents))
}
