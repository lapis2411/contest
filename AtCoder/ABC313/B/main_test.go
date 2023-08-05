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
				input: `3 2
1 2
2 3`,
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "t2",
			args: args{
				input: `3 2
1 3
2 3`,
			},
			want:    "-1",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `6 6
1 6
6 5
6 2
2 3
4 3
4 2`,
			},
			want:    "-1",
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
