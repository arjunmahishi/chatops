package store

import (
	"reflect"
	"testing"
)

func TestCache_Write(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		c       Cache
		args    args
		wantErr bool
	}{
		{
			name:    "Test 1",
			c:       Cache{},
			args:    args{key: "key1", value: "val1"},
			wantErr: false,
		},
		{
			name:    "Test 2",
			c:       Cache{},
			args:    args{key: "key1", value: "val1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cache{}
			if err := c.Write(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Cache.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCache_Read(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		c       Cache
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "Test 1",
			c:       Cache{},
			args:    args{key: "key1"},
			want:    "val1",
			wantErr: false,
		},
		{
			name:    "Test 2",
			c:       Cache{},
			args:    args{key: "wrongKey"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cache{}
			got, err := c.Read(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cache.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
