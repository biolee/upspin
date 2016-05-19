package endpoint

import (
	"strings"
	"testing"
)

func TestParseAndString(t *testing.T) {
	assertParsesAndEncodes(t, "gcp,https://localhost:8080")
	assertParsesAndEncodes(t, "remote,https://localhost:8080")
	assertParsesAndEncodes(t, "inprocess")
}

func TestErrorCases(t *testing.T) {
	assertError(t, "remote", "requires a netaddr")
	assertError(t, "supersonic,https://supersonic.com", "unknown transport type")
	assertError(t, "gcp", "requires a netaddr")
}

func assertError(t *testing.T, epString string, substringError string) {
	_, err := Parse(epString)
	if err == nil {
		t.Fatal("Expected error")
	}
	if !strings.Contains(err.Error(), substringError) {
		t.Errorf("Expected error prefix %q, got %q", substringError, err)
	}
}

func assertParsesAndEncodes(t *testing.T, epString string) {
	ep, err := Parse(epString)
	if err != nil {
		t.Fatal(err)
	}
	retStr := String(ep)
	if retStr != epString {
		t.Errorf("Expected %s, got %s", epString, retStr)
	}
}