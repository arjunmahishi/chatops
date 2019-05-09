package messenger

import (
	"net/http"
	"testing"
)

func TestHangoutsMessenger_setup(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "test 1",
			fields:  fields{Client: nil},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HangoutsMessenger{
				Client: tt.fields.Client,
			}
			if err := hm.setup(); (err != nil) != tt.wantErr {
				t.Errorf("HangoutsMessenger.setup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
