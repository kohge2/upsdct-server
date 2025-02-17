package utils

import "testing"

func TestMultiplyIntByDecimal(t *testing.T) {
	type args struct {
		i int
		d float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "正常形_小数第1位",
			args: args{
				i: 123456,
				d: 0.1,
			},
			want: 12345,
		},
		{
			name: "正常形_小数第2位",
			args: args{
				i: 123456,
				d: 0.04,
			},
			want: 4938,
		},
		{
			name: "正常形_小数第10位",
			args: args{
				i: 123456,
				d: 0.1234567891,
			},
			want: 15241,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyIntByDecimal(tt.args.i, tt.args.d); got != tt.want {
				t.Errorf("MultiplyIntByDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
