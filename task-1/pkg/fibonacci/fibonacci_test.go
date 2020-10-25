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
		{name: "1 - 1", args: args{1}, want: 1},
		{name: "2 - 2", args: args{2}, want: 2},
		{name: "3 - 3", args: args{3}, want: 3},
		{name: "4 - 5", args: args{4}, want: 5},
		{name: "6 - 13", args: args{6}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Number(tt.args.num); got != tt.want {
				t.Errorf("Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
