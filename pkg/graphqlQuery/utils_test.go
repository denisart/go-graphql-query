package graphqlQuery

import "testing"

func Test_lineShift(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"One line",
			args{
				line: "aaabbbccc",
			},
			"aaabbbccc",
		},
		{
			"Three lines",
			args{
				line: "aaa\nbbb\nccc",
			},
			"aaa\n  bbb\n  ccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lineShift(tt.args.line); got != tt.want {
				t.Errorf("lineShift() = %v, want %v", got, tt.want)
			}
		})
	}
}
