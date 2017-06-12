package util

import (
	"os"
	"testing"
)

func TestBuildPath(t *testing.T) {
	os.Chdir("..")
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"workspace",
			args{"test1"},
			"workspace/test1/",
		},
		{
			"workdir",
			args{"test2"},
			"workdir/test2/",
		},
	}
	for _, tt := range tests {
		os.Setenv("CONFIGOR_WORKSPACEDIR", tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildPath(tt.args.name); got != tt.want {
				t.Errorf("BuildPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
