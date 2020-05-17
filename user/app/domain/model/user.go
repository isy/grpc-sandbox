package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type (
	User struct {
		ID        int64
		FirstName string
		LastName  string
		Gender    Gender
		CreatedAt *timestamp.Timestamp
	}
)
