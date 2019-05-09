package payload

import (
	"testing"
	"time"
)

func Test_validateTime(t *testing.T) {
	type args struct {
		eventTime string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Running with current time",
			args: args{eventTime: time.Now().Format(time.RFC3339)},
			want: true,
		},
		{
			name: "Running with wrong time format",
			args: args{eventTime: time.Now().String()},
			want: false,
		},
		{
			name: "Running with invalid timestamp",
			args: args{eventTime: time.Now().Add(-1 * time.Minute).Format(time.RFC3339)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateTime(tt.args.eventTime); got != tt.want {
				t.Errorf("validateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
