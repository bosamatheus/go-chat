package handler

import (
	"io"
	"testing"

	"github.com/go-redis/redis/v8"
)

func Test_unsafeError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil", args{nil}, true},
		{"not nil", args{redis.Nil}, true},
		{"EOF", args{io.EOF}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unsafeError(tt.args.err); got != tt.want {
				t.Errorf("unsafeError() = %v, want %v", got, tt.want)
			}
		})
	}
}
