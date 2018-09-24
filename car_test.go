package gojib

import (
	"testing"
)

func TestGetIranianCarsPrices(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"get a slice of iranian cars prices",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIranianCarsPrices()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIranianCarsPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("len(GetIranianCarsPrices()) = 0, want != 0")
			}
		})
	}
}

func TestGetForeignCarsPrices(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"get a slice of foreign cars prices",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetForeignCarsPrices()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetForeignCarsPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("len(GetForeignCarsPrices()) = 0, want != 0")
			}
		})
	}
}
