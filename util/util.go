package util

import "fmt"

// CheckNamespaceLength checks that the given namespaceID (nID) is
// exactly 8 bytes in length, returning an error if not.
func CheckNamespaceLength(nID []byte) error {
	if len(nID) != 8 {
		return fmt.Errorf("expected namespaceID of 8 bytes, got %d", len(nID))
	}
	return nil
}
