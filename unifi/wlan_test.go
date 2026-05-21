package unifi_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ubiquiti-community/go-unifi/unifi"
)

func TestWLANMarshalJSON_ScheduleWithDurationOmitted(t *testing.T) {
	wlan := unifi.WLAN{}

	data, err := json.Marshal(&wlan)
	if err != nil {
		t.Fatal(err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatal(err)
	}

	_, exists := parsed["schedule_with_duration"]
	assert.False(t, exists, "schedule_with_duration should be omitted when empty, not serialized as null")

	raw := string(data)
	assert.NotContains(t, raw, "schedule_with_duration", "field should not appear in JSON at all")
}

func TestWLANMarshalJSON_ScheduleWithDurationPresent(t *testing.T) {
	wlan := unifi.WLAN{
		ScheduleWithDuration: []unifi.WLANScheduleWithDuration{
			{
				StartDaysOfWeek: []string{"mon"},
				StartHour:       ptrInt64(8),
				StartMinute:     ptrInt64(0),
				DurationMinutes: ptrInt64(600),
				Name:            "test-schedule",
			},
		},
	}

	data, err := json.Marshal(&wlan)
	if err != nil {
		t.Fatal(err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatal(err)
	}

	sched, exists := parsed["schedule_with_duration"]
	assert.True(t, exists, "schedule_with_duration should be present when populated")
	arr, ok := sched.([]interface{})
	assert.True(t, ok)
	assert.Len(t, arr, 1)
}

func ptrInt64(v int64) *int64 {
	return &v
}