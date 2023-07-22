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
				input: "7\n6 7 2 1 3 4 5",
			},
			want:    "4\n7 5 3 2",
			wantErr: false,
		},
		{
			name: "t2",
			args: args{
				input: "3 3\noxo\noxo\noxo\n",
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: "3 3\noox\noxo\nxoo",
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: "1 7\nooooooo",
			},
			want:    "7",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: "5 15\noxooooooooooooo\noxooxooooooooox\noxoooooooooooox\noxxxooooooxooox\noxooooooooxooox",
			},
			want:    "5",
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
