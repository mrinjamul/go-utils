package utils

import (
	"testing"
)

// TestSha256sum tests the Sha256sum function
func TestSha256sum(t *testing.T) {
	testcases := []struct {
		input string
		want  string
	}{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"utils", "7574696c73e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	}
	for _, tc := range testcases {
		got, err := Sha256sum(tc.input)
		if err != nil {
			t.Errorf("Sha256sum(%q) returned error: %v", tc.input, err)
		}

		if got != tc.want {
			t.Errorf("Sha256(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}

}

// TestShortGitCommit tests the ShortCommit function
func TestShortGitCommit(t *testing.T) {
	testcases := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"354395e", "354395e"},
		{"354395e52193585956dd5b409d3e1e425e28c41f", "354395e"},
	}
	for _, tc := range testcases {
		got := ShortGitCommit(tc.input)
		if got != tc.want {
			t.Errorf("ShortCommit(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}
