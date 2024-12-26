package constants

const (
	LetterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	LetterIdxBits = 6
	LetterIdxMask = 1<<LetterIdxBits - 1
	LetterIdxMax  = 63 / LetterIdxBits
)
