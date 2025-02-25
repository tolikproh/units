package units

import (
	"testing"
)

// TestPrefixUint проверяет метод Uint для префиксов
func TestPrefixUint(t *testing.T) {
	prefix := Prefix(1) // Пример префикса
	expected := uint64(1)

	if prefix.Uint() != expected {
		t.Errorf("Expected %d, got %d", expected, prefix.Uint())
	}
}
