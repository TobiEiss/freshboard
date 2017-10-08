package plugins_test

import (
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/TobiEiss/freeboardBackend"
	"github.com/labstack/echo"
)

func TestGoogleCalendar(t *testing.T) {
	router := echo.New()
	req := httptest.NewRequest("GET", "http://test.abc", nil)
	rec := httptest.NewRecorder()
	calendar := freeboardBackend.InitGoogleCalendar(
		5,
		[]string{"http://www.google.com/calendar/ical/bg.bulgarian%23holiday%40group.v.calendar.google.com/public/basic.ics"})
	context := router.NewContext(req, rec)

	// fire up upcoming events
	err := calendar.UpcomingEvents(context)

	if err != nil {
		t.Fail()
	}

	body, _ := ioutil.ReadAll(rec.Result().Body)
	log.Println(string(body))
}
