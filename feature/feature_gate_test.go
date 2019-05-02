package feature

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestIsActive(t *testing.T) {
    assert.False(t, IsActive("no-feature"), "shouldn't have no-feature active")
}

func TestEnable(t *testing.T) {
    assert.False(t, IsActive("not-enabled"), "shouldn't have not-enabled active")
    Enable("not-enabled")
    assert.True(t, IsActive("not-enabled"), "should have no-feature active")
}

func TestDisable(t *testing.T) {
    assert.False(t, IsActive("to-disable"), "shouldn't have to-disable active")
    Enable("to-disable")
    assert.True(t, IsActive("to-disable"), "should have to-disable active")
    Disable("to-disable")
    assert.False(t, IsActive("to-disable"), "shouldn't have to-disable active")
}
