package server

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsErrNotServing(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "not",
			args: args{err: errors.New("some err")},
			want: false,
		},
		{
			name: "just",
			args: args{err: errNotServing},
			want: true,
		},
		{
			name: "wrapped",
			args: args{fmt.Errorf("abc: %w", errNotServing)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsErrNotServing(tt.args.err); got != tt.want {
				t.Errorf("IsErrNotAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotServing(t *testing.T) {
	err := NotServing()
	if !errors.Is(err, errNotServing) {
		t.Errorf("expect is %v, got %v", errNotServing, err)
	}
}
