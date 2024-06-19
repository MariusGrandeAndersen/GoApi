package service

import (
	"go_api/internal/service"
	"testing"
)

func TestConvertNameToUpperCase(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"alice", "ALICE"},
		{"bob", "BOB"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.ConvertNameToUpperCase(tt.name)
			if result != tt.expected {
				t.Errorf("ConvertNameToUpperCase(%s) = %s; want %s", tt.name, result, tt.expected)
			}
		})
	}
}
