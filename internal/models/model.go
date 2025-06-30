package models

import (
	"math/rand"
	"time"

	"github.com/satori/go.uuid"
)

// Model is the base model/struct for any model in the application/system
type Model struct {
	ID        uint      `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// GenerateRandomString returns a random string with `n` as the length
func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func generateUUID() string {
	return uuid.NewV4().String()
}

func (model Model) IsNewRecord() bool {
	return model.ID == 0
}

func (model Model) IsSavedRecord() bool {
	return model.ID != 0
}
