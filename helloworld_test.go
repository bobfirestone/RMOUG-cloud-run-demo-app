package main

import "testing"

func Test_greet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "hello", want: "Hello, world! This is running"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greet(); got != tt.want {
				t.Errorf("greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
