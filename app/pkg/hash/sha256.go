package hash

import (
	"crypto/sha256"
	"fmt"
)

func HashString(s string) string {
	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
