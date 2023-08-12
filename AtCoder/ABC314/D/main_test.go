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
				input: `7
AtCoder
5
1 4 i
3 0 a
1 5 b
2 0 a
1 4 Y`,
			},
			want:    "atcYber",
			wantErr: false,
		},
		{
			name: "t2",
			args: args{
				input: `35
TheQuickBrownFoxJumpsOverTheLazyDog
10
2 0 a
1 19 G
1 13 m
1 2 E
1 21 F
2 0 a
1 27 b
3 0 a
3 0 a
1 15 i`,
			},
			want:    "TEEQUICKBROWMFiXJUGPFOVERTBELAZYDOG",
			wantErr: false,
		},
		{
			name: "t2",
			args: args{
				input: `7
AtCoder
6
1 4 i
3 0 a
1 5 b
2 0 a
1 4 Y
2 0 a`,
			},
			want:    "atcyber",
			wantErr: false,
		},
		{
			name: "t2",
			args: args{
				input: `7
AtCoder
6
1 4 i
3 0 a
1 5 b
2 0 a
1 4 Y
3 0 a`,
			},
			want:    "ATCYBER",
			wantErr: false,
		},
		{
			name: "t2",
			args: args{
				input: `7
AtCoder
6
1 1 i
3 0 a
1 5 b
2 0 a
1 4 Y
3 0 a`,
			},
			{
				name: "t2",
				args: args{
					input: `7
AtCoder
6
1 7 i
3 0 a
1 5 b
2 0 a
1 4 Y
3 0 a`,
				},
			},
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
