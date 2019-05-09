package payload

import (
	"testing"
	"time"
)

func TestHangoutsPayload_Validate(t *testing.T) {
	type fields struct {
		Type      string
		EventTime string
		Token     string
		Message   message
		User      user
		Space     space
	}

	fullInvalidfields := fields{Token: "", EventTime: ""}
	partialInvalidfields1 := fields{Token: "", EventTime: time.Now().Format(time.RFC3339)}
	partialInvalidfields2 := fields{Token: Token, EventTime: ""}
	validfields := fields{Token: Token, EventTime: time.Now().Format(time.RFC3339)}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Running with fully invalid fields",
			fields: fullInvalidfields,
			want:   false,
		},
		{
			name:   "Running with invalid token",
			fields: partialInvalidfields1,
			want:   false,
		},
		{
			name:   "Running with invalid timestamp",
			fields: partialInvalidfields2,
			want:   false,
		},
		{
			name:   "Running all valid fields",
			fields: validfields,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hp := HangoutsPayload{
				Type:      tt.fields.Type,
				EventTime: tt.fields.EventTime,
				Token:     tt.fields.Token,
				Message:   tt.fields.Message,
				User:      tt.fields.User,
				Space:     tt.fields.Space,
			}
			if got := hp.Validate(); got != tt.want {
				t.Errorf("HangoutsPayload.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHangoutsPayload_Other(t *testing.T) {
	samplePayload := HangoutsPayload{
		Message: message{Name: "", ArgumentText: "message"},
		User:    user{DisplayName: "name", Email: "email"},
		Space:   space{Name: "space"},
	}

	if samplePayload.GetMessage() != "message" {
		t.Fail()
	}
	if samplePayload.GetSenderEmail() != "email" {
		t.Fail()
	}
	if samplePayload.GetSenderName() != "name" {
		t.Fail()
	}
	if samplePayload.GetSpace() != "space" {
		t.Fail()
	}
}
