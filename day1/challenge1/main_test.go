package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_CalorieCounter(t *testing.T) {
	type args struct {
		calories []int
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
		want    int
	}{
		{name: "Passing Test", wantErr: false, args: args{calories: []int{7127, 2867}}, want: 9994},
		{name: "Small Numbers", wantErr: false, args: args{calories: []int{10, 10}}, want: 20},
		{name: "Negative Numbers", wantErr: false, args: args{calories: []int{-10, 10}}, want: 0},
		{name: "Negative Numbers 2", wantErr: false, args: args{calories: []int{-10, -10}}, want: -20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := CalorieCounter(tt.args.calories)
			fmt.Println("ERRNIL", err != nil)

			if (err != nil) != tt.wantErr {
				t.Fatalf("got error %v want error %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, wanted %v", got, tt.want)
			}
		})
	}
}

func Test_ElfParser(t *testing.T) {

	type args struct {
		in string
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
		want    []int
	}{
		{
			name:    "Working Case",
			wantErr: false,
			args: args{
				in: `
					123
					456
					`,
			},
			want: []int{123, 456},
		},
		{
			name:    "If Negative",
			wantErr: false,
			args: args{
				in: `
				-2132 
				-11441
					`,
			},
			want: []int{-2132, -11441},
		},
		{
			name:    "if single line",
			wantErr: false,
			args: args{
				in: "12414 343241",
			},
			want: []int{12414, 343241},
		},
	}

	for _, tt := range tests {
		got, err := ElfParser(tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Fatalf("error %v expected want error %v", err, tt.wantErr)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v wanted %v", got, tt.want)
		}
	}
}

func TestMain(t *testing.T) {

	type args struct {
		cal []int
		in  string
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
		want    int
	}{
		{
			name:    "initial test",
			wantErr: false,
			args: args{
				in:  "30123 32932",
				cal: []int{30123, 32932},
			},
			want: 63055,
		},
	}

	for _, tt := range tests {
		gotString, err := ElfParser(tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Fatalf("error %v want error %v", err, tt.wantErr)
		}

		got, err := CalorieCounter(gotString)
		if (err != nil) != tt.wantErr {
			t.Fatalf("error %v want error %v", err, tt.wantErr)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v want %v", got, tt.want)
		}

	}
}
