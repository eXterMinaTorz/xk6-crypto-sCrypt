package scrypt

import (
	"testing"
)

func TestGenerateHash(t *testing.T) {
	m := &ScryptModule{}
	input := "testinput"
	opts := map[string]interface{}{
		"N": 16384,
		"r": 8,
		"p": 1,
		"keyLen": 64,
		"saltLen": 16,
	}
	result, err := m.GenerateHash(input, opts)
	if err != nil {
		t.Fatalf("GenerateHash returned error: %v", err)
	}
	if result.Hash == "" {
		t.Error("GenerateHash returned empty hash")
	}
	if result.Salt == "" {
		t.Error("GenerateHash returned empty salt")
	}
	if result.Input != input {
		t.Errorf("GenerateHash returned wrong input: got %s, want %s", result.Input, input)
	}
	t.Logf("Hash: %s", result.Hash)
	t.Logf("Salt: %s", result.Salt)
	t.Logf("Input: %s", result.Input)
}
