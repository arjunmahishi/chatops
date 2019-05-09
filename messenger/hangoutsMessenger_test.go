package messenger

import (
	"net/http"
	"testing"
)

func TestHangoutsMessenger_Send(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		space   string
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test without error",
			fields:  fields{Client: nil},
			args:    args{space: "spaces/4OxYVAAAAAE", message: "test message"},
			wantErr: false,
		},
		// {
		// 	name:    "Test without error",
		// 	fields:  fields{Client: nil},
		// 	args:    args{space: "spaces/wrong-space", message: "test message"},
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HangoutsMessenger{
				Client: tt.fields.Client,
			}
			hm.setup()
			if err := hm.Send(tt.args.space, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("HangoutsMessenger.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHangoutsMessenger_Update(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		space       string
		messageName string
		newMessage  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test without error",
			fields:  fields{Client: nil},
			args:    args{space: "spaces/4OxYVAAAAAE", messageName: "1zAYb5rsz-w.1zAYb5rsz-w", newMessage: "test message 5"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HangoutsMessenger{
				Client: tt.fields.Client,
			}
			hm.setup()
			if err := hm.Update(tt.args.space, tt.args.messageName, tt.args.newMessage); (err != nil) != tt.wantErr {
				t.Errorf("HangoutsMessenger.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTest(t *testing.T) {
	hm := NewMessenger()
	err := hm.Send("spaces/AAAA8fjBKEQ", `<users/all> test`)
	if err != nil {
		t.Fail()
	}
}
