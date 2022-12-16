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

func Test_Split3Ways(t *testing.T) {
	tests := []struct {
		name    string
		in      []byte
		want    [][]string
		wantErr bool
	}{
		{
			name: "test",
			in: []byte(
				`vJrwpWtwJgWrhcsFMMfFFhFp
			jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
			PmmdzqPrVvPwwTWBwg
			wDWgDfWNTvwvgFfWfddGldJVprrrVdNlrN
			nLnmLSnmMVJvSrHqdV
			MsmsbLvtzMjFsCPDsfBwwT`),
			want:    [][]string{{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}, {"wDWgDfWNTvwvgFfWfddGldJVprrrVdNlrN", "nLnmLSnmMVJvSrHqdV", "MsmsbLvtzMjFsCPDsfBwwT"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Split3Ways(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatal("error", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v expected %v\n", got, tt.want)
			}
		})
	}
}

func Test_AssignNum3Ways(t *testing.T) {
	tests := []struct {
		name    string
		in      [][]string
		want    int
		wantErr bool
	}{
		{
			name:    "test",
			in:      [][]string{{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}, {"wDWgDfWNTvwvgFfWfddGldJVprrrVdNlrN", "nLnmLSnmMVJvSrHqdV", "MsmsbLvtzMjFsCPDsfBwwT"}, {"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}},
			want:    58,
			wantErr: false,
		},
	}

	// 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26
	// a,b,c,d,e,f,g,h,i,j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z

	// 27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52
	// A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z
	// 18 22 18
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AssignNum3Ways(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatal("error", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v exptected %v", got, tt.want)
			}
		})
	}
}
