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
				input: `8 3
apzbqrcs
1 2 3 1 2 2 1 2`,
			},
			want:    "cszapqbr",
			wantErr: false,
		},
		{
			name: "t2",
			args: args{
				input: `2 1
aa
1 1`,
			},
			want:    "aa",
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
