package logic

import (
	"reflect"
	"testing"
)

func Test_Split(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		bag1    []string
		bag2    []string
		wantErr bool
	}{
		{
			name: "Small Test",
			in: []byte(
				`vJrwpWtwJgWrhcsFMMfFFhFp
			jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
			PmmdzqPrVvPwwTWBwg`),
			bag1:    []string{"vJrwpWtwJgWr", "jqHRNqRjqzjGDLGL", "PmmdzqPrV"},
			bag2:    []string{"hcsFMMfFFhFp", "rsFMfFZSrLrFZsSL", "vPwwTWBwg"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		bag1, bag2, err := Split(tt.in)
		if (err != nil) != tt.wantErr {
			t.Fatalf("error %v, wanted error %v", err, tt.wantErr)
		}

		if !reflect.DeepEqual(bag1, tt.bag1) {
			t.Errorf("got %v, wanted %v", bag1, tt.bag1)
		}

		if !reflect.DeepEqual(bag2, tt.bag2) {
			t.Fatalf("got %v, wanted %v", bag1, tt.bag1)
		}
	}
}

func Test_CommonType(t *testing.T) {
	tests := []struct {
		name    string
		bag1    []string
		bag2    []string
		want    string
		wantErr bool
	}{
		{
			name:    "Small Test",
			bag1:    []string{"vJrwpWtwJgWr", "jqHRNqRjqzjGDLGL", "PmmdzqPrV"},
			bag2:    []string{"hcsFMMfFFhFp", "rsFMfFZSrLrFZsSL", "vPwwTWBwg"},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		got, err := CommonType(tt.bag1, tt.bag2)
		if (err != nil) != tt.wantErr {
			t.Fatalf("error %v, wanted error %v", err, tt.wantErr)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, wanted %v", got, tt.want)
		}
	}
}
