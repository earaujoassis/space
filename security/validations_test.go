package security

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/satori/go.uuid"
)

func TestValidUUID(t *testing.T) {
    assert.True(t, ValidUUID(uuid.NewV4().String()), "should validate valid UUID")
    assert.False(t, ValidUUID("not-valid-uuid"), "should invalidate invalid UUID")
    assert.False(t, ValidUUID("n0t-v4l1d-uu1d"), "should invalidate invalid UUID")
}

func TestValidBase64(t *testing.T) {
    assert.True(t,
        ValidBase64("aGxEa2ZRWVBTYlpHcGxXeUdEZUlIaWpPZnN2VkJPWm06RFJudEx6aWRvbVFzeGpuRk1wT3pjUWN2dVZ5Q1RxQ1JIRHpxTEtudW5Fd0x0bE9PQXBrVWxGZ1dhdk9Ec29URQ=="),
        "should validate valid base64 string")
}

func TestValidRandomString(t *testing.T) {
    assert.True(t, ValidRandomString("120salsdsdl912mmdsFadc"), "should validate valid Random string")
    assert.False(t, ValidRandomString("n0t-v4l1d-random-string"), "should invalidate invalid Random string")
    assert.False(t, ValidRandomString(uuid.NewV4().String()), "should invalidate invalid Random string")
    assert.True(t, ValidRandomString("DRntLzidomQsKnunEwLtlOOApkUlFgWavODsoTE"), "should validate valid Random string")
}

func TestValidToken(t *testing.T) {
    TestValidRandomString(t)
}

func TestValidEmail(t *testing.T) {
    assert.True(t, ValidEmail("example@mailer.com"), "should validate valid email")
    assert.False(t, ValidEmail("n0t v$v4lid email"), "should invalidate invalid email")
    assert.False(t, ValidEmail("n0t-v4l1d@email"), "should invalidate invalid email")
}
