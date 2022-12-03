package logic

import (
	"os"
	"reflect"
	"testing"
)

func Test_Play(t *testing.T) {
	f, err := os.ReadFile("../input.txt")
	if err != nil {
		return
	}

	tests := []struct {
		name    string
		wantErr bool
		in      []byte
		comp    int
		usr     int
	}{
		{
			name:    "Test Play",
			wantErr: false,
			in: []byte(`
				B X 
				B Z
				B Z
				A Y
				`),

			comp: 13,
			usr:  27,
		},
		{
			name:    "Full Test",
			wantErr: false,
			in:      f,
			comp:    8634,
			usr:     15337,
		},
	}

	for _, tt := range tests {
		comp, usr, err := Play(tt.in)
		if (err != nil) != tt.wantErr {
			t.Fatalf("got error %v want err %v", err, tt.wantErr)
		}

		if !reflect.DeepEqual(comp, tt.comp) {
			t.Fatalf("got comp %v want %v", comp, tt.comp)
		}

		if !reflect.DeepEqual(usr, tt.usr) {
			t.Fatalf("got usr %v want %v", usr, tt.usr)
		}
	}
}

func Test_PlayR(t *testing.T) {
	f, err := os.ReadFile("../input.txt")
	if err != nil {
		return
	}

	tests := []struct {
		name    string
		wantErr bool
		in      []byte
		comp    int
		usr     int
	}{
		{
			name:    "Test Play",
			wantErr: false,
			in: []byte(`
			B X 
			B Z
			B Z
			A Y
			`),

			comp: 16,
			usr:  23,
		},
		{
			name:    "Full Test",
			wantErr: false,
			in:      f,
			comp:    11439,
			usr:     11696,
		},
	}

	for _, tt := range tests {
		comp, usr, err := PlayR(tt.in)
		if (err != nil) != tt.wantErr {
			t.Fatalf("got error %v want err %v", err, tt.wantErr)
		}

		if !reflect.DeepEqual(comp, tt.comp) {
			t.Fatalf("got comp %v want %v", comp, tt.comp)
		}

		if !reflect.DeepEqual(usr, tt.usr) {
			t.Fatalf("got usr %v want %v", usr, tt.usr)
		}
	}
}
