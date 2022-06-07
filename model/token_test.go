package model

import "testing"

func TestGenerateToken(t *testing.T) {
	result, _ := GenerateToken("http://somehost/somepath/filetotokenize.file")
	if result != "8710f85b45205a6b49e9574ea83deb0a8241922e7a9677c4882c985d6d335d4f" {
		t.Error("Expected to be 8710f85b45205a6b49e9574ea83deb0a8241922e7a9677c4882c985d6d335d4f, but was: %", result)
	}
}

func mockGenerateToken(string) (string, error) {
	return "mock_token", nil
}
func TestGenerateTokenizedUrl(t *testing.T) {
	result, _ := GenerateTokenizedUrl(mockGenerateToken, "http://somehost/somepath/filetotokenize.file")
	if result != "http://somehost/somepath/filetotokenize.file/token=mock_token" {
		t.Error("error, was: %", result)
	}
}
