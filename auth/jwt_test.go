package auth

import (
	"bytes"
	"testing"
)

func TestEmbed(t *testing.T) {
	want := []byte("-----BEGIN PUBLIC KEY-----")
	if !bytes.Contains(rowPubKey, want) {
		t.Errorf("want %s, but got %s", want, rowPubKey)
	}

	want = []byte("-----BEGIN RSA PRIVATE KEY-----")
	if !bytes.Contains(rowPrivKey, want) {
		t.Errorf("want %s, but got %s", want, rowPrivKey)
	}
}
