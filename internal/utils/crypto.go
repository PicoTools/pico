package utils

import (
	"crypto/elliptic"
)

// ChoosEC returns elliptic curve based on its length
func ChooseEC(l int) elliptic.Curve {
	switch l {
	case 224:
		return elliptic.P224()
	case 256:
		return elliptic.P256()
	case 384:
		return elliptic.P384()
	case 521:
		return elliptic.P521()
	default:
		return nil
	}
}
