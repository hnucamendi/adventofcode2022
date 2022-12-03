package logic

import (
	"os"
	"reflect"
	"testing"
)

// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
// A for Rock, B for Paper, and C for Scissors.
// X for Rock, Y for Paper, and Z for Scissors.

// shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

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
