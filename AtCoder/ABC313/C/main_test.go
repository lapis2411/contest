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
4 7 3 7`,
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "t2",
			args: args{
				input: `1
313`,
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `10
999999997 999999999 4 3 2 4 999999990 8 999999991 999999993`,
			},
			want:    "2499999974",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `3
1000 0 1`,
			},
			want:    "666",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `9
0 1 1 1 1 1 1 0 0`,
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `9
1000 0 1 1000 0 1 1000 0 1`,
			},
			want:    "1998",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `2
1000 0`,
			},
			want:    "500",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `6
1000000000 1000000000 1000000000 1000000000 1000000000 1000000000`,
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `6
1000000004 1000000000 1000000000 1000000000 1000000000 1000000000`,
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `6
1000000008 1000000000 1000000000 1000000000 1000000000 1000000000`,
			},
			want:    "6",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `2
1 3`,
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `6
1 1 3 1 1 3`,
			},
			want:    "2",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `7
1 1 1 1 1 1 100`,
			},
			want:    "84",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `3
17 1 32`,
			},
			want:    "15",
			wantErr: false,
		},
		{
			name: "t3",
			args: args{
				input: `3
0 100 101`,
			},
			want:    "67",
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
