package core_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/TobiEiss/freeboardBackend/core"
)

func TestMarshal(t *testing.T) {
	calendarConf := core.CalendarConfig{
		GoogleCalendars:       []string{"http://www.google.com/calendar/ical/bg.bulgarian%23holiday%40group.v.calendar.google.com/public/basic.ics"},
		CountOfUpcomingEvents: 5,
	}
	config := core.Config{[]interface{}{calendarConf}}

	jsonCfg, err := json.Marshal(config)
	if err != nil {
		t.Fail()
	}
	log.Println(jsonCfg)
}
