package constants

const (
	// WyhashSeed is seed used for initialization of wyhash algorithm.
	// In case of changing this value on dirty DB (with data) - blobber
	// will create new instance of blobs in DB, so change it only on clean DB
	// to preserve space.
	WyhashSeed = 0x8d19a91
)
