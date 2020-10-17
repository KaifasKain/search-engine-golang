package fibonacci

import "testing"

func TestNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first - zerro", args{1}, 0},
		{"second - one", args{2}, 1},
		{"eight - thirteen", args{8}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Number(tt.args.num); got != tt.want {
				t.Errorf("Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
