package config_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Odraxs/go-z-v-mail/server/config"
)

func TestGetZincsearchCredentials(t *testing.T) {
	test := []struct {
		name     string
		expected *config.ZincsearchCredentials
	}{
		{
			name:     "Getting zincsearch credentials when not loaded",
			expected: nil,
		},
		{
			name: "Getting zincsearch credentials when loaded",
			expected: &config.ZincsearchCredentials{
				User:     "admin",
				Password: "password",
			},
		},
	}

	for _, test := range test {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(tt.name, "when loaded") {
				config.LoadZincsearchCredentials()
			}

			credentials := config.GetZincsearchCredentials()

			if !reflect.DeepEqual(credentials, tt.expected) {
				t.Fatalf("Expected %v but got %v", tt.expected, credentials)
			}
		})
	}
}

func TestLoadZincsearchCredentials(t *testing.T) {
	test := []struct {
		name     string
		expected config.ZincsearchCredentials
	}{
		{
			name: "Load credentials",
			expected: config.ZincsearchCredentials{
				User:     "admin",
				Password: "password",
			},
		}, {
			name: "Load already loaded credentials",
			expected: config.ZincsearchCredentials{
				User:     "admin",
				Password: "password",
			},
		}}

	for _, test := range test {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			credentials := config.LoadZincsearchCredentials()
			if !reflect.DeepEqual(credentials, tt.expected) {
				t.Fatalf("Expected %v but got %v", tt.expected, credentials)
			}
		})
	}
}
