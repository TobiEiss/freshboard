package handler_test

import (
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/TobiEiss/freeboardBackend/core"
	"github.com/TobiEiss/freeboardBackend/handler"
	"github.com/labstack/echo"
)

func TestGoogleCalendar(t *testing.T) {
	router := echo.New()
	req := httptest.NewRequest("GET", "http://test.abc", nil)
	rec := httptest.NewRecorder()
	calendarConf := core.CalendarConfig{
		GoogleCalendars:       []string{"http://www.google.com/calendar/ical/bg.bulgarian%23holiday%40group.v.calendar.google.com/public/basic.ics"},
		CountOfUpcomingEvents: 5,
	}
	context := core.CContext{
		router.NewContext(req, rec),
		core.Config{[]interface{}{calendarConf}},
	}

	// fire up upcoming events
	googleCalendar := handler.GoogleCalendar{calendarConf}
	err := googleCalendar.UpcomingEvents(&context)

	if err != nil {
		t.Fail()
	}

	body, _ := ioutil.ReadAll(rec.Result().Body)
	log.Println(string(body))
}
