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
				input: "\\//",
			},
			want:    "4\n1 4",
			wantErr: false,
		},
		{
			name: "t2",
			args: args{
				input: '\\///\_/\/\\\\/_/\\///__\\\_\\/_\/_/\',
			},
			want:    "35\n5 4 2 1 19 9",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: '///',
			},
			want:    "0\n0",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: '\///\\\\\\\',
			},
			want:    "1\n1 1",
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
