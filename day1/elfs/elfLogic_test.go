package elfs

import (
	"reflect"
	"testing"
)

func Test_ElfParser(t *testing.T) {

	type args struct {
		in []byte
	}

	type wants struct {
		is [][]int
		i  int
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
		want    wants
	}{
		{
			name:    "Working Case",
			wantErr: false,
			args: args{
				in: []byte(`
					123
					456

					321
					432

					857
					134
					`),
			},

			want: wants{
				is: [][]int{
					{991}, {753}, {579},
				},
				i: 2323,
			},
		},
	}

	for _, tt := range tests {
		got, err := ElfParser(tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Fatalf("error %v expected want error %v", err, tt.wantErr)
		}

		if !reflect.DeepEqual(got, tt.want.is) {
			t.Fatalf("got %v wanted %v", got, tt.want)
		}

		g, err := Topthree(got)
		if (err != nil) != tt.wantErr {
			t.Fatalf("error %v expected want error %v", err, tt.wantErr)
		}

		if !reflect.DeepEqual(g, tt.want.i) {
			t.Fatalf("got %v wanted %v", got, tt.want)
		}

	}
}
