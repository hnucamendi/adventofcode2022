package logic

import (
	"reflect"
	"testing"
)

func Test_Split(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		want    [][]string
		wantErr bool
	}{
		{
			name: "Small Test",
			in: []byte(
				`vJrwpWtwJgWrhcsFMMfFFhFp
			jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
			PmmdzqPrVvPwwTWBwg`),
			want:    [][]string{{"vJrwpWtwJgWr", "jqHRNqRjqzjGDLGL", "PmmdzqPrV"}, {"hcsFMMfFFhFp", "rsFMfFZSrLrFZsSL", "vPwwTWBwg"}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		got, err := Split(tt.in)
		if (err != nil) != tt.wantErr {
			t.Fatalf("error %v, wanted error %v", err, tt.wantErr)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, wanted %v", got, tt.want)
		}
	}
}

func Test_AssignNum(t *testing.T) {
	tests := []struct {
		name    string
		param   [][]string
		want    int
		wantErr bool
	}{
		{
			name:    "First Test",
			param:   [][]string{{"vJrwpWtwJgWr", "jqHRNqRjqzjGDLGL", "PmmdzqPrV"}, {"hcsFMMfFFhFp", "rsFMfFZSrLrFZsSL", "vPwwTWBwg"}},
			want:    96,
			wantErr: false,
		},
		{
			name:    "Small Test",
			param:   [][]string{{"abce"}, {"abcd"}},
			want:    1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AssignNum(tt.param)
			if (err != nil) != tt.wantErr {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got: %v\n want: %v\n", got, tt.want)
			}
		})
	}
}

func Test_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		in      string
		wantErr bool
	}{
		{
			name:    "test",
			want:    26,
			in:      "z",
			wantErr: false,
		},
		{
			name:    "test",
			want:    52,
			in:      "Z",
			wantErr: false,
		},
		{
			name:    "test",
			want:    1,
			in:      "a",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetValue(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got:%v\nexpected:%v\n", got, tt.want)
			}
		})
	}
}
