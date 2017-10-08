package plugins

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	ics "github.com/TobiEiss/ics-golang"
	"github.com/labstack/echo"
)

const (
	CountOfUpcomingEvents = "count"
	Calendar              = "calendar"
)

// UpcomingEvents returns upcoming events
func UpcomingEvents(context echo.Context) error {
	count := context.Request().Header.Get(CountOfUpcomingEvents)
	calendar := context.Request().Header.Get(Calendar)
	if count == "" || calendar == "" {
		return errors.New("set header count and calendar")
	}

	// set count to int
	countInt, err := strconv.Atoi(count)
	if err != nil {
		return err
	}

	// parse calendar
	ics.RepeatRuleApply = true
	parser := ics.New()
	inputChan := parser.GetInputChan()
	inputChan <- calendar
	parser.Wait()

	// get all calendars in this parser
	cal, err := parser.GetCalendars()
	if err != nil {
		log.Println(err)
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, cal[0].GetUpcomingEvents(countInt))
}
