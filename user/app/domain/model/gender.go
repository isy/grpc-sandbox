package model

const (
	unknown Gender = iota
	male
	female
)

type Gender int32

func NewGender() *Gender {
	gender := unknown
	return &gender
}
