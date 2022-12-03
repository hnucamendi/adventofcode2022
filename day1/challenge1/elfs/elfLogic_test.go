package elfs

import (
	"reflect"
	"testing"
)

func Test_ElfParser(t *testing.T) {

	type args struct {
		in []byte
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
		want    [][]int
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
					`),
			},

			want: [][]int{
				{579}, {753},
			},

			// want: []interface{}{
			// 	[]int{
			// 		579,
			// 	},
			// 	[]int{
			// 		753,
			// 	},
			// },
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
