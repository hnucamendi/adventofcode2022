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
