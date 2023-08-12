package main

import (
	"strings"
	"testing"
)

func TestHoge(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "t1",
			args: args{
				input: `4
3
7 19 20
4
4 19 24 0
2
26 10
3
19 31 24
19`,
			},
			want:    "2\n1 4",
			wantErr: false,
		},
		{
			name: "t2",
			args: args{
				input: `3
1
1
1
2
1
3
0`,
			},
			want:    "0",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.args.input)
			got := Do(reader)
			if got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
