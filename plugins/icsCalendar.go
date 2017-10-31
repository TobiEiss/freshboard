package plugins

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

	var customEvents []CustomEvent
	for _, event := range cal[0].GetUpcomingEvents(countInt) {
		customEvents = append(customEvents, customEventAdapter(event))
	}
	return context.JSON(http.StatusOK, customEvents)
}

func customEventAdapter(event ics.Event) CustomEvent {
	return CustomEvent{
		Event:    event,
		OnlyDate: OnlyDate{event.Start},
	}
}

// CustomEvent is the custom event to show specific dates
type CustomEvent struct {
	ics.Event
	OnlyDate OnlyDate `json:"onlyDate"`
}

// OnlyDate is the type to show only the date
type OnlyDate struct {
	time.Time
}

// MarshalJSON specific marshal
func (ct *OnlyDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format("Monday, 02-Jan-06"))), nil
}
