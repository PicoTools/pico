package utils

import (
	"math/rand"
	"strings"

	"github.com/PicoTools/pico/internal/constants"
)

// RandUint32 returns random uint32 value
func RandUint32() uint32 {
	return rand.Uint32()
}

// RandString returns random string with specified length
func RandString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i, cache, remain := n-1, rand.Int63(), constants.LetterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), constants.LetterIdxMax
		}
		if idx := int(cache & constants.LetterIdxMask); idx < len(constants.LetterBytes) {
			sb.WriteByte(constants.LetterBytes[idx])
			i--
		}
		cache >>= constants.LetterIdxBits
		remain--
	}
	return sb.String()
}
