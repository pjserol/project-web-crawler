package main

import (
	"testing"
)

func Test_formatURL(t *testing.T) {

	tests := []struct {
		name        string
		base        string
		url         string
		expectedURL string
	}{
		{
			name:        "test happy path",
			base:        "http://test.com",
			url:         "http://pj.com",
			expectedURL: "http://pj.com",
		},
		{
			name:        "test happy path with /",
			base:        "http://test.com/",
			url:         "/bla",
			expectedURL: "http://test.com/bla",
		},
		{
			name:        "test happy path with #",
			base:        "http://test.com/",
			url:         "#bla",
			expectedURL: "http://test.com/#bla",
		},
	}

	for _, test := range tests {
		fURL := formatURL(test.base, test.url)

		if fURL != test.expectedURL {
			t.Errorf("for %s, \nexpected url %s, \nbut got %s", test.name, test.expectedURL, fURL)
		}
	}
}
