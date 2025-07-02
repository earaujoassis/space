package security

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidUUID(t *testing.T) {
	assert.True(t, ValidUUID(uuid.NewV4().String()))
	assert.False(t, ValidUUID("not-valid-uuid"))
	assert.False(t, ValidUUID("n0t-v4l1d-uu1d"))
}

func TestValidBase64(t *testing.T) {
	assert.True(t,
		ValidBase64("aGxEa2ZRWVBTYlpHcGxXeUdEZUlIaWpPZnN2VkJPWm06RFJudEx6aWRvbVFzeGpuRk1wT3pjUWN2dVZ5Q1RxQ1JIRHpxTEtudW5Fd0x0bE9PQXBrVWxGZ1dhdk9Ec29URQ=="))
}

func TestValidRandomString(t *testing.T) {
	assert.True(t, ValidRandomString("120salsdsdl912mmdsFadc"))
	assert.False(t, ValidRandomString("n0t-v4l1d-random-string"))
	assert.False(t, ValidRandomString(uuid.NewV4().String()))
	assert.True(t, ValidRandomString("DRntLzidomQsKnunEwLtlOOApkUlFgWavODsoTE"))
}

func TestValidToken(t *testing.T) {
	TestValidRandomString(t)
}

func TestValidEmail(t *testing.T) {
	assert.True(t, ValidEmail("example@mailer.com"))
	assert.False(t, ValidEmail("n0t v$v4lid email"))
	assert.False(t, ValidEmail("n0t-v4l1d@email"))
}
