package main

import "testing"

func TestCodelandUsernameValidation(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid username with underscore",
			args: args{str: "user_1"},
			want: "true",
		},
		{
			name: "starts with number",
			args: args{str: "1user"},
			want: "false",
		},
		{
			name: "ends with underscore",
			args: args{str: "user_"},
			want: "false",
		},
		{
			name: "too short",
			args: args{str: "us"},
			want: "false",
		},
		{
			name: "contains invalid character",
			args: args{str: "user@name"},
			want: "false",
		},
		{
			name: "valid username no special chars",
			args: args{str: "username123"},
			want: "true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CodelandUsernameValidation(tt.args.str); got != tt.want {
				t.Errorf("CodelandUsernameValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}
