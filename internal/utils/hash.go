package utils

import (
	"encoding/binary"

	"github.com/PicoTools/pico/internal/constants"
	"github.com/orisano/wyhash"
)

// CalcHash returns wyhash hash
func CalcHash(data []byte) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, wyhash.Sum64(constants.WyhashSeed, data))
	return b
}
