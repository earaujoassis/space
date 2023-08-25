package feature

import (
    "github.com/earaujoassis/space/internal/services/volatile"
)

// IsActive is used to check if a feature-gate `name` is currently active (through Redis keys)
func IsActive(name string) bool {
    var result bool

    volatile.TransactionsWrapper(func () {
        if !volatile.CheckFieldExistence("feature.gates", name) {
            result = false
        } else {
            result = true
        }
    })

    return result
}

// Enable makes a feature-gate `name` currently active (through Redis keys)
func Enable(name string) {
    volatile.TransactionsWrapper(func () {
        volatile.SetFieldAtKey("feature.gates", name, 1)
    })
}

// Disable makes a feature-gate `name` currently inactive (through Redis keys)
func Disable(name string) {
    volatile.TransactionsWrapper(func () {
        volatile.DeleteFieldAtKey("feature.gates", name)
    })
}
