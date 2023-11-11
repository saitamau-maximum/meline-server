package utils

import (
    crand "crypto/rand"
    "fmt"
)

func GenerateState(b int) string {
    k := make([]byte, b)
    if _, err := crand.Read(k); err != nil {
        panic(err)
    }
    return fmt.Sprintf("%x", k)
}
